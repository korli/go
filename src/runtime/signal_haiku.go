// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

var sigtable = [...]sigTabT{
	/*  0 */ {0, "SIGNONE: no trap"},
	/*  1 */ {_SigNotify + _SigKill, "SIGHUP: terminal line hangup"},
	/*  2 */ {_SigNotify + _SigKill, "SIGINT: interrupt"},
	/*  3 */ {_SigNotify + _SigThrow, "SIGQUIT: quit"},
	/*  4 */ {_SigThrow + _SigUnblock, "SIGILL: illegal instruction"},
	/*  5 */ {_SigNotify + _SigUnblock, "SIGCHLD: child status has changed"},
	/*  6 */ {_SigNotify + _SigThrow, "SIGABRT: abort"},
	/*  7 */ {_SigNotify, "SIGPIPE: write to broken pipe"},
	/*  8 */ {_SigPanic + _SigUnblock, "SIGFPE: floating-point exception"},
	/*  9 */ {0, "SIGKILL: kill"},
	/* 10 */ {0, "SIGSTOP: stop"},
	/* 11 */ {_SigPanic + _SigUnblock, "SIGSEGV: segmentation violation"},
	/* 12 */ {_SigNotify + _SigDefault, "SIGCONT: continue after stop"},
	/* 13 */ {_SigNotify + _SigDefault, "SIGTSTP: keyboard stop"},
	/* 14 */ {_SigNotify, "SIGALRM: alarm clock"},
	/* 15 */ {_SigNotify + _SigKill, "SIGTERM: termination"},
	/* 16 */ {_SigNotify + _SigDefault, "SIGTTIN: background read from tty"},
	/* 17 */ {_SigNotify + _SigDefault, "SIGTTOU: background write to tty"},
	/* 18 */ {_SigNotify, "SIGUSR1: user-defined signal 1"},
	/* 19 */ {_SigNotify, "SIGUSR2: user-defined signal 2"},
	/* 20 */ {_SigNotify, "SIGWINCH: window size change"},
	/* 21 */ {_SigNotify + _SigKill, "SIGKILLTHR: kill thread"},
	/* 22 */ {_SigThrow + _SigUnblock, "SIGTRAP: trace trap"},
	/* 23 */ {_SigNotify, "SIGPOLL: poll"},
	/* 24 */ {_SigNotify + _SigUnblock, "SIGPROF: profiling alarm clock"},
	/* 25 */ {_SigThrow, "SIGSYS: bad system call"},
	/* 26 */ {_SigNotify, "SIGURG: urgent condition on socket"},
	/* 27 */ {_SigNotify, "SIGVTALRM: virtual alarm clock"},
	/* 28 */ {_SigNotify, "SIGXCPU: cpu limit exceeded"},
	/* 29 */ {_SigNotify, "SIGXFSZ: file size limit exceeded"},
	/* 30 */ {_SigPanic + _SigUnblock, "SIGBUS: bus error"},
	/* 31 */ {_SigNotify, "SIGRESERVED1: reserved1"},
	/* 32 */ {_SigNotify, "SIGRESERVED2: reserved2"},
}
