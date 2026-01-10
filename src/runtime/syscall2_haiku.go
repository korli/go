// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import _ "unsafe" // for go:linkname

//go:cgo_import_dynamic libc_chdir chdir "libroot.so"
//go:cgo_import_dynamic libc_chroot chroot "libroot.so"
//go:cgo_import_dynamic libc_close close "libroot.so"
//go:cgo_import_dynamic libc_dup2 dup2 "libroot.so"
//go:cgo_import_dynamic libc_execve execve "libroot.so"
//go:cgo_import_dynamic libc_fcntl fcntl "libroot.so"
//go:cgo_import_dynamic libc_gethostname gethostname "libroot.so"
//go:cgo_import_dynamic libc_getpid getpid "libroot.so"
//go:cgo_import_dynamic libc_ioctl ioctl "libroot.so"
//go:cgo_import_dynamic libc_setgid setgid "libroot.so"
//go:cgo_import_dynamic libc_setgroups setgroups "libroot.so"
//go:cgo_import_dynamic libc_setsid setsid "libroot.so"
//go:cgo_import_dynamic libc_setuid setuid "libroot.so"
//go:cgo_import_dynamic libc_setpgid setpgid "libroot.so"
//go:cgo_import_dynamic libc__kern_generic_syscall _kern_generic_syscall "libroot.so"
//go:cgo_import_dynamic libc_fork fork "libroot.so"
//go:cgo_import_dynamic libc_wait4 wait4#LIBBSD_1_BETA3 "libbsd.so"

//go:linkname libc_chdir libc_chdir
//go:linkname libc_chroot libc_chroot
//go:linkname libc_close libc_close
//go:linkname libc_dup2 libc_dup2
//go:linkname libc_execve libc_execve
//go:linkname libc_fcntl libc_fcntl
//go:linkname libc_gethostname libc_gethostname
//go:linkname libc_getpid libc_getpid
//go:linkname libc_ioctl libc_ioctl
//go:linkname libc_setgid libc_setgid
//go:linkname libc_setgroups libc_setgroups
//go:linkname libc_setsid libc_setsid
//go:linkname libc_setuid libc_setuid
//go:linkname libc_setpgid libc_setpgid
//go:linkname libc__kern_generic_syscall libc__kern_generic_syscall
//go:linkname libc_fork libc_fork
//go:linkname libc_wait4 libc_wait4
