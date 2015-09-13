// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#pragma dynimport libc·chdir chdir "libroot.so"
#pragma dynimport libc·chroot chroot "libroot.so"
#pragma dynimport libc·close close "libroot.so"
#pragma dynimport libc·dlclose dlclose "libroot.so"
#pragma dynimport libc·dlopen dlopen "libroot.so"
#pragma dynimport libc·dlsym dlsym "libroot.so"
#pragma dynimport libc·execve execve "libroot.so"
#pragma dynimport libc·fcntl fcntl "libroot.so"
#pragma dynimport libc·gethostname gethostname "libnetwork.so"
#pragma dynimport libc·ioctl ioctl "libroot.so"
#pragma dynimport libc·pipe pipe "libroot.so"
#pragma dynimport libc·setgid setgid "libroot.so"
#pragma dynimport libc·setgroups setgroups "libroot.so"
#pragma dynimport libc·setsid setsid "libroot.so"
#pragma dynimport libc·setuid setuid "libroot.so"
#pragma dynimport libc·setpgid setsid "libroot.so"
#pragma dynimport libc·fork fork "libroot.so"
#pragma dynimport libc·waitpid waitpid "libroot.so"
