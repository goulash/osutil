// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"fmt"
	"os"
)

// IsExists returns ok = true if the given file exists, regardless whether it
// is a file or a directory. Normally you will probably want to use the more
// specific versions: IsFileExists and IsDirectoryExists.
func IsExists(path string) (ok bool, err error) {
	ok, _, err = isExists(path)
	return
}

// IsFileExists returns ok = true if the file exists and is not a directory,
// and returns err != nil if any other error occured (such as permission
// denied).
func IsFileExists(path string) (ok bool, err error) {
	var stat os.FileInfo

	ok, stat, err = isExists(path)
	if err != nil && stat.IsDir() {
		err = errors.New(fmt.Sprintf("%s exists but is a directory not a file", path))
	}
	return
}

// IsDirectoryExists returns ok = true if the file exists and is a directory,
// and returns err != nil if any other error occured (such as permission
// denied).
func IsDirectoryExists(path string) (ok bool, err error) {
	var stat os.FileInfo

	ok, stat, err = isExists(path)
	if err != nil && !stat.IsDir() {
		err = errors.New(fmt.Sprintf("%s exists but is not a directory", path))
	}
	return
}

// isExists does the hard work for IsExist, IsFileExist, and IsDirectoryExists,
// returning ok = true if the file given by path exists.
func isExists(path string) (ok bool, stat os.FileInfo, err error) {
	stat, err = os.Stat(path)

	ok = true
	if err != nil {
		// ok = true if file exists
		ok = !os.IsNotExist(err)
		if !ok {
			err = nil
		}
	}
	return
}
