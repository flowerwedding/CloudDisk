// Code generated for package configs by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.yaml
package configs

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

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x5d\x6f\xda\x4c\x10\x85\xef\x91\xf8\x0f\x23\xbd\xd7\xaf\xb1\x4d\xf8\xda\xab\xe6\x83\x28\x54\xa1\x45\xb5\xa3\x5c\x56\x63\x3c\x98\xad\xd6\xde\x65\x77\x4c\x4c\x7e\x7d\xe5\xb5\x53\xdc\x56\xbd\x83\x67\x66\x8f\xce\x99\xe3\x84\xec\x99\xac\x18\x8f\x00\xbe\xd5\xd5\x56\xe7\x24\x20\xa7\xac\x2e\x5a\xf2\xc4\x6c\x76\xda\xb2\x80\x65\x18\x86\x7e\x87\x30\x4f\x65\x49\xba\x66\x01\x73\x8f\x5e\xad\x64\xfa\x8d\xdd\x1a\xe3\x05\x1f\xe8\x80\xb5\xe2\x1d\x16\x94\xc8\x77\x12\x10\xf9\x07\x5b\x6c\x86\xc8\xb3\x67\x5d\x24\x78\xa6\x1d\xf2\x51\x80\x63\x6d\xb1\xa0\x89\xd2\x85\xeb\x87\x8f\x52\xd1\x17\x2c\x49\x00\x1a\x33\x60\xeb\x86\x05\x04\x4a\x7b\xbb\x2f\x46\x69\xcc\xff\xd6\xa9\x3d\x77\x83\x15\x1f\xfa\xc5\x2a\x01\x47\x66\x23\x26\x93\x28\x5e\x04\x61\x10\x06\x91\x68\x93\x4e\x1c\x23\xcb\xfd\xf5\xc1\xa6\xc4\x82\xb6\xd8\x74\x9e\x67\x00\xff\xc1\xf6\xee\x8f\xf1\xad\x52\xfa\x6d\xdd\xb0\xf3\xd9\x01\xfe\x87\xe0\x87\x29\x06\xbf\xe9\xfa\xc7\x54\xc5\x78\xf4\x80\x8c\x19\x3a\xea\x6e\x75\x97\x5e\x0c\x09\x28\x2f\xee\xa4\xbc\xb2\x23\x5b\xf9\xc4\x56\x6b\x6e\xc9\x0e\x9d\x7b\xd3\x36\xf7\xfb\x4f\xda\xb1\x80\xab\xed\xe9\x34\x9c\x77\x3a\xdd\x99\xca\x4b\xd6\x5f\x25\xc5\x4c\xd1\xce\xd2\x41\x36\x02\x5a\xf8\xbd\xa5\xf7\x47\xb4\x8e\x58\x40\xcd\x87\x65\xa7\x6e\x9d\xaf\x51\x40\x6a\x6b\xea\x8b\xda\xe4\x8a\xee\x75\x55\xb9\x41\x79\x5f\x0d\x55\x3d\x9b\x86\xe3\xd1\xe7\xd7\xd4\x3b\x4a\x68\x6f\x5b\x41\x77\xa4\x4a\xa1\x6c\xd1\xc6\xb9\x9a\xec\xd0\xcc\xba\x31\xd2\x92\x80\x45\xdc\xf6\xbe\x2e\x51\xaa\x41\x1c\x57\xb2\x09\x4e\xa7\x60\xaf\x4b\x6f\xc9\x7f\x7a\xb3\xe5\xe2\xe3\x1e\x5d\xb4\x78\x19\xde\xac\x56\x51\x1c\xc5\x9f\x06\xbb\x1f\xc7\x81\xe6\x52\xe4\x47\x45\xef\xee\x2c\x6d\x4e\x59\xd6\x19\x49\x92\x67\x01\xdc\x07\x7b\xb4\xba\xfc\x87\x4e\xaa\x7f\xd5\x17\xaf\xe6\xb3\x59\x18\xdf\xc4\x51\x3f\xff\x19\x00\x00\xff\xff\xdc\x41\x95\xad\x2c\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 812, mode: os.FileMode(438), modTime: time.Unix(1596793158, 0)}
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
	"configs/config.yaml": configsConfigYaml,
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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
