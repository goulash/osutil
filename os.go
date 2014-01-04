// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"fmt"
	"os"
)

// IsExists returns exists = true if the given file exists, regardless whether
// it is a file or a directory. Normally you will probably want to use the more
// specific versions: IsFileExists and IsDirectoryExists.
func IsExists(path string) (exists bool, err error) {
	exists, _, err = isExists(path)
	return
}

// IsFileExists returns exists = true if the file exists and is not
// a directory, and returns err != nil if any other error occured (such as
// permission denied).
func IsFileExists(path string) (exists bool, err error) {
	var stat os.FileInfo

	exists, stat, err = isExists(path)
	if err != nil && stat.IsDir() {
		err = errors.New(fmt.Sprintf("%s exists but is a directory not a file", path))
	}
	return
}

// IsDirectoryExists returns exists = true if the file exists and is
// a directory, and returns err != nil if any other error occured (such as
// permission denied).
func IsDirectoryExists(path string) (exists bool, err error) {
	var stat os.FileInfo

	exists, stat, err = isExists(path)
	if err != nil && !stat.IsDir() {
		err = errors.New(fmt.Sprintf("%s exists but is not a directory", path))
	}
	return
}

// isExists does the hard work for IsExist, IsFileExist, and IsDirectoryExists,
// returning exists = true if the file given by path exists.
func isExists(path string) (exists bool, stat os.FileInfo, err error) {
	stat, err = os.Stat(path)

	exists = true
	if err != nil {
		// exists = true if file exists
		exists = !os.IsNotExist(err)
		if !exists {
			err = nil
		}
	}
	return
}
