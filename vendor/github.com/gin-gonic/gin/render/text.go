// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"fmt"
<<<<<<< HEAD
	"net/http"

	"github.com/gin-gonic/gin/internal/bytesconv"
=======
	"io"
	"net/http"
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
)

// String contains the given interface object slice and its format.
type String struct {
	Format string
	Data   []interface{}
}

var plainContentType = []string{"text/plain; charset=utf-8"}

// Render (String) writes data with custom ContentType.
func (r String) Render(w http.ResponseWriter) error {
	return WriteString(w, r.Format, r.Data)
}

// WriteContentType (String) writes Plain ContentType.
func (r String) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, plainContentType)
}

// WriteString writes data according to its format and write custom ContentType.
func WriteString(w http.ResponseWriter, format string, data []interface{}) (err error) {
	writeContentType(w, plainContentType)
	if len(data) > 0 {
		_, err = fmt.Fprintf(w, format, data...)
		return
	}
<<<<<<< HEAD
	_, err = w.Write(bytesconv.StringToBytes(format))
=======
	_, err = io.WriteString(w, format)
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
	return
}
