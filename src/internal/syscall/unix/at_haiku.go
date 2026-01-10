// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import "syscall"

// Implemented as sysvicall6 in runtime/syscall_haiku.go.
func syscall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

//go:cgo_import_dynamic libc_fstatat fstatat "libroot.so"
//go:cgo_import_dynamic libc_openat openat "libroot.so"
//go:cgo_import_dynamic libc_unlinkat unlinkat "libroot.so"

const (
	AT_REMOVEDIR        = 0x4
	AT_SYMLINK_NOFOLLOW = 0x1
)
