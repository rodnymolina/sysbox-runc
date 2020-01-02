//
// Copyright: (C) 2019 Nestybox Inc.  All rights reserved.
//

// +build linux

package syscont

import (
	"fmt"

	"github.com/opencontainers/runc/libcontainer/configs"
)

// List of syscalls allowed inside a system container
var syscontSyscallWhitelist = []string{

	// docker allows these by default
	"accept",
	"accept4",
	"access",
	"adjtimex",
	"alarm",
	"bind",
	"brk",
	"capget",
	"capset",
	"chdir",
	"chmod",
	"chown",
	"chown32",
	"clock_getres",
	"clock_gettime",
	"clock_nanosleep",
	"close",
	"connect",
	"copy_file_range",
	"creat",
	"dup",
	"dup2",
	"dup3",
	"epoll_create",
	"epoll_create1",
	"epoll_ctl",
	"epoll_ctl_old",
	"epoll_pwait",
	"epoll_wait",
	"epoll_wait_old",
	"eventfd",
	"eventfd2",
	"execve",
	"execveat",
	"exit",
	"exit_group",
	"faccessat",
	"fadvise64",
	"fadvise64_64",
	"fallocate",
	"fanotify_mark",
	"fchdir",
	"fchmod",
	"fchmodat",
	"fchown",
	"fchown32",
	"fchownat",
	"fcntl",
	"fcntl64",
	"fdatasync",
	"fgetxattr",
	"flistxattr",
	"flock",
	"fork",
	"fremovexattr",
	"fsetxattr",
	"fstat",
	"fstat64",
	"fstatat64",
	"fstatfs",
	"fstatfs64",
	"fsync",
	"ftruncate",
	"ftruncate64",
	"futex",
	"futimesat",
	"getcpu",
	"getcwd",
	"getdents",
	"getdents64",
	"getegid",
	"getegid32",
	"geteuid",
	"geteuid32",
	"getgid",
	"getgid32",
	"getgroups",
	"getgroups32",
	"getitimer",
	"getpeername",
	"getpgid",
	"getpgrp",
	"getpid",
	"getppid",
	"getpriority",
	"getrandom",
	"getresgid",
	"getresgid32",
	"getresuid",
	"getresuid32",
	"getrlimit",
	"get_robust_list",
	"getrusage",
	"getsid",
	"getsockname",
	"getsockopt",
	"get_thread_area",
	"gettid",
	"gettimeofday",
	"getuid",
	"getuid32",
	"getxattr",
	"inotify_add_watch",
	"inotify_init",
	"inotify_init1",
	"inotify_rm_watch",
	"io_cancel",
	"ioctl",
	"io_destroy",
	"io_getevents",
	"ioprio_get",
	"ioprio_set",
	"io_setup",
	"io_submit",
	"ipc",
	"kill",
	"lchown",
	"lchown32",
	"lgetxattr",
	"link",
	"linkat",
	"listen",
	"listxattr",
	"llistxattr",
	"_llseek",
	"lremovexattr",
	"lseek",
	"lsetxattr",
	"lstat",
	"lstat64",
	"madvise",
	"memfd_create",
	"mincore",
	"mkdir",
	"mkdirat",
	"mknod",
	"mknodat",
	"mlock",
	"mlock2",
	"mlockall",
	"mmap",
	"mmap2",
	"mprotect",
	"mq_getsetattr",
	"mq_notify",
	"mq_open",
	"mq_timedreceive",
	"mq_timedsend",
	"mq_unlink",
	"mremap",
	"msgctl",
	"msgget",
	"msgrcv",
	"msgsnd",
	"msync",
	"munlock",
	"munlockall",
	"munmap",
	"nanosleep",
	"newfstatat",
	"_newselect",
	"open",
	"openat",
	"pause",
	"pipe",
	"pipe2",
	"poll",
	"ppoll",
	"prctl",
	"pread64",
	"preadv",
	"preadv2",
	"prlimit64",
	"pselect6",
	"pwrite64",
	"pwritev",
	"pwritev2",
	"read",
	"readahead",
	"readlink",
	"readlinkat",
	"readv",
	"recv",
	"recvfrom",
	"recvmmsg",
	"recvmsg",
	"remap_file_pages",
	"removexattr",
	"rename",
	"renameat",
	"renameat2",
	"restart_syscall",
	"rmdir",
	"rt_sigaction",
	"rt_sigpending",
	"rt_sigprocmask",
	"rt_sigqueueinfo",
	"rt_sigreturn",
	"rt_sigsuspend",
	"rt_sigtimedwait",
	"rt_tgsigqueueinfo",
	"sched_getaffinity",
	"sched_getattr",
	"sched_getparam",
	"sched_get_priority_max",
	"sched_get_priority_min",
	"sched_getscheduler",
	"sched_rr_get_interval",
	"sched_setaffinity",
	"sched_setattr",
	"sched_setparam",
	"sched_setscheduler",
	"sched_yield",
	"seccomp",
	"select",
	"semctl",
	"semget",
	"semop",
	"semtimedop",
	"send",
	"sendfile",
	"sendfile64",
	"sendmmsg",
	"sendmsg",
	"sendto",
	"setfsgid",
	"setfsgid32",
	"setfsuid",
	"setfsuid32",
	"setgid",
	"setgid32",
	"setgroups",
	"setgroups32",
	"setitimer",
	"setpgid",
	"setpriority",
	"setregid",
	"setregid32",
	"setresgid",
	"setresgid32",
	"setresuid",
	"setresuid32",
	"setreuid",
	"setreuid32",
	"setrlimit",
	"set_robust_list",
	"setsid",
	"setsockopt",
	"set_thread_area",
	"set_tid_address",
	"setuid",
	"setuid32",
	"setxattr",
	"shmat",
	"shmctl",
	"shmdt",
	"shmget",
	"shutdown",
	"sigaltstack",
	"signalfd",
	"signalfd4",
	"sigreturn",
	"socket",
	"socketcall",
	"socketpair",
	"splice",
	"stat",
	"stat64",
	"statfs",
	"statfs64",
	"statx",
	"symlink",
	"symlinkat",
	"sync",
	"sync_file_range",
	"syncfs",
	"sysinfo",
	"tee",
	"tgkill",
	"time",
	"timer_create",
	"timer_delete",
	"timerfd_create",
	"timerfd_gettime",
	"timerfd_settime",
	"timer_getoverrun",
	"timer_gettime",
	"timer_settime",
	"times",
	"tkill",
	"truncate",
	"truncate64",
	"ugetrlimit",
	"umask",
	"uname",
	"unlink",
	"unlinkat",
	"utime",
	"utimensat",
	"utimes",
	"vfork",
	"vmsplice",
	"wait4",
	"waitid",
	"waitpid",
	"write",
	"writev",

	"personality",
	"arch_prctl",
	"modify_ldt",
	"clone",
	"chroot",

	// docker blocks these by default; sysbox-runc allows them
	"mount",
	"umount",
	"umount2",
	"add_key",
	"request_key",
	"keyctl",
	"pivot_root",
	"gethostname",
	"sethostname",

	// allow namespace creation inside the system container (for nested containers)
	"setns",
	"unshare",
}

// List of syscalls trapped & emulated inside a system container
//
// NOTE: all of these must also be in the syscontSyscallWhitelist, as otherwise seccomp
// will block them.
var syscontSyscallTrapList = []string{
	"mount",
}

// AddSyscallTraps modifies the given libcontainer config to add seccomp notification
// actions for syscall trapping
func AddSyscallTraps(config *configs.Config) error {

	if config.SeccompNotif != nil {
		return fmt.Errorf("conflicting seccomp notification config found.")
	}

	if len(syscontSyscallTrapList) > 0 {
		list := []*configs.Syscall{}
		for _, call := range syscontSyscallTrapList {
			s := &configs.Syscall{
				Name:   call,
				Action: configs.Notify,
			}
			list = append(list, s)
		}

		config.SeccompNotif = &configs.Seccomp{
			DefaultAction: configs.Allow,
			Architectures: []string{"amd64"},
			Syscalls:      list,
		}
	}

	return nil
}
