// Code generated by go-bindata.
// sources:
// templates/report.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesReportHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\xb0\xea\x1e\xda\x83\xc8\xa4\xa9\x73\xf0\xd2\xba\x2c\x5a\xa0\x58\x14\x28\xba\x45\xf7\x4c\x93\x23\x8b\x09\x4d\xaa\xe4\x38\xb1\x2b\xf0\xbf\x17\xb4\x3e\x22\x09\xce\xae\x5d\xec\x21\x3a\x51\xc3\x99\xf7\x38\xef\x8d\x08\xf1\x0a\x77\xa6\x58\xf0\x0a\x84\x2a\x16\x84\x10\xc2\x51\xa3\x81\xe2\x77\xa7\xc0\x10\x84\x80\x9c\xb5\x91\x76\x37\x48\xaf\x6b\x24\x78\xac\x61\x9d\x21\x1c\x90\x3d\x88\x27\xd1\x46\x33\x12\xbc\x5c\x67\x15\x62\x1d\x56\x8c\x49\xa7\x80\x3e\xfc\xb3\x07\x7f\xa4\xd2\xed\x58\xbb\xcc\x7f\xa2\xb7\xf4\x67\xba\xd3\x96\x3e\x84\xac\xe0\xac\xad\xbd\x1e\x7e\x27\x0e\x52\x59\xba\x71\x0e\x03\x7a\x51\xa7\x97\x44\x33\x04\xd8\x1d\xbd\xa3\x4b\xf6\x10\x5e\x42\xaf\xd1\x1a\x6d\x1f\x89\x07\xb3\xce\x02\x1e\x0d\x84\x0a\x00\xb3\xf1\x29\x64\x08\x19\xa9\x3c\x94\xd7\xf3\xcb\x30\x3f\x40\x02\x2b\x38\x4b\xa4\x7d\xdb\x89\xb5\x5d\xa7\x07\x15\xad\xbd\xab\x73\x2b\x76\x40\x9a\x21\x9c\x1e\xe9\x8c\xf3\x2b\xf2\xfd\xfd\xfd\xfd\xfb\x61\x23\x0e\x2b\xaa\xad\x74\xde\x83\xc4\x59\xdd\x46\xc8\xc7\xad\x77\x7b\xab\xf2\x1e\xa2\x2c\xd5\x52\xdd\x9e\x45\xb9\x18\x03\x6e\xca\x12\xce\x9e\x44\xcd\xab\x9f\xb5\xc2\x6a\x45\xee\x6e\x6e\xea\x03\xf9\x4e\xef\x6a\xe7\x51\x58\x7c\x3f\x49\xda\x09\xbf\xd5\x36\xf7\x7a\x5b\xe1\x8a\xdc\x2e\x5f\x49\x6d\x69\x38\xeb\x74\xe3\xac\x9d\x5f\xbe\x71\xea\x58\x2c\xb8\xd2\x4f\x44\x1a\x11\xc2\x3a\x93\xce\xa2\xd0\x16\x7c\xd6\x49\x3d\xda\xf3\xee\x39\x7b\x11\x9d\x57\xcb\xe2\x43\xdb\xb7\x39\x12\xad\xc0\xa2\x2e\x35\x28\x92\x5c\xc8\x9d\xcf\x1f\xe1\xf8\xec\xbc\x5a\x91\xa6\xa1\x9f\x50\x60\xa0\x5d\xfa\x6f\xe1\x63\xbb\x15\x23\x1b\xf6\xfe\x72\x28\x4c\x8c\x9c\x55\xcb\x0b\x38\x42\x0d\x52\x97\x5a\x92\x57\x49\x3e\x57\x5a\x56\xaf\xf1\x0c\xf1\x17\x3a\xce\x94\x7e\x2a\x16\x8b\xa6\xf1\xc2\x6e\x81\xbc\xf3\x10\xf6\x06\xc9\x6a\x4d\xe8\x9f\xa7\x65\x88\xf1\x6b\x92\xd4\x1e\x8a\xa6\xe9\x4a\xe9\x27\x2f\x63\xe4\x01\xbd\xb3\xdb\xe2\x5d\x52\xff\xb4\xe2\x2c\xa5\x4d\x48\xbf\x0c\xfb\xd9\x79\x15\x52\x87\x3d\xf0\x29\xd0\x9f\xa6\x85\xf8\x32\xc2\x2f\x87\x1a\x24\x82\x1a\x83\xf4\xb1\x29\xce\xd7\x3a\x9c\x8c\x8a\xc9\x77\x2a\xbf\x1f\x6d\xf7\x9e\xfd\x0a\x02\xf7\x1e\xc2\xd4\xcd\x16\xc0\xf4\xf5\xca\xe4\x95\xf3\xfa\xdf\x34\x71\x66\x06\x32\xb8\x50\x82\x38\x79\xd0\x9f\xba\x47\x8e\x71\x92\xdf\x42\x63\x12\x3f\x55\xd0\x8f\x70\x4c\xe6\x2a\x2c\xce\xa4\xa9\x21\xed\xef\x76\xe4\x94\x9a\xb3\x83\x55\x33\x06\xce\x94\x19\xe9\x30\x52\xeb\x0a\x5d\x86\xc9\x27\xb5\x07\xa5\x25\x6a\x67\x93\x27\xba\x1c\x1a\x1c\x7d\x1c\xdd\x6c\x37\x0d\x98\x00\x31\xa6\xef\xaa\x3b\x19\xf9\xa1\x26\xeb\x91\x97\x43\xd1\x1f\xde\x6d\x62\xfc\xf1\x72\xdd\xcf\x81\x7c\x48\x79\x31\x5e\x6a\xc9\x4b\x9d\xb3\xe8\xf5\x66\x9f\xba\x7a\x2b\x06\x7d\xfb\xd9\x1e\x5f\x2c\x33\x1b\x87\xef\x73\x72\xf7\xcc\xcd\xfa\x06\x56\x8d\x09\xae\x74\x6b\x52\xfa\x06\x0d\xbb\xd2\x8d\xf1\xcd\xa6\xad\x82\x03\x99\xdf\x6f\xe4\xa6\xb3\x60\x7d\xe6\xf2\xfb\xbf\x5e\x5c\xaa\xf6\x8c\xe7\x0d\x0a\x3e\x5a\xf6\x35\xdd\x2b\x67\xdd\xef\x01\x3b\xfd\xf4\xfe\x17\x00\x00\xff\xff\x94\xb1\x25\x08\xfb\x0a\x00\x00")

func templatesReportHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesReportHtml,
		"templates/report.html",
	)
}

func templatesReportHtml() (*asset, error) {
	bytes, err := templatesReportHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/report.html", size: 2811, mode: os.FileMode(420), modTime: time.Unix(1551143297, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/report.html": templatesReportHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"report.html": &bintree{templatesReportHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

