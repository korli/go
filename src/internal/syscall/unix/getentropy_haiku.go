// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build haiku

package unix

import (
	"unsafe"
)

//go:cgo_import_dynamic libc_getentropy getentropy "libroot.so"
//go:linkname libc_getentropy libc_getentropy

var libc_getentropy uintptr

type GetRandomFlag uintptr

// GetEntropy calls the Haiku getentropy system call.
func GetRandom(p []byte, flags GetRandomFlag) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&libc_getentropy)),
		2,
		uintptr(unsafe.Pointer(&p[0])),
		uintptr(len(p)),
		0, 0, 0, 0)
	if errno != 0 {
		return 0, errno
	}
	return len(p), nil
}

