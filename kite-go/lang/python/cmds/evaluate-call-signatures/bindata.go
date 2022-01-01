// Code generated by go-bindata.
// sources:
// templates/coverage/method.html
// templates/coverage/package.html
// templates/coverage/toplevel.html
// templates/diff/missing.html
// templates/diff/toplevel.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _templatesCoverageMethodHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\xcd\x6e\xdb\x30\x0c\xbe\xef\x29\x34\x1f\x0b\x54\xda\xda\xed\x92\x29\x01\x0a\x14\xd8\x61\x3f\x28\x16\xec\x01\x14\x89\x89\x95\xda\x92\x26\x31\x5d\x8c\x34\xef\x3e\x2a\xb6\xeb\xb4\xeb\xd0\x66\x48\x7c\x88\x49\x8a\x3f\xdf\x47\xd3\x8c\xe5\x5b\xe3\x35\x36\x01\x58\x89\x75\x35\x79\x23\xdb\x1b\x63\xb2\x04\x65\xb2\x90\x2f\x89\x16\x2b\x98\x6c\x36\xfc\x1b\x60\xe9\x0d\xff\xae\x6a\xd8\x6e\xa5\x68\xed\xbd\x57\xd2\xd1\x06\x64\x39\xdd\xb8\x40\x58\xa3\x58\xaa\x3b\xd5\x5a\x0b\x96\xa2\x1e\x17\x25\x62\x48\x23\x21\xb4\x37\xc0\x97\xbf\x56\x10\x1b\xae\x7d\x2d\x5a\xf1\xfc\x82\xbf\xe7\x1f\x78\x6d\x1d\x5f\xa6\x62\x22\x45\x1b\xfb\x3f\x05\x6a\xb5\xd6\xc6\xf1\x99\xf7\x98\x30\xaa\x90\x95\x5c\xe8\xc1\x20\x2e\xf9\x25\xff\x28\x96\x69\x30\xfd\xbb\x70\x65\xdd\x2d\x8b\x50\x8d\x8b\x84\x4d\x05\xa9\x04\xa0\x82\x7b\x38\x74\x4a\x05\x2b\x23\xcc\x0f\x47\x40\xa1\x4f\x20\xe4\x64\x84\x21\x17\x1d\xa8\xe7\xba\xbd\x36\xf3\xa6\x61\x9b\x4e\xc9\x57\x50\xc6\x58\xb7\x18\xb1\x8b\x77\x61\xfd\xa9\x3b\xd8\x76\x77\x34\x8f\x7c\xe7\xde\xe1\xf9\x5c\xd5\xb6\x6a\x46\xac\xf6\xce\xa7\xa0\x34\x3c\x0d\xa2\x16\xf4\x15\xa5\xe8\x67\x41\xe6\xc2\x2d\x08\x69\xec\x1d\xd3\x95\x4a\x89\x3a\xa0\x66\x15\x14\xc3\xb0\x64\xf5\xd1\x19\xdb\xfd\x9e\x13\x43\x1b\xc0\x3c\x78\x66\xdf\x21\x63\x6f\x89\xfb\x6a\x36\x18\xa6\x7d\x45\x20\xdd\xb8\xb8\xa4\xbe\xcc\x26\x79\xfa\xa4\x98\x51\x8b\xd0\xbc\xe8\xfc\x99\xba\x5a\xbe\xda\xfb\xc6\x87\x55\xa5\xa2\xc5\xe6\xb9\x10\xd2\xe3\x41\x60\x9f\x79\x63\x5e\xc2\xa0\xba\x31\x1a\x42\xaf\x5c\x93\xa3\xd9\x3d\x73\xf4\xd6\xfc\xfc\xf1\x75\xbb\xdd\xcf\xdc\x1d\xe7\xe4\xea\x15\x24\x87\xc0\x81\xeb\xdf\xc0\x0e\x66\x4a\xbd\xbb\x8a\x0b\x36\x0d\xa0\x8f\xd2\x39\x19\xe2\xfe\xc2\xa1\xdc\x39\x35\xf5\x40\xb5\xd2\x94\x66\xc9\x2d\x32\xf0\xec\x78\x0c\xf8\x53\xbb\x70\x0a\x57\x11\xd8\x8d\x42\x84\xe8\xd2\x29\x88\xf4\xb9\x89\x49\xe8\xc4\xe3\x53\xb9\xf6\x27\x79\x08\x94\x96\x60\xd3\x1f\xc6\xf1\x11\x9f\x9d\xdd\xfe\xa6\x27\x7b\x92\x96\x7f\xd9\x65\x26\xe4\x6d\x89\x13\x4c\x8e\x5f\x45\x0d\xec\x1a\xe6\xd6\x59\xb4\xde\x9d\x82\x45\x57\xe4\x9e\xa5\x9d\x70\x08\x0b\xd2\xf6\xd7\x2c\xa9\x79\x1d\x77\x7b\x5c\xd0\x22\x9f\x30\xb6\xdb\xf3\xad\x17\xed\xfb\xdd\x47\xc0\x9f\x00\x00\x00\xff\xff\x47\xee\xc5\xba\x1c\x08\x00\x00")

func templatesCoverageMethodHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCoverageMethodHtml,
		"templates/coverage/method.html",
	)
}

func templatesCoverageMethodHtml() (*asset, error) {
	bytes, err := templatesCoverageMethodHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/coverage/method.html", size: 2076, mode: os.FileMode(420), modTime: time.Unix(1471647904, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesCoveragePackageHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\xdd\x6e\xdb\x36\x14\xbe\xdf\x53\x70\x42\x80\xae\xc5\x22\x25\xcd\x76\xe3\xc9\x01\x82\x06\xdb\x80\x6d\x41\xb0\x62\x0f\x40\x4b\xb4\xad\x58\x22\x35\x92\xce\x62\x68\x7a\xf7\x1d\xea\x97\x92\x6c\xcf\xa4\x78\x53\xa0\xb9\xa8\xc5\x43\x9e\xef\x7c\x64\x0f\xcf\x8f\x14\x7e\x1b\xb3\x48\x1e\x72\x82\xb6\x32\x4b\xef\xbf\x09\xeb\x1f\x04\x7f\xe1\x96\xe0\xb8\x7e\xac\x86\x32\x91\x29\xb9\x2f\x0a\xff\x79\xb7\xf1\x9f\x70\x46\xca\x32\x0c\x6a\x61\xbf\x48\x44\x3c\xc9\x25\x52\x88\x4b\x4f\x92\x37\x19\xbc\xe0\x57\x5c\x4b\x3d\x24\x78\xb4\xf4\xb6\x52\xe6\x62\x11\x04\x11\x8b\x89\xff\xf2\xf7\x9e\xf0\x83\x1f\xb1\x2c\xa8\x1f\xaf\x3f\xfa\xb7\xfe\x0f\x7e\x96\x50\xff\x45\x78\xf7\x61\x50\xeb\xda\x99\xc8\xf0\x5b\x14\x53\x7f\xc5\x98\x14\x92\xe3\x5c\x0d\x94\xa9\x4e\x10\xdc\xf9\x77\xfe\x8f\xc1\x8b\xe8\x45\xe7\x4c\xa7\x09\xdd\x21\x4e\xd2\xa5\x27\xe4\x21\x25\x62\x4b\x08\x98\xd4\x98\x44\x42\x78\x68\xcb\xc9\xda\x9c\x03\xa8\x8e\x48\x28\x30\x60\xa1\x8c\xea\xdb\x57\x96\xfb\xf1\x8a\xc5\x07\x54\x74\x43\xf5\x97\xe3\x38\x4e\xe8\x66\x81\x3e\xde\xe4\x6f\x3f\x75\x53\x65\xf7\x24\xe3\x91\xc6\x9a\x51\x79\xbd\xc6\x59\x92\x1e\x16\x28\x63\x94\x89\x1c\x47\xe4\x98\x2a\x1c\x4a\x6f\x3f\x0c\x7a\x17\x09\x15\x91\x96\x56\x18\x27\xaf\x28\x4a\xb1\x10\x70\x2e\x78\x95\x12\x4f\xf7\x23\x25\x18\xcc\xa2\xea\xdf\x6b\xd8\x79\x92\x93\x58\x5b\xab\x56\xeb\xb8\xad\x8c\x0f\x05\x4a\x14\xa3\x88\xa5\x40\x9b\x2e\xbd\x3b\x38\xb5\xd5\xbd\xf2\xd0\x30\x58\xc1\x01\xca\xf8\x82\xe5\xbf\xc0\xa9\x6f\x8f\xaf\x07\x09\x37\x66\x30\xbe\x27\xff\xcf\x01\x37\x8e\xd3\x68\x3e\xd0\x83\x52\x46\xff\x22\x0a\x17\xe5\xaf\x3f\x7f\x2f\xcb\x0e\xb5\x99\x53\xc0\xd8\x15\x61\x38\x83\x67\x96\xef\x53\xcc\x13\x79\x30\x38\xb8\xa7\x7d\x86\x7e\xde\xd3\x48\x26\x8c\x0a\x43\xbd\x07\xbe\x41\x9f\x73\x12\x99\xe8\x3d\x47\xd2\x4a\x4f\xd9\x7b\xc6\x52\x12\x6e\x44\x53\x99\xb3\x50\x53\xd6\x1e\x99\xf1\xc6\x0c\x55\x94\x95\x0f\x1f\x76\xff\x60\xbe\x11\xe8\x11\x4b\x6c\x68\xce\x56\x57\xd9\xfd\xcc\xf6\x3c\x22\xe8\x91\xac\x13\x9a\x98\xfe\xe7\x2b\xe3\x97\x02\xcc\xb9\x7e\xbd\x47\xc3\x3d\x82\xa8\x9b\xe1\xcb\x2e\x63\xa3\xfe\x07\x91\x5b\x16\x0b\xd0\x4d\x09\x35\xd2\x84\x13\xaa\x1c\xd4\xd4\x6c\xb6\x4f\x25\xba\xbd\xb9\xf1\x6f\x50\xbd\x81\x48\xb6\x38\x6b\xc9\xcc\xd8\x03\x87\xd6\x73\x1d\xd0\xd0\xa0\x6c\x98\x28\xcf\x76\xc0\xa2\x81\xb1\x61\xf0\x5b\xed\xea\xf3\x39\x74\x40\x8a\x05\x32\x76\x8b\xda\xed\x1d\xf8\x45\x0b\x74\x92\x86\x6d\x1e\x68\xdc\xfe\xd2\xeb\x58\x14\x1c\xd3\x0d\x41\x57\xc9\xf7\x57\x19\x5a\x2c\xd1\xe8\xee\x08\xc6\x65\x33\x2a\x4b\xf3\x8b\x0c\xd5\x0c\x40\xa3\xdb\xb2\x7c\x8f\xb4\x24\x79\x95\x69\x29\x32\xab\xe0\xbb\x24\x09\x73\xe7\xf2\xe3\xd9\xe4\x3b\xc0\x1d\xa4\xde\x7e\xe6\x72\xe0\x4a\xab\x8f\x42\x65\x89\xbe\xeb\x47\xef\x2f\xa2\xa6\x6a\xa9\xaa\xe8\x5a\xbe\x03\x39\xe3\x8b\xa2\xa8\x7e\x11\x1c\xb6\x27\x20\x38\x78\x65\xf9\x0e\x0c\x25\x6b\x90\xf8\xbf\xe2\x57\x02\xe9\x51\x05\x8d\xb2\x54\x83\x2e\x59\x16\x05\x49\x05\x50\x7f\x62\xba\x88\xc6\x6a\x33\x60\xe3\xb2\x73\x3a\x4b\x26\x6f\x42\xc4\x84\x50\x1b\x3b\x1a\x46\xed\x50\x63\xa4\x89\x5c\x32\x82\xe6\x66\xca\x06\x62\x48\x43\x04\x9e\x34\x0e\xf5\xc8\xa5\xf9\x3a\xbb\x4e\x08\xd4\x01\xa4\xe1\x30\xc8\xc1\x1a\x9b\xb1\xdc\x25\x2f\x51\x85\x8e\x09\xaf\x3a\xa2\x34\xbc\x26\xe9\x59\xe3\x76\x6c\xee\x1c\xbf\x63\x31\xa3\x52\x98\x55\xa4\xa2\x4f\xdb\x24\x8d\x39\xa1\xb6\xd5\xaa\x2d\x40\x57\x7e\xda\x00\x0c\xea\x57\x5b\x06\x5d\x32\xb6\x24\x30\x47\xbf\x2d\x6d\x6d\x6d\xdb\xea\x4e\x8a\x5d\x5b\x02\xb3\x41\x8e\x97\xbf\xb6\x74\x8c\x91\xdc\x14\xc4\xad\x0d\xab\xc2\x18\x4e\x40\xdd\x22\x31\x17\xa4\xba\x04\xb6\x20\xa7\x6a\x65\x0d\x6f\x4e\xcd\xec\x90\xd6\x11\x48\xdb\x1a\xda\x21\xab\x11\x9c\x7d\x4d\xed\x90\xd3\x04\xd0\x86\x55\x7d\xa3\x5c\xba\xd5\x18\xf0\x14\x2b\xfb\x8a\x3b\x5b\x11\xee\xb2\x01\xee\x40\xdd\xbc\x03\xfb\xfa\xbe\xe8\xeb\xfb\xa2\x2f\xf6\x7d\x91\xde\xa0\xe6\x42\xeb\x50\xab\x4b\xd7\x75\xa8\xd5\xc8\x65\x87\x9a\x0b\xad\x95\xcc\x77\x9b\xbe\x93\x84\x19\xfb\x06\x75\x00\x3b\xec\x50\xfb\x29\xa3\x16\x15\xd4\xa6\x85\xc1\x03\x8d\x3f\xed\x39\x3c\x48\xd3\x08\x5a\x6d\x6f\x54\x21\xcc\x46\x1b\xa4\x76\x7b\x34\x2d\xba\x57\xdb\x1e\xd5\x0c\x03\x60\x93\xdc\xd3\xb0\x1c\x67\x7a\x97\x44\xcf\x62\x5b\x70\xd5\xf3\xbf\x4b\x9e\x27\x71\x2d\x38\x0e\xeb\x01\x97\x2c\xcf\x20\x5b\xf0\x1c\x56\x08\x4e\xdd\xf3\x34\xf2\xe5\x55\xc8\xb4\xdd\x86\x35\xc3\x2f\x6b\x20\x50\x5f\xe1\xba\x4f\x78\x55\x2b\xdf\x7c\xe6\xab\x57\x86\x41\xf5\x71\xf8\xbf\x00\x00\x00\xff\xff\x46\x2e\xa5\xc9\x33\x1e\x00\x00")

func templatesCoveragePackageHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCoveragePackageHtml,
		"templates/coverage/package.html",
	)
}

func templatesCoveragePackageHtml() (*asset, error) {
	bytes, err := templatesCoveragePackageHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/coverage/package.html", size: 7731, mode: os.FileMode(420), modTime: time.Unix(1471899653, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesCoverageToplevelHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x58\x51\x6f\xdb\x2c\x14\x7d\xff\x7e\x05\x9f\xd5\x97\x55\xab\x9d\xb6\xdb\x4b\xe6\x44\xaa\x5a\xed\x65\x53\x15\xad\xdb\x0f\x20\x40\x1c\x12\x0c\x1e\xe0\xae\x91\xe7\xff\xbe\x4b\x6c\x27\x4e\xb6\x56\x01\xa7\x0f\x8d\xb9\xe1\x9e\x73\xa0\xf7\x72\x4c\xd3\xff\xa9\x22\x76\x53\x30\xb4\xb4\xb9\x98\xfe\x97\x36\x1f\x08\x7e\xd2\x25\xc3\xb4\x79\x84\x81\xe5\x56\xb0\xe9\x3d\x16\x02\x3d\xf1\x4c\x62\x5b\x6a\x66\xd2\xa4\x09\x77\x93\x0c\xd1\xbc\xb0\xc8\xe1\x4d\x22\xcb\x5e\x6c\xb2\xc2\xcf\xb8\x89\x46\xc8\x68\x32\x89\x96\xd6\x16\x66\x9c\x24\x44\x51\x16\xaf\x7e\x96\x4c\x6f\x62\xa2\xf2\xa4\x79\xbc\xba\x89\xaf\xe3\x0f\x71\xce\x65\xbc\x32\xd1\x34\x4d\x9a\xdc\x10\x82\x1c\xbf\x10\x2a\xe3\xb9\x52\xd6\x58\x8d\x0b\x37\x70\x44\xbb\x40\x72\x1b\xdf\xc6\x1f\x93\x95\xd9\x87\x5e\x27\x16\x5c\xae\x91\x66\x62\x12\x19\xbb\x11\xcc\x2c\x19\x03\xc2\x9e\x0e\x62\x4c\x84\x96\x9a\x2d\xfc\x15\x40\xea\x91\x04\x07\x06\x1a\x1c\xe9\x7e\xe9\x8e\xb7\x1b\xcd\x15\xdd\xa0\xaa\x1d\x20\x54\x60\x4a\xb9\xcc\xc6\xe8\x66\x54\xbc\x7c\x6a\xc3\x75\xfb\x69\x69\x6f\xe6\x42\x49\x7b\xb5\xc0\x39\x17\x9b\x31\xca\x95\x54\xa6\xc0\x84\x1d\xa7\xc0\xf2\xf7\x6c\x69\xb2\x2f\x84\xd4\x11\xef\x24\x51\xfe\x8c\x88\xc0\xc6\xc0\x1e\xe0\xb9\x60\xd1\x74\xc7\x93\x6e\x03\x07\xdf\xa2\xed\xef\x2b\x58\x25\x2f\x18\xed\xcd\x75\xb3\xfb\xb8\x5d\x4c\x1f\x06\x5c\x88\x22\xa2\x04\x48\x96\x93\xe8\x16\x76\x68\x3e\xfd\xae\x2c\x16\x68\x86\xc9\x1a\x67\xae\x1e\xe7\xb0\x6d\x96\x9e\x9c\xf8\x58\xe6\xe8\x73\x29\x89\xe5\x4a\x86\x65\xdf\xe9\x0c\x3d\x15\x8c\xf8\x67\xcf\x88\x1d\x90\xed\xb8\x67\xd8\x5a\xa6\x03\x84\x3b\xea\xe0\x64\xc7\xfc\xa0\x02\x17\x1c\x94\xe8\x18\x2f\x2f\xd7\xbf\xb0\xce\x0c\x7a\xc0\x16\x07\x51\x0f\x43\x70\x1a\x9e\x54\xa9\x09\x43\x0f\x6c\xc1\x65\x58\xc5\x38\x1d\x7d\x14\xfe\x06\x0c\x44\xb4\x77\x4b\x54\x55\xdc\x30\xfd\x86\x70\x9e\xe3\xba\x3e\x45\x1f\x64\xc1\xfa\x5c\x23\x98\x90\xc4\x6d\x05\xfb\x26\xe6\xa5\xb0\xe8\x7a\x34\x8a\x47\x28\x86\x6d\xe9\x30\x16\x56\x79\x71\x77\x65\x3c\x90\xbe\x07\xe3\xab\xc0\x95\xf4\x40\xf6\x16\xc2\x97\xf9\x4b\x53\xcf\xc3\xb8\x77\x20\xbe\xec\x6d\x1d\x0f\xfc\xb3\x77\x20\xaf\xb1\x07\x75\x01\xf4\x5b\x6b\x09\x1e\x1d\x3a\x53\x45\x29\xb0\xe6\x76\xe3\x91\x14\x6a\x20\xa1\xd6\x11\x6a\x1a\x81\x76\x11\x68\x14\x01\x16\x11\x60\x0e\x43\x6c\x61\x88\x21\x1c\x5b\xc1\x5b\x87\xf8\xab\xe4\xe1\x2e\x50\x55\x1a\xcb\x8c\xa1\x0b\xfe\xfe\xa2\x30\x68\x3c\x81\x46\x5a\x67\xa6\xae\xfd\xbd\x02\x5e\x1b\x01\x06\x5d\xd7\xf5\x3b\x94\xe2\xf6\xf5\xb5\xaa\x00\x36\xbe\x93\x9b\x47\x9c\xbb\xce\x2c\xd6\xd9\x8f\x6f\x5f\xeb\xda\x25\xf4\xbe\x71\xcd\x8a\x4f\x5a\x71\x93\xb6\x6f\xaf\xfb\x25\x17\x54\x33\x79\x27\xe9\x7d\xa9\xe1\xc1\xfa\x1e\x22\x0e\xaf\x33\xac\x0e\x2d\x10\x63\xdb\x48\xa1\x18\xbd\xc3\x6c\xbb\xc4\xd6\xc6\x7a\x70\x3e\x87\x6a\xab\xa8\xeb\xb7\xf3\x89\xfa\x07\x62\x80\x2e\xd7\x9d\xe7\xd3\x74\x84\x16\xa0\xa7\x31\xae\xf3\x29\xfa\x0b\x2f\x40\x53\xd3\xd5\x67\x2c\xa7\x63\xbc\xd3\x6d\xb2\xaa\x98\xa4\x07\x87\x02\xcc\x39\xbc\x66\x41\xc0\x5d\xc9\x76\xf7\xb9\x04\x2e\x74\xdd\x9d\xaf\x99\x09\x77\x3f\xf7\xff\x80\x3f\x01\x00\x00\xff\xff\x34\x8d\x1d\x34\x26\x10\x00\x00")

func templatesCoverageToplevelHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCoverageToplevelHtml,
		"templates/coverage/toplevel.html",
	)
}

func templatesCoverageToplevelHtml() (*asset, error) {
	bytes, err := templatesCoverageToplevelHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/coverage/toplevel.html", size: 4134, mode: os.FileMode(420), modTime: time.Unix(1471647623, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesDiffMissingHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\xcd\x8e\xd3\x30\x10\xbe\xf3\x14\x43\xd4\x6b\x6d\xd8\xc2\xa5\x38\x95\x38\x71\x61\x11\x02\xf1\x00\x4e\x3c\x6d\x52\x62\x3b\xd8\xde\x55\xa3\x92\x77\x67\x9c\xdf\x6e\x76\x2b\xd8\x15\x39\x34\x99\xdf\xef\xf3\x37\x9e\x8a\xd7\xca\xe6\xa1\xa9\x11\x8a\xa0\xab\xdd\x2b\xd1\xbf\x00\x44\x81\x52\xc5\x8f\xf8\x88\x50\x86\x0a\x77\xdf\xcb\x83\x91\xe1\xce\x21\x7c\x95\x21\xa0\x33\x5e\xf0\x3e\x32\xe6\xf9\xdc\x95\x75\x80\xd8\x30\x4d\x02\x9e\x02\x3f\xca\x7b\xd9\x7b\x13\xf0\x2e\x4f\x93\x22\x84\xda\x6f\x39\xcf\xad\x42\x76\xfc\x75\x87\xae\x61\xb9\xd5\xbc\xff\x5c\xdf\xb0\xb7\xec\x1d\xd3\xa5\x61\x47\x9f\xec\x04\xef\x6b\x5f\x02\xa0\xe5\x29\x57\x86\x65\xd6\x06\x1f\x9c\xac\xa3\x11\x81\x26\x07\xdf\xb0\x0d\x7b\xcf\x8f\x7e\x76\x5d\x07\xae\x4a\xf3\x13\x1c\x56\x69\xe2\x43\x53\xa1\x2f\x10\x09\xf0\x82\x47\xee\x7d\x02\x85\xc3\xfd\xf3\x19\x50\xe9\x82\x42\x6c\x46\x1c\x22\xe8\x7c\xf4\x88\x3b\x5a\x99\x55\x0d\x9c\x07\x23\x3e\xb5\x54\xaa\x34\x87\x2d\xdc\xbc\xa9\x4f\x1f\x86\x40\x3b\xbc\x83\x7a\x90\xbb\xb7\x26\xac\xf7\x52\x97\x55\xb3\x05\x6d\x8d\xf5\xb5\xcc\x71\x59\x44\x12\x8c\x88\x82\x8f\xb7\x41\x44\xe0\x9e\x84\x50\xe5\x3d\xe4\x95\xf4\x9e\x14\x90\x59\x85\xc9\x7c\x5d\xa2\xf9\x20\x06\xdd\xef\x9a\x4e\x58\xd6\xa8\xa6\xcc\x98\x3b\x77\x1c\x3d\xee\xd2\x8c\x0e\x05\xb9\xad\x88\xa4\x49\x93\x0d\xe9\x92\xed\xbe\x48\x8d\x82\x67\x24\x51\x50\x7f\x49\x8e\x2c\xbb\x83\xa4\x09\xf9\xad\xdb\x9e\xcf\xec\xb6\xf4\x9e\xc4\x82\xdf\xd0\xb9\xda\x36\xd9\xcd\x5e\x16\x9b\xb7\xad\xe0\x54\xf9\x08\x80\x6c\xf7\x5c\xb2\x9f\x68\xaa\xc5\x3f\xb2\x95\xc3\x15\xba\xa0\xf3\xd1\x34\x91\x11\x91\x35\xb4\x32\x3f\xbe\x7d\x5e\xd0\x1d\xe2\x91\xb1\xfc\x2f\x7c\x6f\x51\x67\xe8\xfc\x53\x8c\x97\xed\xce\x67\x27\xcd\x01\x61\xa5\x61\x9b\xc2\x44\x69\xe8\xd0\xb6\x0b\xa8\x05\x76\xac\x5f\x75\x13\x88\xd5\xd4\x63\x9a\xc7\x22\xed\x09\x99\x16\x23\x5d\x8d\x83\x9c\xf4\xeb\xda\x69\x22\x34\x29\xb6\xd2\xec\xba\x50\x2f\x87\xb9\x36\xa0\x39\x72\x0d\xf1\xb1\x9a\x68\xd4\xc5\xe1\x29\x7e\xb9\x1c\x64\xc6\x25\x1a\xb6\xaf\xbb\x9e\x00\xdd\x76\xf6\x59\xb4\xa5\xdd\x9f\xf7\x9f\x00\x00\x00\xff\xff\xa7\xb9\x2b\x0b\xd4\x05\x00\x00")

func templatesDiffMissingHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDiffMissingHtml,
		"templates/diff/missing.html",
	)
}

func templatesDiffMissingHtml() (*asset, error) {
	bytes, err := templatesDiffMissingHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/diff/missing.html", size: 1492, mode: os.FileMode(420), modTime: time.Unix(1468963465, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesDiffToplevelHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\x4d\x8f\xd3\x3c\x10\xbe\xbf\xbf\x62\xde\x68\xb9\x51\x07\xb6\x70\x29\x69\x24\x04\x12\x42\x62\x51\xc5\xc7\x19\x4d\x6d\xa7\x71\x37\xb1\x83\x3d\xbb\x34\x2a\xf9\xef\xd8\x71\x9a\xa6\x15\xb0\xea\xd2\x43\x3a\x1e\xcf\x33\xcf\x93\xf9\x48\xf6\xbf\x30\x9c\xda\x46\x42\x49\x75\x95\xff\x97\xc5\x3f\x80\xac\x94\x28\x82\xe1\x4d\x52\x54\xc9\xfc\xb3\xda\x68\xa4\x3b\x2b\x61\x85\x44\xd2\x6a\x78\xab\x8a\x22\x4b\xe3\x6d\x8c\x74\xdc\xaa\x86\x20\x24\x5c\x26\x24\x77\x94\x6e\xf1\x1e\xa3\x37\x01\x67\xf9\x32\x29\x89\x1a\xb7\x48\x53\x6e\x84\x64\xdb\xef\x77\xd2\xb6\x8c\x9b\x3a\x8d\xe6\xec\x9a\x3d\x67\x2f\x58\xad\x34\xdb\xba\x24\xcf\xd2\x88\xbd\x3c\x7d\x8d\x3b\x2e\x34\x5b\x1b\x43\x8e\x2c\x36\xe1\x10\x68\x46\x47\x3a\x67\x73\xf6\x32\xdd\xba\xa3\xeb\x4f\xb4\x95\xd2\xb7\x60\x65\xb5\x4c\x1c\xb5\x95\x74\xa5\x94\x9e\x6e\xa2\x82\x3b\x97\x40\x69\x65\x71\x39\xbf\x87\x9e\x09\x08\xc9\xbc\x82\x40\x7a\x78\xed\xc0\x1a\xed\xb5\x11\x2d\xec\x7b\x13\xa0\x41\x21\x94\xde\x2c\xe0\xfa\x59\xb3\x7b\xd5\x3b\xbb\xfe\xc9\x9c\x44\xcb\xcb\x99\x35\x3f\xc6\xe0\x1a\xed\x46\xe9\xd9\xda\x10\x99\x7a\x01\xf3\x33\x08\x89\x31\xb2\x30\x9a\x66\x05\xd6\xaa\x6a\x17\x50\x1b\x6d\x5c\x83\x5c\x4e\x83\x7d\x7d\x0e\x92\xb2\xf4\x30\x28\x59\xd0\x36\x28\x16\xea\x1e\x78\x85\xce\xf9\x02\xe1\xba\x92\x49\x3e\xe4\xce\xfa\xe3\xc9\x1d\xf4\xcf\x99\x2f\x80\x6a\xa4\x18\x23\x43\xec\x71\x06\x0f\x1e\x3b\x3d\xf6\x21\xc0\x4d\xe5\x05\xea\x65\x32\x4f\xf2\x15\xf2\x5b\xdc\x48\x3f\x94\xe5\xdf\x03\xdf\xf9\x6a\x97\x0f\x87\xdd\x28\xe7\x7c\x85\xe1\xbd\x86\x37\xa8\x85\x12\x48\xe3\xf8\xbb\x87\xe1\x5f\x0c\x61\x15\xc0\x37\xe8\x3c\xe4\x02\xe4\x13\xb8\x88\xda\x9f\x27\x95\x09\xb7\x27\x95\xcb\xe8\xd8\x9b\xf8\xdb\xef\x2d\xea\x8d\x84\xab\x6f\x4f\xe1\xaa\x1e\xa8\x16\x4b\x60\x03\x6d\xd7\x9d\xc9\x3b\x2b\x7c\x70\x89\x13\xc5\x19\x0e\x1b\xb0\xdf\x8f\x09\x7f\x42\xb0\xbe\x7e\xfa\xd0\x75\x49\x7e\xf4\xb3\x8f\x58\xcb\xae\xcb\x52\xf4\x73\x4e\xe2\x11\x99\xd9\x6b\xdd\x86\x24\x9e\x41\xfb\xef\xc8\xbf\x32\x4c\x80\x37\xa3\x74\xbf\xad\x35\x86\x1c\x17\xe1\x63\xcb\x1f\x8b\x5e\x71\x3a\x0a\x28\xc8\xfc\x36\xc3\x69\xb3\x43\x2f\xa5\x16\x93\x86\xf9\xfb\x69\xbb\xfd\x31\x2c\xd8\xb0\x99\xa9\x5f\xcd\xb8\xb7\x31\xc6\xef\x6f\xf8\xe2\xff\x0a\x00\x00\xff\xff\x4b\x00\x4e\x19\x08\x06\x00\x00")

func templatesDiffToplevelHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDiffToplevelHtml,
		"templates/diff/toplevel.html",
	)
}

func templatesDiffToplevelHtml() (*asset, error) {
	bytes, err := templatesDiffToplevelHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/diff/toplevel.html", size: 1544, mode: os.FileMode(420), modTime: time.Unix(1468963465, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"templates/coverage/method.html": templatesCoverageMethodHtml,
	"templates/coverage/package.html": templatesCoveragePackageHtml,
	"templates/coverage/toplevel.html": templatesCoverageToplevelHtml,
	"templates/diff/missing.html": templatesDiffMissingHtml,
	"templates/diff/toplevel.html": templatesDiffToplevelHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"coverage": &bintree{nil, map[string]*bintree{
			"method.html": &bintree{templatesCoverageMethodHtml, map[string]*bintree{
			}},
			"package.html": &bintree{templatesCoveragePackageHtml, map[string]*bintree{
			}},
			"toplevel.html": &bintree{templatesCoverageToplevelHtml, map[string]*bintree{
			}},
		}},
		"diff": &bintree{nil, map[string]*bintree{
			"missing.html": &bintree{templatesDiffMissingHtml, map[string]*bintree{
			}},
			"toplevel.html": &bintree{templatesDiffToplevelHtml, map[string]*bintree{
			}},
		}},
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

