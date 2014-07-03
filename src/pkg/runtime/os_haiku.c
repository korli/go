// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "runtime.h"
#include "defs_GOOS_GOARCH.h"
#include "os_GOOS.h"
#include "signal_unix.h"
#include "stack.h"
#include "../../cmd/ld/textflag.h"

#pragma dynexport end _end
#pragma dynexport etext _etext
#pragma dynexport edata _edata

#pragma dynimport libc·_errnop _errnop "libroot.so"
#pragma dynimport libc·clock_gettime clock_gettime "libroot.so"
#pragma dynimport libc·close close "libroot.so"
#pragma dynimport libc·exit exit "libroot.so"
#pragma dynimport libc·fstat fstat "libroot.so"
#pragma dynimport libc·getcontext getcontext "libroot.so"
#pragma dynimport libc·getrlimit getrlimit "libroot.so"
#pragma dynimport libc·malloc malloc "libroot.so"
#pragma dynimport libc·mmap mmap "libroot.so"
#pragma dynimport libc·munmap munmap "libroot.so"
#pragma dynimport libc·open open "libroot.so"
#pragma dynimport libc·pthread_attr_destroy pthread_attr_destroy "libroot.so"
#pragma dynimport libc·pthread_attr_getstack pthread_attr_getstack "libroot.so"
#pragma dynimport libc·pthread_attr_init pthread_attr_init "libroot.so"
#pragma dynimport libc·pthread_attr_setdetachstate pthread_attr_setdetachstate "libroot.so"
#pragma dynimport libc·pthread_attr_setstack pthread_attr_setstack "libroot.so"
#pragma dynimport libc·pthread_create pthread_create "libroot.so"
#pragma dynimport libc·raise raise "libroot.so"
#pragma dynimport libc·read read "libroot.so"
#pragma dynimport libc·select select "libroot.so"
#pragma dynimport libc·sched_yield sched_yield "libroot.so"
#pragma dynimport libc·sem_init sem_init "libroot.so"
#pragma dynimport libc·sem_post sem_post "libroot.so"
#pragma dynimport libc·sem_reltimedwait_np sem_reltimedwait_np "libroot.so"
#pragma dynimport libc·sem_wait sem_wait "libroot.so"
#pragma dynimport libc·setitimer setitimer "libroot.so"
#pragma dynimport libc·sigaction sigaction "libroot.so"
#pragma dynimport libc·sigaltstack sigaltstack "libroot.so"
#pragma dynimport libc·sigprocmask sigprocmask "libroot.so"
#pragma dynimport libc·sysconf sysconf "libroot.so"
#pragma dynimport libc·usleep usleep "libroot.so"
#pragma dynimport libc·write write "libroot.so"
#pragma dynimport libc·_kern_reserve_address_range _kern_reserve_address_range "libroot.so"

extern uintptr libc·_errnop;
extern uintptr libc·clock_gettime;
extern uintptr libc·close;
extern uintptr libc·exit;
extern uintptr libc·fstat;
extern uintptr libc·getcontext;
extern uintptr libc·getrlimit;
extern uintptr libc·malloc;
extern uintptr libc·mmap;
extern uintptr libc·munmap;
extern uintptr libc·open;
extern uintptr libc·pthread_attr_destroy;
extern uintptr libc·pthread_attr_getstack;
extern uintptr libc·pthread_attr_init;
extern uintptr libc·pthread_attr_setdetachstate;
extern uintptr libc·pthread_attr_setstack;
extern uintptr libc·pthread_create;
extern uintptr libc·raise;
extern uintptr libc·read;
extern uintptr libc·sched_yield;
extern uintptr libc·select;
extern uintptr libc·sem_init;
extern uintptr libc·sem_post;
extern uintptr libc·sem_reltimedwait_np;
extern uintptr libc·sem_wait;
extern uintptr libc·setitimer;
extern uintptr libc·sigaction;
extern uintptr libc·sigaltstack;
extern uintptr libc·sigprocmask;
extern uintptr libc·sysconf;
extern uintptr libc·usleep;
extern uintptr libc·write;
extern uintptr libc·_kern_reserve_address_range;

void	runtime·getcontext(Ucontext *context);
int32	runtime·pthread_attr_destroy(PthreadAttr* attr);
int32	runtime·pthread_attr_init(PthreadAttr* attr);
//TODO: uint32/uint64 problem needs to be solved
int32	runtime·pthread_attr_getstack(PthreadAttr* attr, void** addr, uint32* size);
int32	runtime·pthread_attr_setdetachstate(PthreadAttr* attr, int32 state);
int32	runtime·pthread_attr_setstack(PthreadAttr* attr, void* addr, uint32 size);
int32	runtime·pthread_create(Pthread* thread, PthreadAttr* attr, void(*fn)(void), void *arg);
uint32	runtime·tstart_sysvicall(M *newm);
int32	runtime·sem_init(SemT* sem, int32 pshared, uint32 value);
int32	runtime·sem_post(SemT* sem);
int32	runtime·sem_reltimedwait_np(SemT* sem, Timespec* timeout);
int32	runtime·sem_wait(SemT* sem);
int64	runtime·sysconf(int32 name);

extern SigTab runtime·sigtab[];
/*static Sigset sigset_none;*/
/*static Sigset sigset_all = { ~(uint32)0, ~(uint32)0, ~(uint32)0, ~(uint32)0, };*/

// Calling sysvcall on os stack.
#pragma textflag NOSPLIT
uintptr
runtime·sysvicall6(uintptr fn, int32 count, ...)
{
	runtime·memclr((byte*)&m->scratch, sizeof(m->scratch));
	m->libcall.fn = (void*)fn;
	m->libcall.n = (uintptr)count;
	for(;count; count--)
		m->scratch.v[count - 1] = *((uintptr*)&count + count);
	m->libcall.args = (uintptr*)&m->scratch.v[0];
	runtime·asmcgocall(runtime·asmsysvicall6, &m->libcall);
	return m->libcall.r1;
}

static int32
getncpu(void)
{
/*	int32 n;

	n = (int32)runtime·sysconf(_SC_NPROCESSORS_ONLN);
	if(n < 1)
		return 1;
	return n;
*/
	return 1;
}

void
runtime·osinit(void)
{
	runtime·ncpu = getncpu();
}

void
runtime·newosproc(M *mp, void *stk)
{
	PthreadAttr attr;
	Sigset oset;
	Pthread tid;
	int32 ret;

	USED(stk);
	if(runtime·pthread_attr_init(&attr) != 0)
		runtime·throw("pthread_attr_init");
	//if(runtime·pthread_attr_setstack(&attr, 0, 0x200000) != 0)
	//	runtime·throw("pthread_attr_setstack");
	//if(runtime·pthread_attr_getstack(&attr, (void**)&mp->g0->stackbase, &mp->g0->stacksize) != 0)
	//	runtime·throw("pthread_attr_getstack");
	if(runtime·pthread_attr_setdetachstate(&attr, PTHREAD_CREATE_DETACHED) != 0)
		runtime·throw("pthread_attr_setdetachstate");

	// Disable signals during create, so that the new thread starts
	// with signals disabled.  It will enable them in minit.
	//runtime·sigprocmask(SIG_SETMASK, &sigset_all, &oset);
	mp->g0->stackbase = 0xdeadf00d;
	mp->g0->stacksize = 0xdead;
	ret = runtime·pthread_create(&tid, &attr, (void (*)(void))runtime·tstart_sysvicall, mp);
	//runtime·sigprocmask(SIG_SETMASK, &oset, nil);
	if(ret != 0) {
		runtime·printf("runtime: failed to create new OS thread (have %d already; errno=%d)\n", runtime·mcount(), ret);
		runtime·throw("runtime.newosproc");
	}
}

void
runtime·get_random_data(byte **rnd, int32 *rnd_len)
{
	static byte urandom_data[HashRandomBytes];
	int32 fd;
	fd = runtime·open("/dev/urandom", 0 /* O_RDONLY */, 0);
	if(runtime·read(fd, urandom_data, HashRandomBytes) == HashRandomBytes) {
		*rnd = urandom_data;
		*rnd_len = HashRandomBytes;
	} else {
		*rnd = nil;
		*rnd_len = 0;
	}
	runtime·close(fd);
}

void
runtime·goenvs(void)
{
	runtime·goenvs_unix();
}

// Called to initialize a new m (including the bootstrap m).
// Called on the parent thread (main thread in case of bootstrap), can allocate memory.
void
runtime·mpreinit(M *mp)
{
	mp->gsignal = runtime·malg(32*1024);
}

// Called to initialize a new m (including the bootstrap m).
// Called on the new thread, can not allocate memory.
void
runtime·minit(void)
{
	runtime·asmcgocall(runtime·miniterrno, (void *) &libc·_errnop);
	// Initialize signal handling
	runtime·signalstack((byte*)m->gsignal->stackguard - StackGuard, 32*1024);
	//runtime·sigprocmask(SIG_SETMASK, &sigset_none, nil);
}

// Called from dropm to undo the effect of an minit.
void
runtime·unminit(void)
{
	runtime·signalstack(nil, 0);
}

void
runtime·sigpanic(void)
{
	/*if(!runtime·canpanic(g))
		runtime·throw("unexpected signal during runtime execution");

	switch(g->sig) {
	case SIGBUS:
		if(g->sigcode0 == BUS_ADRERR && g->sigcode1 < 0x1000 || g->paniconfault) {
			if(g->sigpc == 0)
				runtime·panicstring("call of nil func value");
			runtime·panicstring("invalid memory address or nil pointer dereference");
		}
		runtime·printf("unexpected fault address %p\n", g->sigcode1);
		runtime·throw("fault");
	case SIGSEGV:
		if((g->sigcode0 == 0 || g->sigcode0 == SEGV_MAPERR || g->sigcode0 == SEGV_ACCERR) && g->sigcode1 < 0x1000 || g->paniconfault) {
			if(g->sigpc == 0)
				runtime·panicstring("call of nil func value");
			runtime·panicstring("invalid memory address or nil pointer dereference");
		}
		runtime·printf("unexpected fault address %p\n", g->sigcode1);
		runtime·throw("fault");
	case SIGFPE:
		switch(g->sigcode0) {
		case FPE_INTDIV:
			runtime·panicstring("integer divide by zero");
		case FPE_INTOVF:
			runtime·panicstring("integer overflow");
		}
		runtime·panicstring("floating point error");
	}
	runtime·panicstring(runtime·sigtab[g->sig].name);
	*/
}

uintptr
runtime·memlimit(void)
{
	/*Rlimit rl;
	extern byte text[], end[];
	uintptr used;

	if(runtime·getrlimit(RLIMIT_AS, &rl) != 0)
		return 0;
	if(rl.rlim_cur >= 0x7fffffff)
		return 0;

	// Estimate our VM footprint excluding the heap.
	// Not an exact science: use size of binary plus
	// some room for thread stacks.
	used = end - text + (64<<20);
	if(used >= rl.rlim_cur)
		return 0;

	// If there's not at least 16 MB left, we're probably
	// not going to be able to do much.  Treat as no limit.
	rl.rlim_cur -= used;
	if(rl.rlim_cur < (16<<20))
		return 0;

	return rl.rlim_cur - used;
	*/
	return 0;
}

void
runtime·setprof(bool on)
{
	USED(on);
}

extern void runtime·sigtramp(void);

void
runtime·setsig(int32 i, GoSighandler *fn, bool restart)
{
	/*Sigaction sa;

	runtime·memclr((byte*)&sa, sizeof sa);
	sa.sa_flags = SA_SIGINFO|SA_ONSTACK;
	if(restart)
		sa.sa_flags |= SA_RESTART;
	sa.sa_mask.__sigbits[0] = ~(uint32)0;
	sa.sa_mask.__sigbits[1] = ~(uint32)0;
	sa.sa_mask.__sigbits[2] = ~(uint32)0;
	sa.sa_mask.__sigbits[3] = ~(uint32)0;
	if(fn == runtime·sighandler)
		fn = (void*)runtime·sigtramp;
	*((void**)&sa._funcptr[0]) = (void*)fn;
	runtime·sigaction(i, &sa, nil);
	*/
}

GoSighandler*
runtime·getsig(int32 i)
{
	/*Sigaction sa;

	runtime·memclr((byte*)&sa, sizeof sa);
	runtime·sigaction(i, nil, &sa);
	if(*((void**)&sa._funcptr[0]) == runtime·sigtramp)
		return runtime·sighandler;
	return *((void**)&sa._funcptr[0]);
	*/
	return (void*) 0x1c1c1e;
}

void
runtime·signalstack(byte *p, int32 n)
{
	/*StackT st;

	st.ss_sp = (void*)p;
	st.ss_size = n;
	st.ss_flags = 0;
	if(p == nil)
		st.ss_flags = SS_DISABLE;
	runtime·sigaltstack(&st, nil);
	*/
}

void
runtime·unblocksignals(void)
{
	//runtime·sigprocmask(SIG_SETMASK, &sigset_none, nil);
}

#pragma textflag NOSPLIT
uintptr
runtime·semacreate(void)
{
	SemT* sem;

	// Call libc's malloc rather than runtime·malloc.  This will
	// allocate space on the C heap.  We can't call runtime·malloc
	// here because it could cause a deadlock.
	m->libcall.fn = (void*)(uintptr) &libc·malloc;
	m->libcall.n = 1;
	runtime·memclr((byte*)&m->scratch, sizeof(m->scratch));
	m->scratch.v[0] = (uintptr)sizeof(*sem);
	m->libcall.args = (uintptr*)&m->scratch;
	runtime·asmcgocall(runtime·asmsysvicall6, &m->libcall);
	sem = (void*)m->libcall.r1;
	if(runtime·sem_init(sem, 0, 0) != 0)
		runtime·throw("sem_init");
	return (uintptr)sem;
}

#pragma textflag NOSPLIT
int32
runtime·semasleep(int64 ns)
{
	if(ns >= 0) {
		/*
		m->ts.tv_sec = ns / 1000000000LL;
		m->ts.tv_nsec = ns % 1000000000LL;

		m->libcall.fn = (void*)(uintptr) &libc·sem_reltimedwait_np;
		m->libcall.n = 2;
		runtime·memclr((byte*)&m->scratch, sizeof(m->scratch));
		m->scratch.v[0] = m->waitsema;
		m->scratch.v[1] = (uintptr)&m->ts;
		m->libcall.args = (uintptr*)&m->scratch;
		runtime·asmcgocall(runtime·asmsysvicall6, &m->libcall);
		if(*m->perrno != 0) {
			if(*m->perrno == ETIMEDOUT || *m->perrno == EAGAIN || *m->perrno == EINTR)
				return -1;
			runtime·throw("sem_reltimedwait_np");
		}
		return 0;
		*/
		runtime·throw("Todo: implement semasleep");
		return -1;
	}
	for(;;) {
		m->libcall.fn = (void*)(uintptr) &libc·sem_wait;
		m->libcall.n = 1;
		runtime·memclr((byte*)&m->scratch, sizeof(m->scratch));
		m->scratch.v[0] = m->waitsema;
		m->libcall.args = (uintptr*)&m->scratch;
		runtime·asmcgocall(runtime·asmsysvicall6, &m->libcall);
		if(m->libcall.r1 == 0)
			break;
		//if(*m->perrno == EINTR)
		//	continue;
		runtime·throw("sem_wait");
	}
	return 0;
}

#pragma textflag NOSPLIT
void
runtime·semawakeup(M *mp)
{
	SemT* sem = (SemT*)mp->waitsema;
	if(runtime·sem_post(sem) != 0)
		runtime·throw("sem_post");
}

int32
runtime·close(int32 fd)
{
	return runtime·sysvicall6((uintptr) &libc·close, 1, (uintptr)fd);
}

void
runtime·exit(int32 r)
{
	runtime·sysvicall6((uintptr) &libc·exit, 1, (uintptr)r);
}

/* int32 */ void
runtime·getcontext(Ucontext* context)
{
	runtime·sysvicall6((uintptr) &libc·getcontext, 1, (uintptr)context);
}

int32
runtime·getrlimit(int32 res, Rlimit* rlp)
{
	return runtime·sysvicall6((uintptr) &libc·getrlimit, 2, (uintptr)res, (uintptr)rlp);
}

uint8*
runtime·mmap(byte* addr, uintptr len, int32 prot, int32 flags, int32 fildes, uint32 off)
{
	return (uint8*)runtime·sysvicall6((uintptr) &libc·mmap, 6, (uintptr)addr, (uintptr)len, (uintptr)prot, (uintptr)flags, (uintptr)fildes, (uintptr)off);
}

void
runtime·munmap(byte* addr, uintptr len)
{
	runtime·sysvicall6((uintptr) &libc·munmap, 2, (uintptr)addr, (uintptr)len);
}

//extern int64 runtime·nanotime1(void);
#pragma textflag NOSPLIT
int64
runtime·nanotime(void)
{
	return 0;//runtime·sysvicall6((uintptr)runtime·nanotime1, 0);
}

void
time·now(int64 sec, int32 usec)
{
	int64 ns;

	ns = runtime·nanotime();
	sec = ns / 1000000000LL;
	usec = ns - sec * 1000000000LL;
	FLUSH(&sec);
	FLUSH(&usec);
}

int32
runtime·open(int8* path, int32 oflag, int32 mode)
{
	return runtime·sysvicall6((uintptr) &libc·open, 3, (uintptr)path, (uintptr)oflag, (uintptr)mode);
}

int32
runtime·pthread_attr_destroy(PthreadAttr* attr)
{
	return runtime·sysvicall6((uintptr) &libc·pthread_attr_destroy, 1, (uintptr)attr);
}

//TODO size should be defined as a size_t for compat with 64-bit

int32
runtime·pthread_attr_getstack(PthreadAttr* attr, void** addr, uint32* size)
{
	return runtime·sysvicall6((uintptr) &libc·pthread_attr_getstack, 3, (uintptr)attr, (uintptr)addr, (uintptr)size);
}

int32
runtime·pthread_attr_init(PthreadAttr* attr)
{
	return runtime·sysvicall6((uintptr) &libc·pthread_attr_init, 1, (uintptr)attr);
}

int32
runtime·pthread_attr_setdetachstate(PthreadAttr* attr, int32 state)
{
	return runtime·sysvicall6((uintptr) &libc·pthread_attr_setdetachstate, 2, (uintptr)attr, (uintptr)state);
}

int32
runtime·pthread_attr_setstack(PthreadAttr* attr, void* addr, uint32 size)
{
	return runtime·sysvicall6((uintptr) &libc·pthread_attr_setstack, 3, (uintptr)attr, (uintptr)addr, (uintptr)size);
}

int32
runtime·pthread_create(Pthread* thread, PthreadAttr* attr, void(*fn)(void), void *arg)
{
	return runtime·sysvicall6((uintptr) &libc·pthread_create, 4, (uintptr)thread, (uintptr)attr, (uintptr)fn, (uintptr)arg);
}

/* int32 */ void
runtime·raise(int32 sig)
{
	runtime·sysvicall6((uintptr) &libc·raise, 1, (uintptr)sig);
}

int32
runtime·read(int32 fd, void* buf, int32 nbyte)
{
	return runtime·sysvicall6((uintptr) &libc·read, 3, (uintptr)fd, (uintptr)buf, (uintptr)nbyte);
}

#pragma textflag NOSPLIT
int32
runtime·sem_init(SemT* sem, int32 pshared, uint32 value)
{
	return runtime·sysvicall6((uintptr) &libc·sem_init, 3, (uintptr)sem, (uintptr)pshared, (uintptr)value);
}

#pragma textflag NOSPLIT
int32
runtime·sem_post(SemT* sem)
{
	return runtime·sysvicall6((uintptr) &libc·sem_post, 1, (uintptr)sem);
}

#pragma textflag NOSPLIT
int32
runtime·sem_reltimedwait_np(SemT* sem, Timespec* timeout)
{
	return runtime·sysvicall6((uintptr) &libc·sem_reltimedwait_np, 2, (uintptr)sem, (uintptr)timeout);
}

#pragma textflag NOSPLIT
int32
runtime·sem_wait(SemT* sem)
{
	return runtime·sysvicall6((uintptr) &libc·sem_wait, 1, (uintptr)sem);
}

/* int32 */ void
runtime·setitimer(int32 which, Itimerval* value, Itimerval* ovalue)
{
	runtime·sysvicall6((uintptr) &libc·setitimer, 3, (uintptr)which, (uintptr)value, (uintptr)ovalue);
}

/* int32 */ void
runtime·sigaction(int32 sig, struct Sigaction* act, struct Sigaction* oact)
{
	runtime·sysvicall6((uintptr) &libc·sigaction, 3, (uintptr)sig, (uintptr)act, (uintptr)oact);
}

/* int32 */ void
runtime·sigaltstack(Sigaltstack* ss, Sigaltstack* oss)
{
	runtime·sysvicall6((uintptr) &libc·sigaltstack, 2, (uintptr)ss, (uintptr)oss);
}

/* int32 */ void
runtime·sigprocmask(int32 how, Sigset* set, Sigset* oset)
{
	runtime·sysvicall6((uintptr) &libc·sigprocmask, 3, (uintptr)how, (uintptr)set, (uintptr)oset);
}

int64
runtime·sysconf(int32 name)
{
	return runtime·sysvicall6((uintptr) &libc·sysconf, 1, (uintptr)name);
}

void
runtime·usleep(uint32 us)
{
	runtime·sysvicall6((uintptr) &libc·usleep, 1, (uintptr)us);
}

int32
runtime·write(int32 fd, void* buf, int32 nbyte)
{
	return runtime·sysvicall6((uintptr) &libc·write, 3, (uintptr)fd, (uintptr)buf, (uintptr)nbyte);
}

void
runtime·osyield(void)
{
	runtime·sysvicall6((uintptr) &libc·sched_yield, 0);
}

int32
runtime·kern_reserve_address_range(byte* addr, int32 address_spec, uintptr len)
{
	return (int32)runtime·sysvicall6((uintptr) &libc·_kern_reserve_address_range, 3, (uintptr)addr, (uintptr) address_spec, (uintptr)len);
}


// STUB
void
runtime·pipe1(void)
{
	runtime·throw("Pipe not implemented");
}
