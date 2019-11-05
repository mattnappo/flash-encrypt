// Code generated for package common by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/help.txt
// assets/help_standalone.txt
package common

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsHelpTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\x31\x0e\x83\x30\x0c\x45\x77\x4e\xf1\x47\x3a\x94\x53\xc0\xd0\xb5\x37\x88\xc0\x28\x96\x5c\x83\xb0\x69\xe1\xf6\x15\x24\x61\xeb\xd0\x6c\x5f\xff\xfd\xa7\xc8\x00\x40\xda\x2f\xfb\xec\x35\xab\x63\x58\xf8\x4d\x8f\xf6\x56\xe1\x7c\x5d\xaa\x10\x30\x4a\xb0\x98\x6a\x7c\xd8\x23\x78\x28\x70\x53\x9d\xf4\x40\x3f\x35\x2d\xfd\xa1\x39\x92\xd5\x65\xfa\x24\x5f\x17\x35\x04\x08\x9b\x63\x1a\xe1\x91\xca\xc4\x8e\x1c\x44\xd0\x4f\xaa\xd4\x3b\x65\x99\x65\x57\x24\x99\x2f\x53\xcb\x36\x4b\xd8\xe1\x91\xed\x6c\xf0\x22\x5d\x33\x49\x1b\xfb\x45\x02\xdd\xc6\x9e\xfe\x7a\xcf\xd7\x69\xbe\x01\x00\x00\xff\xff\x6f\x96\xc0\xab\x2a\x01\x00\x00")

func assetsHelpTxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsHelpTxt,
		"assets/help.txt",
	)
}

func assetsHelpTxt() (*asset, error) {
	bytes, err := assetsHelpTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/help.txt", size: 298, mode: os.FileMode(420), modTime: time.Unix(1572886422, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsHelp_standaloneTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x41\x0a\x83\x40\x0c\x45\xf7\x9e\xe2\x2f\xdb\x45\xa7\x97\xa8\x07\x11\xfd\x92\x40\x6a\x87\x49\x04\xbd\x7d\x61\x1c\x18\xba\x68\x76\x8f\xf7\x12\x02\x00\xdc\xe6\x72\xe6\xb8\xdd\x07\xd4\x19\x2f\x46\x7a\x3a\xe7\xbd\x10\xeb\xc7\x16\x96\x34\x54\xbf\xf0\xb7\x7e\x5d\x8c\x10\xfe\xd9\x10\x5a\xee\xb9\x7a\xb6\xe9\x44\x88\x7a\x35\x78\x73\xdb\x5b\xc9\x43\xfb\x61\x60\x3c\x34\xb0\xda\xe4\xf2\x68\x3f\xa6\x6f\x00\x00\x00\xff\xff\x9d\x09\x79\x3c\xb0\x00\x00\x00")

func assetsHelp_standaloneTxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsHelp_standaloneTxt,
		"assets/help_standalone.txt",
	)
}

func assetsHelp_standaloneTxt() (*asset, error) {
	bytes, err := assetsHelp_standaloneTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/help_standalone.txt", size: 176, mode: os.FileMode(420), modTime: time.Unix(1572886417, 0)}
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
	"assets/help.txt":            assetsHelpTxt,
	"assets/help_standalone.txt": assetsHelp_standaloneTxt,
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
	"assets": &bintree{nil, map[string]*bintree{
		"help.txt":            &bintree{assetsHelpTxt, map[string]*bintree{}},
		"help_standalone.txt": &bintree{assetsHelp_standaloneTxt, map[string]*bintree{}},
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
