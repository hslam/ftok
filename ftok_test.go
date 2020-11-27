// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package ftok

import (
	"testing"
)

func TestFtok(t *testing.T) {
	key, err := Ftok("/tmp", 0x22)
	if err != nil {
		t.Error(err)
	}
	key1, err := Ftok("/tmp", 0x22)
	if err != nil {
		t.Error(err)
	}
	if key != key1 && key <= 0 {
		t.Error(key, key1)
	}
}

func TestStat(t *testing.T) {
	_, err := Ftok("/tmp/hslam", 0x22)
	if err == nil {
		t.Error()
	}
}
