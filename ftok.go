// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package ftok provides a way to generate a System V IPC key, suitable for using with
// msgget, semget, or shmget.
package ftok

import (
	"syscall"
)

// Ftok uses the given pathname (which must refer to an existing, accessible file) and
// the least significant 8 bits of proj_id (which must be nonzero) to generate
// a key_t type System V IPC key.
func Ftok(pathname string, projectid uint8) (int, error) {
	var stat = syscall.Stat_t{}
	if err := syscall.Stat(pathname, &stat); err != nil {
		return 0, err
	}
	return int(uint(projectid&0xff)<<24 | uint((stat.Dev&0xff)<<16) | (uint(stat.Ino) & 0xffff)), nil
}
