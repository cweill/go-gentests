// Code generated by go-bindata.
// sources:
// templates/call.tmpl
// templates/function.tmpl
// templates/header.tmpl
// templates/inline.tmpl
// templates/inputs.tmpl
// templates/message.tmpl
// templates/results.tmpl
// DO NOT EDIT!

package bindata

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

var _templatesCallTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8e\x51\xca\xc2\x40\x0c\x84\xaf\x12\x4a\x1f\xfe\x1f\x4a\x0e\x20\x78\x00\x5f\x44\x54\xf4\x79\xd9\xa6\x1a\xa8\xab\xa4\x51\x91\x90\xbb\xdb\x2d\xd5\xf5\x29\x30\xf3\x4d\x66\xcc\x5a\xea\x38\x11\x54\x31\xf4\x7d\xe5\x6e\xf6\x64\x3d\x03\x6e\x29\x12\x3f\x48\xb2\xc2\x1d\xa4\xab\x02\xae\x86\x9d\xca\x3d\xaa\xbb\x2a\x9a\x51\x6a\xb3\xfb\x21\x01\xdd\x8b\x8a\xeb\x70\x21\xf7\x3f\x33\x09\xe9\x44\x50\x73\x03\x35\xf5\xb0\x58\x02\x6e\x82\x8c\xa6\x92\x0c\xf3\xf7\x9a\xdd\x1b\xf8\x66\x4b\xdf\x51\x58\xf3\x86\xdf\xbe\x29\x9d\xcb\x26\x10\xf7\xaf\x1b\x8d\xe4\x21\x08\x87\x96\xe3\xb8\x01\x0b\x3b\x9d\xff\xf9\xbe\x03\x00\x00\xff\xff\x9d\x9f\x57\x19\xec\x00\x00\x00")

func templatesCallTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCallTmpl,
		"templates/call.tmpl",
	)
}

func templatesCallTmpl() (*asset, error) {
	bytes, err := templatesCallTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/call.tmpl", size: 236, mode: os.FileMode(436), modTime: time.Unix(1455586613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x55\x4d\x4f\xdc\x30\x10\x3d\xef\xfe\x8a\x61\x05\x68\x53\xad\xcc\x1d\xc4\xa1\x08\x5a\x71\x28\x8b\x00\x95\x43\x55\x55\x26\x99\x50\xab\xc6\x49\x6d\x07\x8a\xac\xfc\xf7\x8e\x1d\xe7\x63\x77\x13\xca\xa5\x27\xf0\x64\xe6\xbd\x37\x9e\xe7\x59\xe7\x32\xcc\x85\x42\x58\xe4\x95\x4a\xad\x28\xd4\xa2\xae\xc1\xb9\xfd\x1c\x8e\x4f\x81\xd5\xf5\x7c\xee\x3f\x50\x84\xdd\xa1\xb1\x57\xfc\x09\xeb\x7a\x69\xe1\x83\xa5\x93\x50\x8f\xec\x2e\x01\x37\x9f\xf9\x93\xf1\x15\xdf\xbe\x1b\xab\xab\xd4\x82\xa3\x12\x91\x83\x2a\x2c\xec\xe7\xec\x5a\x0b\x65\x2f\x55\x59\x59\x43\x90\xb3\xd9\xd1\x11\x78\x38\xc8\xd0\xa4\x5a\x94\x9e\x97\x51\x58\x11\x3c\x10\x00\x01\x53\x39\xaa\x2c\x68\x79\x11\xf6\x27\xb0\x1b\x4c\x51\x3c\xa3\x0e\x21\x42\x66\x97\xe6\x36\x50\x75\x81\x4f\x02\x65\xd6\xe1\xb7\xf9\x90\x87\x30\xeb\xf1\x34\x57\x8f\xb8\x91\xee\x5c\x38\xf8\x7e\x43\xa7\xaf\x25\x86\xff\xda\x0a\x94\x06\xb7\x71\x59\xa8\xeb\x58\x26\x4b\xe3\x5f\x2f\xd0\xb7\x7c\xcd\x35\x35\x69\x51\x77\x42\xfb\x08\x83\x1d\x91\x23\x25\xce\x85\xc0\x24\x23\x31\x15\xba\xa9\xbc\x41\x53\x49\x1a\x0c\x5d\x9e\xad\xb4\x32\x17\x5a\x17\xba\xe5\xbd\xf8\x53\x62\x6a\x31\x03\xdd\x64\x4d\xb0\x47\x8c\x48\x7d\xcf\x95\x7d\x8b\x79\x97\xe9\x85\x2a\xe8\x04\x0f\x45\x21\xdb\xdc\xf9\xac\x76\xd1\x05\xeb\xf3\xf5\x31\x7c\xcc\x32\xf0\x16\x82\x94\x1b\x34\x74\xb3\x94\x91\x53\x13\x9d\x85\x68\xd8\x57\xfc\x17\x7a\x9a\x1f\x2b\xb0\xd6\x5b\xad\xe5\x6d\xb4\x36\x0e\x74\x23\x76\xd9\x76\xcb\xee\xdc\x02\x58\x18\x10\x75\x14\x32\x39\xd5\x1d\x46\x82\xd8\x29\xfb\xca\x65\x45\xfd\xba\x09\x07\xb1\xe6\x6d\x1c\x93\x3a\x36\xf0\xd3\xaa\xef\x79\x36\xe6\x8b\x08\x35\x9c\x71\x2b\xf8\x5e\x0b\x8b\x7a\x67\xe6\xa4\xf6\xf0\xe1\x95\xfa\x65\x67\x55\x9e\xa3\x76\x13\x76\xe3\x2a\x83\x65\xb8\xbc\xb5\x92\xaf\xc3\xb9\x24\xbb\xf1\xb5\xc2\xd0\x5f\x02\x91\xcf\xe2\x53\x29\xb9\xa5\xc5\x10\xed\xb1\xa0\x77\x1c\xa0\xfb\x2f\x29\x97\xb2\x0b\xbf\x69\x02\x0a\x36\x5f\xb6\xa5\x50\x09\x92\x39\xc2\x04\xc6\x80\x4f\x3a\xe4\xa5\xcf\xdb\x3b\x05\x25\x64\xe2\xff\xd2\x35\xb7\xd6\xf2\x5e\x9a\x59\x16\x00\xf3\xe5\x62\x88\xf4\x84\xc6\xf0\x47\x8c\x2a\xd1\x67\xc0\x29\x1c\x3c\xaf\xa0\x2d\x3e\x78\x5e\xac\x36\xc8\x45\x58\x53\x7d\xc5\x6a\x40\x95\x0c\x5e\xf2\xf0\x61\xcc\xd2\x42\xd1\x3a\xac\x70\x62\xda\x63\xef\x09\x46\xe6\x1c\xee\xe9\x73\x61\x7b\x5b\x76\x73\x67\xb7\x61\x2f\x2e\x93\x93\x41\x4a\x73\x0f\xc3\x67\x19\xb7\x15\x34\xd0\x67\xdc\x88\xb4\x79\xa7\x83\x29\xd0\x42\x1e\x19\xbd\x37\xde\x06\xf5\xf0\x4a\x24\xfd\x46\x6c\x4f\xe4\x1d\x32\xfe\x07\xed\x9e\xc6\x5c\xd2\xe6\x62\xe7\x88\xe5\xc5\xef\x8a\xcb\x65\x87\xb0\xda\xd4\x91\x34\x42\xe2\x40\xde\x65\x91\x56\x69\x54\xf9\x85\x66\x25\x4a\xb9\xa1\x32\x22\xf6\x36\xfa\x87\x87\x26\xd5\x0d\x4d\xe2\x77\x1e\xfd\xd8\xb6\xa7\xbf\x01\x00\x00\xff\xff\x79\xa2\x79\x79\x98\x07\x00\x00")

func templatesFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesFunctionTmpl,
		"templates/function.tmpl",
	)
}

func templatesFunctionTmpl() (*asset, error) {
	bytes, err := templatesFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/function.tmpl", size: 1944, mode: os.FileMode(420), modTime: time.Unix(1458492180, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x48\x4d\x4c\x49\x2d\x52\xaa\xad\xe5\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xa8\xae\xd6\x0b\x80\x30\x81\x82\x5c\x99\xb9\x05\xf9\x45\x25\x0a\x1a\x5c\xd5\xd5\x45\x89\x79\x40\x69\x3d\x4f\xb0\x48\x71\x6d\x2d\x50\xa1\x5f\x62\x2e\x50\x15\x44\x4b\x49\x06\x50\x7d\x75\x75\x6a\x5e\x0a\x90\xd6\x84\xb3\x00\x01\x00\x00\xff\xff\x81\x22\x53\x6f\x6b\x00\x00\x00")

func templatesHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeaderTmpl,
		"templates/header.tmpl",
	)
}

func templatesHeaderTmpl() (*asset, error) {
	bytes, err := templatesHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/header.tmpl", size: 107, mode: os.FileMode(436), modTime: time.Unix(1455586613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\xcb\x01\xd2\x4a\xb5\xb5\x0a\xd5\xd5\x25\xa9\xb9\x05\x39\x89\x25\x40\xd1\xe4\xc4\x9c\x1c\x25\x05\x3d\xb0\x68\x6a\x5e\x4a\x6d\x2d\x20\x00\x00\xff\xff\xaa\xeb\x41\xff\x31\x00\x00\x00")

func templatesInlineTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInlineTmpl,
		"templates/inline.tmpl",
	)
}

func templatesInlineTmpl() (*asset, error) {
	bytes, err := templatesInlineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inline.tmpl", size: 49, mode: os.FileMode(420), modTime: time.Unix(1458418263, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInputsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcb\x41\x0a\x02\x51\x08\xc6\xf1\xab\xc8\x30\xcb\x78\x07\x08\x3a\x40\xbb\xae\xf0\x60\x3e\x43\x28\x09\xc7\x56\x1f\x73\xf7\x46\x57\xad\x94\x9f\x7f\xc9\x0d\x6a\x0e\x59\xcc\x3f\xdf\xdc\x97\xe3\x20\x57\x95\xeb\x4d\x46\xad\xa6\xb2\xea\x78\x84\x79\xde\x3b\x28\x8c\xe9\x4f\xb4\xcf\x98\x6f\x24\xe2\xe4\xcc\x41\x36\xd4\xe7\x45\x48\xf8\x56\x35\x5e\x3b\xfa\xec\x67\xfb\xef\x35\x7e\x01\x00\x00\xff\xff\x43\x89\x5c\xae\x80\x00\x00\x00")

func templatesInputsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInputsTmpl,
		"templates/inputs.tmpl",
	)
}

func templatesInputsTmpl() (*asset, error) {
	bytes, err := templatesInputsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inputs.tmpl", size: 128, mode: os.FileMode(436), modTime: time.Unix(1455586613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMessageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\xd1\xca\x82\x40\x10\x46\x5f\x65\x10\x85\xff\x07\x99\x07\x08\x7a\x80\x6e\x22\x22\xba\x5f\xf2\xd3\x06\x74\xb3\xdd\xd5\x88\x61\xde\x3d\x15\x13\xba\xfa\xe0\x70\x38\x33\xaa\x15\x6a\xf1\xa0\xac\x43\x8c\xae\x41\x66\xa6\x2a\x35\xf9\x47\x22\x3e\x05\xf1\xe9\xe0\xfb\x21\x45\xb3\xe2\xc9\xa4\x0a\x5f\xcd\xc6\x4b\xd2\x9d\xf8\x8c\x1b\x64\x44\x98\x09\x5f\xde\x3d\xf8\xea\xda\x01\x66\xbc\x89\x7c\x74\xdd\x04\xfe\x96\xe8\x6f\x50\x35\x38\xdf\x80\x72\x29\x29\x47\x4b\xbb\xfd\x24\xb8\x30\xf9\x09\x21\xae\x7f\xe4\x62\x56\x7e\xef\x16\xe3\xd6\x5d\xe6\x7f\xdd\x4f\x00\x00\x00\xff\xff\xeb\x6d\x22\x24\xc6\x00\x00\x00")

func templatesMessageTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMessageTmpl,
		"templates/message.tmpl",
	)
}

func templatesMessageTmpl() (*asset, error) {
	bytes, err := templatesMessageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/message.tmpl", size: 198, mode: os.FileMode(436), modTime: time.Unix(1455586613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResultsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcd\x41\x0a\xc2\x40\x0c\x05\xd0\xab\x7c\x4a\x97\xa5\x07\x10\x5c\x8a\x7b\x6f\x20\x34\x95\x81\x92\x81\x3f\xd3\x55\xc8\xdd\x4d\x6a\x51\x70\x35\x93\xfc\x97\xc4\x6c\x91\xb5\xa8\x60\xa0\xb4\x7d\xeb\x6d\x70\x87\x19\x9f\xfa\x12\x8c\x65\xc2\x28\x1b\x2e\x57\xcc\x8f\x4f\xec\x6e\x56\xd6\x48\xdc\xa7\x70\xa2\x4b\x76\xee\xb5\x63\xce\xcf\x59\x87\x88\x81\xbe\x53\xdb\x8d\xac\x4c\x2c\xe4\x99\xe3\x00\x95\xdf\xa5\xff\x38\x0f\xfe\xec\xf1\xbe\x03\x00\x00\xff\xff\xb0\x4f\xcf\x61\xa8\x00\x00\x00")

func templatesResultsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesResultsTmpl,
		"templates/results.tmpl",
	)
}

func templatesResultsTmpl() (*asset, error) {
	bytes, err := templatesResultsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/results.tmpl", size: 168, mode: os.FileMode(436), modTime: time.Unix(1455586613, 0)}
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
	"templates/call.tmpl": templatesCallTmpl,
	"templates/function.tmpl": templatesFunctionTmpl,
	"templates/header.tmpl": templatesHeaderTmpl,
	"templates/inline.tmpl": templatesInlineTmpl,
	"templates/inputs.tmpl": templatesInputsTmpl,
	"templates/message.tmpl": templatesMessageTmpl,
	"templates/results.tmpl": templatesResultsTmpl,
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
		"call.tmpl": &bintree{templatesCallTmpl, map[string]*bintree{}},
		"function.tmpl": &bintree{templatesFunctionTmpl, map[string]*bintree{}},
		"header.tmpl": &bintree{templatesHeaderTmpl, map[string]*bintree{}},
		"inline.tmpl": &bintree{templatesInlineTmpl, map[string]*bintree{}},
		"inputs.tmpl": &bintree{templatesInputsTmpl, map[string]*bintree{}},
		"message.tmpl": &bintree{templatesMessageTmpl, map[string]*bintree{}},
		"results.tmpl": &bintree{templatesResultsTmpl, map[string]*bintree{}},
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

