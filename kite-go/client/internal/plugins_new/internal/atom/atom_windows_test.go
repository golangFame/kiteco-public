package atom

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kiteco/kiteco/kite-go/client/internal/plugins_new/editor"
	"github.com/kiteco/kiteco/kite-go/client/internal/plugins_new/internal/shared"
	"github.com/kiteco/kiteco/kite-go/client/internal/plugins_new/process"
	"github.com/stretchr/testify/require"
)

func TestInstallConfig(t *testing.T) {
	mockProcessManager := &process.MockManager{
		IsProcessRunningData: func(id string) (bool, error) {
			return id == atomBinaryName, nil
		},
	}
	mgr := newTestManager(mockProcessManager)
	require.Equal(t,
		&editor.InstallConfig{
			RequiresRestart:       true,
			Running:               true,
			InstallWhileRunning:   true,
			UpdateWhileRunning:    true,
			UninstallWhileRunning: true,
		}, mgr.InstallConfig(context.Background()))
}

func TestInstallFlow(t *testing.T) {
	dir, cleanup := shared.SetupTempDir(t, "kite-atom")
	defer cleanup()

	atomFile, atomCleanup := setupAtomBinary(t, dir)
	defer atomCleanup()

	// state use by the mock manager
	isInstalled := false

	// this process manager tracks the install and uninstall commands to mock a typical apm workflow
	processMgr := &process.MockManager{
		RunResult: func(name string, arg ...string) ([]byte, error) {
			args := strings.Join(arg, " ")
			if filepath.Base(name) == "apm" && args == "list --installed --bare --packages" {
				if isInstalled {
					return []byte("image-view@0.64.0\nkite@0.106.0\nwelcome@0.36.8"), nil
				}
				return []byte("image-view@0.64.0\nwelcome@0.36.8"), nil
			}

			if filepath.Base(name) == "apm" && args == "install kite" {
				isInstalled = true
				return []byte("Installing kite to ..."), nil
			}

			if filepath.Base(name) == "apm" && args == "uninstall kite" {
				if !isInstalled {
					return nil, fmt.Errorf("Uninstalling kite ✗\nFailed to delete kite: Does not exist")
				}

				isInstalled = false
				return []byte("Installing kite to ..."), nil
			}

			if filepath.Base(name) == "apm" && args == updateArg+" "+apmPluginID+" "+noConfirmArg {
				return []byte{}, nil
			}

			if name == atomFile && args == "--version" {
				return []byte("Atom    : 1.34.0\nElectron: 3.1.4\nChrome  : 66.0.3359.181\nNode    : 10.2.0"), nil
			}

			return nil, fmt.Errorf("unknown command")
		},
	}

	mgr := newTestManager(processMgr)
	testBasicInstallFlow(t, mgr, atomFile)
}

func TestIsInstalled(t *testing.T) {
	dir, cleanup := shared.SetupTempDir(t, "kite-atom")
	defer cleanup()

	// command output includes kite package
	processMgr := &process.MockManager{
		RunResult: func(name string, arg ...string) ([]byte, error) {
			if filepath.Base(name) == "apm" && strings.Join(arg, " ") == "list --installed --bare --packages" {
				return []byte("image-view@0.64.0\nkite@0.106.0\nwelcome@0.36.8"), nil
			}
			return nil, fmt.Errorf("unknown command")
		},
	}

	mgr := newTestManager(processMgr)
	require.True(t, mgr.IsInstalled(context.Background(), dir))
}

func TestIsNotInstalled(t *testing.T) {
	dir, cleanup := shared.SetupTempDir(t, "kite-atom")
	defer cleanup()

	// the command package doesn't include the kite package
	processMgr := &process.MockManager{
		RunResult: func(name string, arg ...string) ([]byte, error) {
			if name == "apm" && strings.Join(arg, " ") == "list --installed --bare --packages" {
				return []byte("image-view@0.64.0\nwelcome@0.36.8"), nil
			}
			return nil, fmt.Errorf("unknown command")
		},
	}

	mgr := newTestManager(processMgr)
	require.False(t, mgr.IsInstalled(context.Background(), dir))
}

func TestIsInstallApmError(t *testing.T) {
	dir, cleanup := shared.SetupTempDir(t, "kite-atom")
	defer cleanup()

	// the command package doesn't include the kite package
	processMgr := &process.MockManager{
		RunResult: func(name string, arg ...string) ([]byte, error) {
			if name == "apm" && strings.Join(arg, " ") == "list --installed --bare --packages" {
				return nil, fmt.Errorf("error executing apm command")
			}
			return nil, fmt.Errorf("unknown command")
		},
	}

	mgr := newTestManager(processMgr)
	require.False(t, mgr.IsInstalled(context.Background(), dir))
}

func TestDetectEditors(t *testing.T) {
	dir, cleanup := shared.SetupTempDir(t, "kite-atom")
	defer cleanup()

	atomPath, atomCleanup := setupAtomBinary(t, dir)
	defer atomCleanup()

	commonPath := filepath.Join(dir, "Programs", "atom", "bin", "atom")
	err := os.MkdirAll(filepath.Dir(commonPath), 0700)
	require.NoError(t, err)
	err = ioutil.WriteFile(commonPath, nil, 0700)
	require.NoError(t, err)

	// the command package doesn't include the kite package
	processMgr := &process.MockManager{
		RunResult: func(name string, arg ...string) ([]byte, error) {
			if (name == atomPath || name == commonPath) && strings.Join(arg, " ") == "--version" {
				return []byte("Atom    : 1.34.0\nElectron: 3.1.4\nChrome  : 66.0.3359.181\nNode    : 10.2.0"), nil
			}
			return nil, fmt.Errorf("unknown command")
		},
	}

	mgr := newTestManager(processMgr)
	mgr.commonPaths = []string{commonPath}

	paths, err := mgr.DetectEditors(context.Background())
	require.NoError(t, err)
	require.Len(t, paths, 2)
	editors := shared.MapEditors(context.Background(), paths, mgr)
	require.Len(t, editors, 2)

	// remove the atom file and the common dir and make sure that it's not detected
	err = os.Remove(atomPath)
	require.NoError(t, err)
	paths, err = mgr.DetectEditors(context.Background())
	require.NoError(t, err)
	editors = shared.MapEditors(context.Background(), paths, mgr)
	require.Len(t, editors, 1)

	// remove common and make sure that it's empty
	err = os.Remove(commonPath)
	require.NoError(t, err)

	paths, err = mgr.DetectEditors(context.Background())
	require.NoError(t, err)
	editors = shared.MapEditors(context.Background(), paths, mgr)
	require.Empty(t, editors)
}

func TestOpenFile(t *testing.T) {
	dir, cleanup := shared.SetupTempDir(t, "kite-atom")
	defer cleanup()

	// this process manager tracks the install and uninstall commands to mock a typical apm workflow
	processMgr := &process.MockManager{
		RunResult: func(name string, arg ...string) ([]byte, error) {
			args := strings.Join(arg, " ")
			expected := filepath.Join(dir, "atom\\bin\\atom")
			if name == expected && args == "/foo/bar:1" {
				return []byte(""), nil
			}
			return nil, fmt.Errorf("unknown command %s", name)
		},
	}

	mgr := newTestManager(processMgr)
	editorPath := filepath.Join(dir, "atom\\app-1.52.0\\resources\\app\\apm\\bin\\apm.cmd")
	_, err := mgr.OpenFile(context.Background(), "atom", editorPath, "/foo/bar", 1)
	require.NoError(t, err)
}

// setups a dummy atom binary at the given path and returns the path to it and a cleanup function
func setupAtomBinary(t *testing.T, baseDir string) (string, func()) {
	// dummy atom binary
	atomPath := filepath.Join(baseDir, "atom", "atom.cmd")
	err := os.MkdirAll(filepath.Dir(atomPath), 0700)
	require.NoError(t, err)
	err = ioutil.WriteFile(atomPath, []byte(""), 0700)
	require.NoError(t, err)

	// update path to contain our temp dir and atom sub dir
	oldPath := os.Getenv("PATH")
	err = os.Setenv("PATH", strings.Join([]string{baseDir, filepath.Dir(atomPath)}, string(os.PathListSeparator)))
	require.NoError(t, err)

	return atomPath, func() {
		os.Setenv("PATH", oldPath)
	}
}
