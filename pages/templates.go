// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// generated by build/assets.sh; DO NOT EDIT

// Code generated by go-bindata.
// sources:
// pages/assets/html/containers.html
// DO NOT EDIT!

package pages

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

var _pagesAssetsHtmlContainersHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\xcd\x72\xdb\x38\x12\x3e\x4b\x4f\xd1\xc3\xda\xc3\x6c\x55\x48\xd9\x89\x2f\x9b\x95\x55\xa5\x51\x92\x1d\xed\x38\x76\xca\xb2\x67\x6a\x8e\x20\xd9\x22\x11\x43\x04\x06\x00\x25\x6b\x5d\x7e\xf7\x2d\x00\xa4\xc4\x5f\x29\xfe\xa9\x64\x74\xb1\x44\xa0\xbb\xbf\xfe\xba\x1b\x68\x10\x1e\xff\xe4\xfb\x43\x80\x19\x17\x5b\x49\x93\x54\xc3\xdb\x93\xd3\x33\xf8\x0f\xe7\x09\x43\x98\x67\x51\x00\x53\xc6\xe0\xda\x0c\x29\xb8\x46\x85\x72\x8d\x71\x30\x1c\x02\x5c\xd0\x08\x33\x85\x31\xe4\x59\x8c\x12\x74\x8a\x30\x15\x24\x4a\xb1\x1c\x79\x03\xbf\xa3\x54\x94\x67\xf0\x36\x38\x81\x9f\xcd\x04\xaf\x18\xf2\xfe\xf9\xef\x21\xc0\x96\xe7\xb0\x22\x5b\xc8\xb8\x86\x5c\x21\xe8\x94\x2a\x58\x52\x86\x80\xf7\x11\x0a\x0d\x34\x83\x88\xaf\x04\xa3\x24\x8b\x10\x36\x54\xa7\xd6\x4c\xa1\x24\x18\x02\xfc\x59\xa8\xe0\xa1\x26\x34\x03\x02\x11\x17\x5b\xe0\xcb\xea\x3c\x20\xda\xe0\x35\x9f\x54\x6b\xf1\x7e\x34\xda\x6c\x36\x01\xb1\x58\x03\x2e\x93\x11\x73\xf3\xd4\xe8\x62\x3e\xfb\x78\xb9\xf8\xe8\xbf\x0d\x4e\x8c\xc4\x6d\xc6\x50\x29\x90\xf8\x57\x4e\x25\xc6\x10\x6e\x81\x08\xc1\x68\x44\x42\x86\xc0\xc8\x06\xb8\x04\x92\x48\xc4\x18\x34\x37\x68\x37\x92\x6a\x9a\x25\x6f\x40\xf1\xa5\xde\x10\x89\x43\x80\x98\x2a\x2d\x69\x98\xeb\x1a\x55\x25\x36\xaa\x6a\x13\x78\x06\x24\x03\x6f\xba\x80\xf9\xc2\x83\x5f\xa6\x8b\xf9\xe2\xcd\x10\xe0\x8f\xf9\xcd\xaf\x57\xb7\x37\xf0\xc7\xf4\xfa\x7a\x7a\x79\x33\xff\xb8\x80\xab\x6b\x98\x5d\x5d\x7e\x98\xdf\xcc\xaf\x2e\x17\x70\xf5\x09\xa6\x97\x7f\xc2\x6f\xf3\xcb\x0f\x6f\x00\xa9\x4e\x51\x02\xde\x0b\x69\xf0\x73\x09\xd4\x90\x68\xe2\x06\xb0\x40\xac\x01\x58\x72\x07\x48\x09\x8c\xe8\x92\x46\xc0\x48\x96\xe4\x24\x41\x48\xf8\x1a\x65\x46\xb3\x04\x04\xca\x15\x55\x26\x94\x0a\x48\x16\x0f\x01\x18\x5d\x51\x4d\xb4\x7d\xd2\x72\x2a\x18\xfa\xfe\x64\x38\x1c\xa7\x7a\xc5\x26\x43\x80\x71\x8a\x24\x9e\xd8\x10\x8c\x35\xd5\x0c\x27\xd1\x34\x5e\x53\xc5\x25\xf8\xf0\xf0\x10\x7c\xa0\x4a\x30\xb2\xbd\x24\x2b\x7c\x7c\x1c\x8f\xdc\x14\x37\x5d\x45\x92\x0a\x0d\x4a\x46\xe7\xde\xc3\x43\x70\xcd\xb9\x7e\x7c\x54\xc6\x72\x34\x12\x5c\x08\x94\xc1\x8a\x66\xc1\x57\xe5\x4d\xc6\x23\x37\xb9\x90\xfc\xc9\xf7\xe1\x82\x68\x54\xda\xe6\x10\x65\x18\x1b\xec\xb0\xa2\x19\x5d\x52\x8c\x61\xb6\x58\x80\xc1\x69\x67\x33\x9a\xdd\x81\x44\x76\xee\x29\xbd\x65\xa8\x52\x44\xed\x41\x2a\x71\xd9\xb6\x1b\x72\xae\x95\x96\x44\xf8\x67\xc1\x49\x70\xe2\x87\xa8\x49\xf0\xd6\xe2\x88\x94\xf2\x26\xc3\x3d\x80\x2b\x61\x28\x22\xcc\xb0\xb3\xc2\x97\x9a\xb3\x4a\xfc\x77\xc1\x69\x70\xda\xb2\xf6\x14\x8d\x11\xcf\x4c\xb5\xa0\x54\x2d\xc0\x07\x19\xfb\x2f\x59\x93\x85\x0b\xc8\xce\x93\x43\x01\xfa\xfa\x57\x8e\x72\xeb\xbf\x33\x2c\xf5\x85\xe9\x90\xfc\x01\xa2\xfb\x35\xe9\xad\xc0\x73\x4f\xe3\xbd\x1e\x7d\x25\x6b\xe2\x9e\x7a\xdd\x06\x12\xbb\xcc\xf9\x5f\x15\x11\xb4\xa1\xf2\xd9\x3a\x2b\xe4\xbe\x12\xc8\x28\x25\x52\xb7\xb5\x8d\x47\x65\x59\x8d\x43\x1e\x6f\x0b\x03\x31\x5d\x43\xc4\x88\x52\xe7\xde\x0e\x89\xcb\x3e\x5f\xa5\x7c\x13\x11\x85\x1e\x4c\x8a\xe5\x70\x4c\x9a\x19\xe2\xed\x85\x99\xaf\x56\xfe\xe9\x5b\x0f\x68\x7c\xee\x31\x9e\x70\x6f\x27\x36\x22\xbb\xaf\x35\x7b\xa5\xc8\x64\x38\xa8\x0e\x08\x92\xa0\x6f\xc0\xa2\x34\x43\x66\x41\x38\x9d\xb4\xeb\x3e\x3d\x35\x72\xa3\x98\xae\xcd\x5f\xce\x4a\xf1\x50\x22\x89\x23\x99\xaf\x42\x27\xfd\xf0\x20\x49\x96\x20\xfc\x43\x10\x89\x99\x9e\xed\xdc\x7c\x7f\x0e\xc1\x97\xfa\x33\xf5\xf8\x68\x0d\x32\x3a\xa9\x38\xdb\x94\x0c\x2e\x68\x76\xf7\xf8\xe8\x4d\x3a\x86\x6e\xf0\x5e\x1b\x74\x64\x32\x1e\x31\x5a\x00\xc0\x2c\x36\x8a\xc7\x23\xce\xf6\xa4\x58\xe0\xee\xc7\xc3\x03\x5d\x42\x30\x57\x8e\xd4\x23\x5c\x41\xf1\x19\xa7\x67\x7b\x90\x41\x30\x8a\x79\x74\x67\x18\xfb\x60\xff\xc2\xde\x27\x07\x26\x3d\xeb\x31\xed\xc0\x55\x81\x2c\xf2\x30\xaa\x32\xf2\xb2\xd8\xbd\x9b\xd4\xf4\x8d\x47\xe9\xbb\x6a\xe0\x2a\xc2\x8c\x2a\xed\x27\x92\xe7\xa2\x11\x39\x55\x51\x60\xc3\xd6\x44\x38\xa8\x25\x67\x6d\x7e\x19\xac\xb6\x11\x9f\x6a\x5c\xd9\x20\xd6\xe6\xef\x23\xd8\x08\x5e\x85\xb5\x7e\x0a\x1d\x83\x2e\x06\x0b\x4d\x74\xfe\x1a\x04\x7e\x90\x74\x8d\x12\x9c\xbe\x26\x81\x39\x3b\xca\x9f\x4b\x0d\x65\xc5\x2d\x7f\x0d\x7c\x2e\xe5\x9d\x1a\xe8\xa0\x68\xac\x04\xc9\x4a\x2b\x46\x8d\xcf\x48\x88\xcc\x72\x57\xd5\x1d\xfc\x86\x5b\x43\x9d\x99\x3e\x81\xe6\xe0\xef\x84\xe5\xb6\x72\x9b\x75\x51\x67\xcd\x39\xbb\xc7\x36\x78\x1e\xb4\x85\xe6\x92\x24\x38\x0e\xe5\xa4\x00\x64\x54\xf5\x91\x35\xd8\x73\x65\xcd\xb7\xb8\xea\x47\xf5\x54\xbe\x2a\xfa\xdb\x7c\x55\x07\xeb\x7c\x0d\x76\x74\x0d\xc6\xa3\x9c\x59\x6f\x4a\x26\x8b\x07\x7d\xd9\xda\x55\xe3\xce\xab\xf9\x8a\x24\x78\x3c\x43\x61\xf7\xe9\x4f\x55\xa8\x7c\x4c\xce\x3a\xd5\x2e\x59\x2b\x23\x55\x5c\x4e\x9b\xd9\x2f\x5c\x9e\xf8\xd4\xca\x98\x7d\xab\x36\xcb\x84\x30\x94\xfb\xdf\xc7\x7c\xbb\x46\xc5\x73\x19\xa1\x9a\xae\x09\x65\xa6\xfb\x7e\x85\x1a\x9c\x2b\xce\x6c\x07\xdb\xa8\x3f\x67\x72\x26\xf2\xaa\xb1\xde\x44\xab\x30\xd1\x9b\x3f\x40\x22\x4d\xd7\xa6\xd7\x2f\x2c\xfa\xb6\xc5\x05\x41\x32\x64\xee\xbb\x37\x99\x7d\xb9\x75\xe1\xdf\x6b\x2c\x16\x6f\x81\x91\x81\x13\x5c\x98\x9e\x7b\xe7\xf8\x61\x93\x87\xea\x28\x25\xd2\xc4\xb1\xcc\x51\x21\x69\xa6\xdd\xc3\xb6\x31\xa8\xa9\xc9\x33\xba\x53\xa3\xaa\x6a\xda\xc8\xab\x41\xec\xf0\xe5\x33\xb9\x7f\x25\x77\x3e\x93\x7b\xb0\xaa\x1a\x1e\xcd\x78\xdd\xa1\xbd\xc5\x7e\x9f\x22\xfe\x22\x97\xd4\xdd\xcb\xdd\x99\x32\xc6\x37\xe6\x74\xc2\xdb\x41\x32\x16\x1a\x06\x21\xf8\x4c\xa2\x94\x66\x38\xcf\x96\x3c\xb8\xcc\x57\x56\xae\x5c\x63\xda\xe8\xcb\xa5\x66\xf7\xdb\x39\xf1\x19\x57\x5c\x6e\xbf\x6f\xc2\x3b\x9b\x07\x72\xde\x4d\x08\xdc\x4b\x07\xab\xe6\xe5\xf4\x56\x94\x35\x2b\x80\xfe\x0f\x0f\x18\xee\x4f\x9a\x42\xfe\x36\xa3\xfa\x80\xfc\x73\xb2\xaa\xd0\xf3\x4a\x85\xd2\x55\x24\x6d\xa7\x8f\xd6\x48\xaf\xbb\x85\xe4\x0b\x1c\x5d\x6c\x88\x78\xad\x45\x6e\x43\x44\xe7\xb2\xd0\xf6\xb8\x62\xf5\x19\x5e\x57\xa4\x8f\x78\xde\x2c\xbd\xc2\xbb\x5a\x17\xfa\xec\xcd\xec\x56\x99\xd6\xa8\xbf\x13\xb7\x95\x57\xd4\x9f\x90\x74\x45\xe4\xf6\x40\x1b\x60\x66\x19\x0b\x34\x4b\xda\x8d\x40\x7d\x5a\x51\xcc\x57\x6b\x94\x6b\x8a\x9b\xc3\xed\x41\xb5\x43\xc8\x0d\x62\x3f\x21\x79\x82\x5e\x5d\xa5\x39\xcd\xee\x5a\x86\x1f\xe2\xcd\x17\xc9\x23\x54\xea\x58\xb7\x53\x75\x47\x94\x22\xbe\xe6\xe2\x9b\x1c\xea\xe9\x33\xbe\xa3\x9b\xb6\xe5\xf8\x16\x07\x3b\xbc\x69\x18\x38\x9b\xdc\x70\x4d\x18\x94\x79\x78\x66\x33\xb3\xc2\x4f\x24\x72\x5f\x9b\x29\xbe\x0b\xbc\x7d\xa9\xb1\x27\x05\xca\x17\x50\x46\xd5\xec\xcb\x2d\x5c\x70\x12\xc3\x74\x8d\xf2\x80\x3e\xc6\x49\x5c\x57\xb4\x7b\x2f\x55\x45\x66\x31\x81\xb0\x47\x68\xd9\xab\x4c\xa0\xf4\xcd\xfe\xdf\x89\xaf\x5b\xe5\x2f\x12\xc9\x5d\xcc\x37\x59\x9f\x4e\xa7\x2a\x2c\xa7\xf5\x2a\x6d\xa7\xc6\xd1\xdd\xf9\x3b\xa6\x49\xb9\x51\x7f\xa7\x4c\x59\x59\x73\xc7\xc3\x10\xca\x51\xe3\x49\x05\x80\xe4\x1b\xe8\x3e\xf0\x1c\x0c\x61\x63\x5a\x7b\x39\xfe\x97\x3d\x5b\xd6\x5c\x95\x3c\x91\x68\x5f\xa3\x42\xeb\xd3\x35\xd1\x0f\x89\x84\xea\x0f\x3f\x36\x07\x55\xe9\x95\xeb\x88\x1b\x48\xb9\xf6\x1d\x15\x9d\x9a\xa1\xbe\x57\x29\xe9\xf3\x8c\x6d\xbd\xc9\xaf\x5c\x43\x19\x30\x77\x48\xee\x90\x6c\xb3\xf9\x14\xb8\x34\x5b\xf2\x06\xd8\x88\xb3\xf8\x39\x68\x67\x9c\xc5\xdf\x0a\x77\x30\xe8\xc4\xdd\xfd\xb0\x1d\xb9\x77\x5e\x35\xbb\x34\xde\x37\x57\x9f\x27\x16\xe5\x25\xea\x0d\x97\x77\x4f\xac\xca\xc1\xcb\xcb\xb1\x30\x5c\x6c\xf6\x4f\x29\xc4\x41\x73\x34\x96\x5c\x98\xe4\x6f\x17\x48\x98\x6b\xcd\x77\xf1\x0a\x75\x06\xa1\xce\xfc\x18\x97\x24\x67\x1a\x4a\x39\x5f\xf3\x24\x61\xe8\x15\xef\xb3\x9d\x90\xe3\x39\x73\x28\x7d\x85\x0c\x23\x7b\x04\xd8\x19\x83\x98\x68\x52\x88\x56\x30\x00\x91\x94\xf8\x29\x51\x82\x8b\x5c\x9c\x7b\x5a\xe6\x58\x3c\xc4\x7b\x41\xb2\x18\xe3\x73\x6f\x49\x98\xc2\x8e\x14\x73\xe9\xd5\x6d\xb8\x8c\x75\x77\x7e\xd5\x12\x33\x22\x12\x2b\x73\x07\x65\x26\x38\xcf\x5a\x2c\xe5\xac\xdb\xa4\xd7\x24\xd8\x5f\x61\x96\x7b\x20\xb9\xf1\xd8\x7d\xb7\x8e\xd9\xee\x92\x61\x1c\x6e\x0f\x32\xd6\xce\xf9\xe2\xf5\xd0\x81\xb4\x7d\xca\x82\x9c\x4a\x9e\x27\xa9\xc8\x75\x7b\x15\xdc\x2d\xcb\x25\xbc\x70\xab\x51\xb5\xb7\xef\x67\x98\xfd\x28\x25\xb7\xaf\x8f\x5b\x5b\x40\x69\x0b\xed\x8c\x7e\x63\x0d\xe7\x1b\x15\xfa\x49\xfd\xb0\x2d\xf3\x13\x65\xa8\xb6\x4a\xe3\xea\xdb\x3b\xc8\xe5\x4e\xc6\xed\x7d\x9d\x4d\x64\xbf\xa6\x9e\x65\x6a\x96\x2b\xcd\x57\x9f\x51\x4b\x1a\x3d\x95\x8f\x23\x8b\xd5\xe0\x10\x03\x53\x77\x51\x6e\xf2\x18\x0a\xeb\xcd\x15\xeb\x50\xae\x34\x7a\x29\xeb\x84\xbf\x72\x7a\x8e\xe6\xc3\xa0\x79\xd8\xec\xb8\x05\xf9\x61\xa9\xd1\x71\x77\x72\x2c\x3b\xbe\xad\xa9\x12\x60\xfa\x66\xdb\xd6\xbc\x6f\xae\x17\x34\x13\xb9\xae\xb5\xba\xd5\x1b\x12\x3f\x76\x17\x71\x7e\xc4\xf3\x4c\x7b\x9d\xfb\xf7\x6e\xeb\xee\x92\xb3\xea\x7b\xe4\xd6\x84\xe5\x78\x7e\x7a\xd2\x80\xdc\xbf\xd0\x74\x22\xac\x75\x83\x0d\x4d\xdd\x0b\xe0\x33\x39\x74\xcd\xc8\x51\x1a\x8b\x36\xe2\xef\xc9\x64\xad\xd5\x72\x56\x24\x67\xac\x62\x26\x64\x3c\xba\x6b\x32\xd0\xde\x1f\x9b\x3d\xf9\x2b\x86\xa5\x67\xe9\xee\x18\xac\x0e\x55\x06\x0e\x5f\xa5\x97\xc2\x4a\x13\xa9\xbf\x90\x04\x7f\x7e\x78\x08\x76\x37\xa8\xee\xc6\xf9\x0d\x98\x67\xb5\xf3\xb7\x7d\xd4\x3a\x6e\xd9\xa7\xee\x2a\xd7\x7e\x2d\xef\x75\xed\x3f\x31\x99\x4f\x2c\xc9\xc6\x5d\x8f\x18\x33\xf5\x9b\x98\x62\x52\xfd\xe6\xde\x5d\xd8\x8f\x47\xee\x3f\x64\xfe\x1f\x00\x00\xff\xff\x17\x71\xfc\x09\x84\x25\x00\x00")

func pagesAssetsHtmlContainersHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pagesAssetsHtmlContainersHtml,
		"pages/assets/html/containers.html",
	)
}

func pagesAssetsHtmlContainersHtml() (*asset, error) {
	bytes, err := pagesAssetsHtmlContainersHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pages/assets/html/containers.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"pages/assets/html/containers.html": pagesAssetsHtmlContainersHtml,
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
	"pages": {nil, map[string]*bintree{
		"assets": {nil, map[string]*bintree{
			"html": {nil, map[string]*bintree{
				"containers.html": {pagesAssetsHtmlContainersHtml, map[string]*bintree{}},
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
