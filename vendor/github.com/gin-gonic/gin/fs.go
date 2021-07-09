// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"net/http"
	"os"
)

<<<<<<< HEAD
type onlyFilesFS struct {
=======
type onlyfilesFS struct {
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
	fs http.FileSystem
}

type neuteredReaddirFile struct {
	http.File
}

// Dir returns a http.Filesystem that can be used by http.FileServer(). It is used internally
// in router.Static().
// if listDirectory == true, then it works the same as http.Dir() otherwise it returns
// a filesystem that prevents http.FileServer() to list the directory files.
func Dir(root string, listDirectory bool) http.FileSystem {
	fs := http.Dir(root)
	if listDirectory {
		return fs
	}
<<<<<<< HEAD
	return &onlyFilesFS{fs}
}

// Open conforms to http.Filesystem.
func (fs onlyFilesFS) Open(name string) (http.File, error) {
=======
	return &onlyfilesFS{fs}
}

// Open conforms to http.Filesystem.
func (fs onlyfilesFS) Open(name string) (http.File, error) {
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

// Readdir overrides the http.File default implementation.
func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	// this disables directory listing
	return nil, nil
}
