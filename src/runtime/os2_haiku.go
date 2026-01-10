// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/abi"
	"runtime/internal/atomic"
	"unsafe"
)

// Remove
//libc_kill,
//libc_madvise,

//go:cgo_export_dynamic runtime.end _end
//go:cgo_export_dynamic runtime.etext _etext
//go:cgo_export_dynamic runtime.edata _edata

//go:cgo_import_dynamic libc__errnop _errnop "libroot.so"
//go:cgo_import_dynamic libc_clock_gettime clock_gettime "libroot.so"
//go:cgo_import_dynamic libc_exit exit "libroot.so"
//go:cgo_import_dynamic libc_fstat fstat#LIBROOT_1_ALPHA1 "libroot.so"
//go:cgo_import_dynamic libc_getcontext getcontext "libroot.so"
//go:cgo_import_dynamic libc_kill kill "libroot.so"
//go:cgo_import_dynamic libc_getrlimit getrlimit "libroot.so"
//go:cgo_import_dynamic libc_malloc malloc "libroot.so"
//go:cgo_import_dynamic libc_mmap mmap "libroot.so"
//go:cgo_import_dynamic libc_munmap munmap "libroot.so"
//go:cgo_import_dynamic libc_open open "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_destroy pthread_attr_destroy "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_getstacksize pthread_attr_getstacksize "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_init pthread_attr_init "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_setdetachstate pthread_attr_setdetachstate "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_setstack pthread_attr_setstack "libroot.so"
//go:cgo_import_dynamic libc_pthread_create pthread_create "libroot.so"
//go:cgo_import_dynamic libc_pthread_self pthread_self "libroot.so"
//go:cgo_import_dynamic libc_pthread_kill pthread_kill "libroot.so"
//go:cgo_import_dynamic libc_raise raise "libroot.so"
//go:cgo_import_dynamic libc_read read "libroot.so"
//go:cgo_import_dynamic libc_select select "libroot.so"
//go:cgo_import_dynamic libc_sched_yield sched_yield "libroot.so"
//go:cgo_import_dynamic libc_sem_init sem_init "libroot.so"
//go:cgo_import_dynamic libc_sem_post sem_post "libroot.so"
//go:cgo_import_dynamic libc_sem_timedwait sem_timedwait "libroot.so"
//go:cgo_import_dynamic libc_sem_wait sem_wait "libroot.so"
//go:cgo_import_dynamic libc_setitimer setitimer "libroot.so"
//go:cgo_import_dynamic libc_sigaction sigaction#LIBROOT_1_ALPHA4 "libroot.so"
//go:cgo_import_dynamic libc_sigaltstack sigaltstack "libroot.so"
//go:cgo_import_dynamic libc_sigprocmask sigprocmask#LIBROOT_1_ALPHA4 "libroot.so"
//go:cgo_import_dynamic libc_sysconf sysconf#LIBROOT_1_ALPHA4 "libroot.so"
//go:cgo_import_dynamic libc_usleep usleep "libroot.so"
//go:cgo_import_dynamic libc_write write "libroot.so"
//go:cgo_import_dynamic libc_pipe pipe "libroot.so"
//go:cgo_import_dynamic libc_pipe2 pipe2 "libroot.so"
//go:cgo_import_dynamic libc_area_for area_for "libroot.so"

//go:linkname libc__errnop libc__errnop
//go:linkname libc_clock_gettime libc_clock_gettime
//go:linkname libc_exit libc_exit
//go:linkname libc_fstat libc_fstat
//go:linkname libc_getcontext libc_getcontext
//go:linkname libc_kill libc_kill
//go:linkname libc_getrlimit libc_getrlimit
//go:linkname libc_malloc libc_malloc
//go:linkname libc_mmap libc_mmap
//go:linkname libc_munmap libc_munmap
//go:linkname libc_open libc_open
//go:linkname libc_pthread_attr_destroy libc_pthread_attr_destroy
//go:linkname libc_pthread_attr_getstacksize libc_pthread_attr_getstacksize
//go:linkname libc_pthread_attr_init libc_pthread_attr_init
//go:linkname libc_pthread_attr_setdetachstate libc_pthread_attr_setdetachstate
//go:linkname libc_pthread_attr_setstack libc_pthread_attr_setstack
//go:linkname libc_pthread_create libc_pthread_create
//go:linkname libc_pthread_self libc_pthread_self
//go:linkname libc_pthread_kill libc_pthread_kill
//go:linkname libc_raise libc_raise
//go:linkname libc_read libc_read
//go:linkname libc_select libc_select
//go:linkname libc_sched_yield libc_sched_yield
//go:linkname libc_sem_init libc_sem_init
//go:linkname libc_sem_post libc_sem_post
//go:linkname libc_sem_timedwait libc_sem_timedwait
//go:linkname libc_sem_wait libc_sem_wait
//go:linkname libc_setitimer libc_setitimer
//go:linkname libc_sigaction libc_sigaction
//go:linkname libc_sigaltstack libc_sigaltstack
//go:linkname libc_sigprocmask libc_sigprocmask
//go:linkname libc_sysconf libc_sysconf
//go:linkname libc_usleep libc_usleep
//go:linkname libc_write libc_write
//go:linkname libc_pipe libc_pipe
//go:linkname libc_pipe2 libc_pipe2
//go:linkname libc_area_for libc_area_for

var (
	libc__errnop,
	libc_clock_gettime,
	libc_exit,
	libc_fstat,
	libc_getcontext,
	libc_getrlimit,
	libc_kill,
	libc_malloc,
	libc_mmap,
	libc_munmap,
	libc_open,
	libc_pthread_attr_destroy,
	libc_pthread_attr_getstacksize,
	libc_pthread_attr_init,
	libc_pthread_attr_setdetachstate,
	libc_pthread_attr_setstack,
	libc_pthread_create,
	libc_pthread_self,
	libc_pthread_kill,
	libc_raise,
	libc_read,
	libc_sched_yield,
	libc_select,
	libc_sem_init,
	libc_sem_post,
	libc_sem_timedwait,
	libc_sem_wait,
	libc_setitimer,
	libc_sigaction,
	libc_sigaltstack,
	libc_sigprocmask,
	libc_sysconf,
	libc_usleep,
	libc_write,
	libc_area_for,
	libc_pipe,
	libc_pipe2 libcFunc
)

const (
	_CLOCK_REALTIME  = 0xffffffff
	_CLOCK_MONOTONIC = 0
)

var sigset_all = ^sigset(0)
var sigset_none = sigset(0)

func getncpu() int32 {
	n := int32(sysconf(__SC_NPROCESSORS_ONLN))
	if n < 1 {
		return 1
	}
	return n
}

func getPageSize() uintptr {
	n := int32(sysconf(__SC_PAGESIZE))
	if n <= 0 {
		return 0
	}
	return uintptr(n)
}

func osinit() {
	ncpu = getncpu()
	if physPageSize == 0 {
		physPageSize = getPageSize()
	}
}

func tstart_sysvicall(newm *m) uint32

// May run with m.p==nil, so write barriers are not allowed.
//go:nowritebarrier
func newosproc(mp *m) {
	var (
		attr pthreadattr
		oset sigset
		tid  pthread
		ret  int32
	)

	if pthread_attr_init(&attr) != 0 {
		throw("pthread_attr_init")
	}

	if pthread_attr_setdetachstate(&attr, _PTHREAD_CREATE_DETACHED) != 0 {
		throw("pthread_attr_setdetachstate")
	}

	// Disable signals during create, so that the new thread starts
	// with signals disabled. It will enable them in minit.
	sigprocmask(_SIG_SETMASK, &sigset_all, &oset)
	ret = pthread_create(&tid, &attr, abi.FuncPCABI0(tstart_sysvicall), unsafe.Pointer(mp))
	sigprocmask(_SIG_SETMASK, &oset, nil)
	if ret != 0 {
		print("runtime: failed to create new OS thread (have ", mcount(), " already; errno=", ret, ")\n")
		if ret == -_EAGAIN {
			println("runtime: may need to increase max user processes (ulimit -u)")
		}
		throw("newosproc")
	}
}


var urandom_dev = []byte("/dev/random\x00")

//go:nosplit
func getRandomData(r []byte) {
	fd := open(&urandom_dev[0], 0 /* O_RDONLY */, 0)
	n := read(fd, unsafe.Pointer(&r[0]), int32(len(r)))
	closefd(fd)
	extendRandom(r, int(n))
}

func goenvs() {
	goenvs_unix()
}

// Called to initialize a new m (including the bootstrap m).
// Called on the parent thread (main thread in case of bootstrap), can allocate memory.
func mpreinit(mp *m) {
	mp.gsignal = malg(32 * 1024)
	mp.gsignal.m = mp
}

func miniterrno()

//go:nosplit
func setSignalstackSP(s *stackt, sp uintptr) {
	*(*uintptr)(unsafe.Pointer(&s.ss_sp)) = sp
}

//go:nosplit
//go:nowritebarrierrec
func sigaddset(mask *sigset, i int) {
	*mask |= 1 << (sigset(i) - 1)
}

func sigdelset(mask *sigset, i int) {
	*mask &^= 1 << (sigset(i) - 1)
}

//go:nosplit
func (c *sigctxt) fixsigcode(sig uint32) {
}


// Called to initialize a new m (including the bootstrap m).
// Called on the new thread, cannot allocate memory.
func minit() {
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(miniterrno)), unsafe.Pointer(&libc__errnop))

	minitSignals()

	getg().m.procid = uint64(pthread_self())
}

// Called from dropm to undo the effect of an minit.
func unminit() {
	unminitSignals()
}

// Called from exitm, but not from drop, to undo the effect of thread-owned
// resources in minit, semacreate, or elsewhere. Do not take locks after calling this.
func mdestroy(mp *m) {
}

func sigtramp()

//go:nosplit
//go:nowritebarrierrec
func setsig(i uint32, fn uintptr) {
	var sa sigactiont

	sa.sa_flags = _SA_SIGINFO | _SA_ONSTACK | _SA_RESTART
	sa.sa_mask = sigset_all
	if fn == abi.FuncPCABIInternal(sighandler) { // abi.FuncPCABIInternal(sighandler) matches the callers in signal_unix.go
		fn = abi.FuncPCABI0(sigtramp)
	}
	*((*uintptr)(unsafe.Pointer(&sa._funcptr))) = fn
	sigaction(i, &sa, nil)
}

//go:nosplit
//go:nowritebarrierrec
func setsigstack(i uint32) {
	var sa sigactiont
	sigaction(i, nil, &sa)
	if sa.sa_flags&_SA_ONSTACK != 0 {
		return
	}
	sa.sa_flags |= _SA_ONSTACK
	sigaction(i, &sa, nil)
}

//go:nosplit
//go:nowritebarrierrec
func getsig(i uint32) uintptr {
	var sa sigactiont
	sigaction(i, nil, &sa)
	return *((*uintptr)(unsafe.Pointer(&sa._funcptr)))
}

func setProcessCPUProfiler(hz int32) {
	setProcessCPUProfilerTimer(hz)
}

func setThreadCPUProfiler(hz int32) {
	setThreadCPUProfilerHz(hz)
}

//go:nosplit
func validSIGPROF(mp *m, c *sigctxt) bool {
	return true
}

//go:nosplit
func semacreate(mp *m) {
	if mp.waitsema != 0 {
		return
	}

	var sem *semt
	_g_ := getg()

	// Call libc's malloc rather than malloc. This will
	// allocate space on the C heap. We can't call malloc
	// here because it could cause a deadlock.
	_g_.m.libcall.fn = uintptr(unsafe.Pointer(&libc_malloc))
	_g_.m.libcall.n = 1
	_g_.m.scratch = mscratch{}
	_g_.m.scratch.v[0] = unsafe.Sizeof(*sem)
	_g_.m.libcall.args = uintptr(unsafe.Pointer(&_g_.m.scratch))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&_g_.m.libcall))
	sem = (*semt)(unsafe.Pointer(_g_.m.libcall.r1))
	if sem_init(sem, 0, 0) != 0 {
		throw("sem_init")
	}
	mp.waitsema = uintptr(unsafe.Pointer(sem))
}

//go:nosplit
func semasleep(ns int64) int32 {
	mp := getg().m
	if ns >= 0 {
		var ts timespec

		if clock_gettime(_CLOCK_REALTIME, &ts) != 0 {
			throw("clock_gettime")
		}
		ts.tv_sec += ns / 1e9
		ts.tv_nsec += ns % 1e9
		if ts.tv_nsec >= 1e9 {
			ts.tv_sec++
			ts.tv_nsec -= 1e9
		}

		if r, err := sem_timedwait((*semt)(unsafe.Pointer(mp.waitsema)), &ts); r != 0 {
			if err == _ETIMEDOUT || err == _EAGAIN || err == _EINTR {
				return -1
			}
			println("sem_timedwait err ", err, " ts.tv_sec ", ts.tv_sec, " ts.tv_nsec ", ts.tv_nsec, " ns ", ns, " id ", mp.id)
			throw("sem_timedwait")
		}
		return 0
	}
	for {
		r1, err := sem_wait((*semt)(unsafe.Pointer(mp.waitsema)))
		if r1 == 0 {
			break
		}
		if err == _EINTR {
			continue
		}
		throw("sem_wait")
	}
	return 0
}

func exitThread(wait *atomic.Uint32) {
	// We should never reach exitThread on Haiku because we let
	// libroot clean up threads.
	throw("exitThread")
}

//go:nosplit
func semawakeup(mp *m) {
	if sem_post((*semt)(unsafe.Pointer(mp.waitsema))) != 0 {
		throw("sem_post")
	}
}

//go:nosplit
func closefd(fd int32) int32 {
	return int32(sysvicall1(&libc_close, uintptr(fd)))
}

//go:nosplit
func exit(r int32) {
	sysvicall1(&libc_exit, uintptr(r))
}

//go:nosplit
func getcontext(context *ucontext) /* int32 */ {
	sysvicall1(&libc_getcontext, uintptr(unsafe.Pointer(context)))
}

/////go:nosplit
// func madvise(addr unsafe.Pointer, n uintptr, flags int32) {
// 	sysvicall3(&libc_madvise, uintptr(addr), uintptr(n), uintptr(flags))
// }

//go:nosplit
func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) (unsafe.Pointer, int) {
	p, err := doMmap(uintptr(addr), n, uintptr(prot), uintptr(flags), uintptr(fd), uintptr(off))
	if p == ^uintptr(0) {
		return nil, int(err)
	}
	return unsafe.Pointer(p), 0
}

//go:nosplit
//go:cgo_unsafe_args
func doMmap(addr, n, prot, flags, fd, off uintptr) (uintptr, uintptr) {
	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(&libc_mmap))
	libcall.n = 6
	libcall.args = uintptr(noescape(unsafe.Pointer(&addr)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	return libcall.r1, libcall.err
}

//go:nosplit
func munmap(addr unsafe.Pointer, n uintptr) {
	sysvicall2(&libc_munmap, uintptr(addr), uintptr(n))
}

//go:nosplit
func nanotime1() int64 {
	var ts mts
	sysvicall2(&libc_clock_gettime, _CLOCK_MONOTONIC, uintptr(unsafe.Pointer(&ts)))
	return ts.tv_sec*1e9 + ts.tv_nsec
}

//go:nosplit
func fcntl(fd, cmd, arg int32) int32 {
	return int32(sysvicall3(&libc_fcntl, uintptr(fd), uintptr(cmd), uintptr(arg)))
}

//go:nosplit
func closeonexec(fd int32) {
	fcntl(fd, _F_SETFD, _FD_CLOEXEC)
}

//go:nosplit
func setNonblock(fd int32) {
	flags := fcntl(fd, _F_GETFL, 0)
	fcntl(fd, _F_SETFL, flags|_O_NONBLOCK)
}

//go:nosplit
func open(path *byte, mode, perm int32) int32 {
	return int32(sysvicall3(&libc_open, uintptr(unsafe.Pointer(path)), uintptr(mode), uintptr(perm)))
}

func pthread_attr_destroy(attr *pthreadattr) int32 {
	return int32(sysvicall1(&libc_pthread_attr_destroy, uintptr(unsafe.Pointer(attr))))
}

// func pthread_attr_getstack(attr *pthreadattr, addr unsafe.Pointer, size *uint64) int32 {
// 	return int32(sysvicall3(&libc_pthread_attr_getstack, uintptr(unsafe.Pointer(attr)), uintptr(addr), uintptr(unsafe.Pointer(size))))
// }

func pthread_attr_init(attr *pthreadattr) int32 {
	return int32(sysvicall1(&libc_pthread_attr_init, uintptr(unsafe.Pointer(attr))))
}

func pthread_attr_setdetachstate(attr *pthreadattr, state int32) int32 {
	return int32(sysvicall2(&libc_pthread_attr_setdetachstate, uintptr(unsafe.Pointer(attr)), uintptr(state)))
}

func pthread_attr_setstack(attr *pthreadattr, addr uintptr, size uint64) int32 {
	return int32(sysvicall3(&libc_pthread_attr_setstack, uintptr(unsafe.Pointer(attr)), uintptr(addr), uintptr(size)))
}

func pthread_create(thread *pthread, attr *pthreadattr, fn uintptr, arg unsafe.Pointer) int32 {
	return int32(sysvicall4(&libc_pthread_create, uintptr(unsafe.Pointer(thread)), uintptr(unsafe.Pointer(attr)), uintptr(fn), uintptr(arg)))
}

func pthread_self() pthread {
	return pthread(sysvicall0(&libc_pthread_self))
}

func signalM(mp *m, sig int) {
	sysvicall2(&libc_pthread_kill, uintptr(pthread(mp.procid)), uintptr(sig))
}

//go:nosplit
//go:nowritebarrierrec
func raise(sig uint32) /* int32 */ {
	sysvicall1(&libc_raise, uintptr(sig))
}

func raiseproc(sig uint32) /* int32 */ {
	pid := sysvicall0(&libc_getpid)
	sysvicall2(&libc_kill, pid, uintptr(sig))
}

//go:nosplit
func read(fd int32, buf unsafe.Pointer, nbyte int32) int32 {
	r1, err := sysvicall3Err(&libc_read, uintptr(fd), uintptr(buf), uintptr(nbyte))
	if c := int32(r1); c >= 0 {
		return c
	}
	return int32(err)
}

//go:nosplit
func sem_init(sem *semt, pshared int32, value uint32) int32 {
	return int32(sysvicall3(&libc_sem_init, uintptr(unsafe.Pointer(sem)), uintptr(pshared), uintptr(value)))
}

//go:nosplit
func sem_post(sem *semt) int32 {
	return int32(sysvicall1(&libc_sem_post, uintptr(unsafe.Pointer(sem))))
}

//go:nosplit
func sem_wait(sem *semt) (int32, int32) {
	r, err := sysvicall1Err(&libc_sem_wait, uintptr(unsafe.Pointer(sem)))
	return int32(r), int32(err)
}

//go:nosplit
func sem_timedwait(sem *semt, timeout *timespec) (int32, int32) {
	r, err := sysvicall2Err(&libc_sem_timedwait, uintptr(unsafe.Pointer(sem)), uintptr(unsafe.Pointer(timeout)))
	return int32(r), int32(err)
}

//go:nosplit
func clock_gettime(clockid uint32, tp *timespec) int32 {
	r := sysvicall2(&libc_clock_gettime, uintptr(clockid), uintptr(unsafe.Pointer(tp)))
	return int32(r)
}

func setitimer(which int32, value *itimerval, ovalue *itimerval) /* int32 */ {
	sysvicall3(&libc_setitimer, uintptr(which), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(ovalue)))
}

//go:nosplit
//go:nowritebarrierrec
func sigaction(sig uint32, act *sigactiont, oact *sigactiont) /* int32 */ {
	sysvicall3(&libc_sigaction, uintptr(sig), uintptr(unsafe.Pointer(act)), uintptr(unsafe.Pointer(oact)))
}

//go:nosplit
//go:nowritebarrierrec
func sigaltstack(ss *stackt, oss *stackt) /* int32 */ {
	sysvicall2(&libc_sigaltstack, uintptr(unsafe.Pointer(ss)), uintptr(unsafe.Pointer(oss)))
}

//go:nosplit
//go:nowritebarrierrec
func sigprocmask(how int32, set *sigset, oset *sigset) /* int32 */ {
	sysvicall3(&libc_sigprocmask, uintptr(how), uintptr(unsafe.Pointer(set)), uintptr(unsafe.Pointer(oset)))
}

func sysconf(name int32) int64 {
	return int64(sysvicall1(&libc_sysconf, uintptr(name)))
}

func usleep1(usec uint32)

//go:nosplit
func usleep_no_g(usec uint32) {
	sysvicall1(&libc_usleep, uintptr(usec))
}

//go:nosplit
func usleep(usec uint32) {
	sysvicall1(&libc_usleep, uintptr(usec))
}

func walltime() (sec int64, nsec int32) {
	var ts mts
	sysvicall2(&libc_clock_gettime, uintptr(_CLOCK_REALTIME), uintptr(unsafe.Pointer(&ts)))
	return ts.tv_sec, int32(ts.tv_nsec)
}

//go:nosplit
func write1(fd uintptr, buf unsafe.Pointer, nbyte int32) int32 {
	r1, err := sysvicall3Err(&libc_write, fd, uintptr(buf), uintptr(nbyte))
	if c := int32(r1); c >= 0 {
		return c
	}
	return int32(err)
}

//go:nosplit
func pipe() (r, w int32, errno int32) {
	var p [2]int32
	_, e := sysvicall1Err(&libc_pipe, uintptr(noescape(unsafe.Pointer(&p))))
	return p[0], p[1], int32(e)
}

//go:nosplit
func pipe2(flags int32) (r, w int32, errno int32) {
	var p [2]int32
	_, e := sysvicall2Err(&libc_pipe2, uintptr(noescape(unsafe.Pointer(&p))), uintptr(flags))
	return p[0], p[1], int32(e)
}

func osyield1()

//go:nosplit
func osyield_no_g() {
	osyield1()
}

//go:nosplit
func osyield() {
	_g_ := getg()

	// Check the validity of m because we might be called in cgo callback
	// path early enough where there isn't a m available yet.
	if _g_ != nil && _g_.m != nil {
		sysvicall0(&libc_sched_yield)
		return
	}
	osyield1()
}

func area_for(addr uintptr) {
	sysvicall1(&libc_area_for, uintptr(addr))
}

// sigPerThreadSyscall is only used on linux, so we assign a bogus signal
// number.
const sigPerThreadSyscall = 1 << 31

//go:nosplit
func runPerThreadSyscall() {
	throw("runPerThreadSyscall only valid on linux")
}
