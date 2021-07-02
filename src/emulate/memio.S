.equ	memio, 0xffff0000
.equ	memio_exit, 0xffff0003

jumptable:
		la		t0, 1f
		slli	a7, a7, 2
		add		t0, t0, a7
		jr		t0
1:		
		j		syscall_io_setup                     # 0
		j		syscall_io_destroy                   # 1
		j		syscall_io_submit                    # 2
		j		syscall_io_cancel                    # 3
		j		syscall_unimp                        # 4
		j		syscall_setxattr                     # 5
		j		syscall_lsetxattr                    # 6
		j		syscall_fsetxattr                    # 7
		j		syscall_getxattr                     # 8
		j		syscall_lgetxattr                    # 9
		j		syscall_fgetxattr                    # 10
		j		syscall_listxattr                    # 11
		j		syscall_llistxattr                   # 12
		j		syscall_flistxattr                   # 13
		j		syscall_removexattr                  # 14
		j		syscall_lremovexattr                 # 15
		j		syscall_fremovexattr                 # 16
		j		syscall_getcwd                       # 17
		j		syscall_lookup_dcookie               # 18
		j		syscall_eventfd2                     # 19
		j		syscall_epoll_create1                # 20
		j		syscall_epoll_ctl                    # 21
		j		syscall_epoll_pwait                  # 22
		j		syscall_dup                          # 23
		j		syscall_dup3                         # 24
		j		syscall_fcntl64                      # 25
		j		syscall_inotify_init1                # 26
		j		syscall_inotify_add_watch            # 27
		j		syscall_inotify_rm_watch             # 28
		j		syscall_ioctl                        # 29
		j		syscall_ioprio_set                   # 30
		j		syscall_ioprio_get                   # 31
		j		syscall_flock                        # 32
		j		syscall_mknodat                      # 33
		j		syscall_mkdirat                      # 34
		j		syscall_unlinkat                     # 35
		j		syscall_symlinkat                    # 36
		j		syscall_linkat                       # 37
		j		syscall_unimp                        # 38
		j		syscall_umount2                      # 39
		j		syscall_mount                        # 40
		j		syscall_pivot_root                   # 41
		j		syscall_nfsservctl                   # 42
		j		syscall_statfs64                     # 43
		j		syscall_fstatfs64                    # 44
		j		syscall_truncate64                   # 45
		j		syscall_ftruncate64                  # 46
		j		syscall_fallocate                    # 47
		j		syscall_faccessat                    # 48
		j		syscall_chdir                        # 49
		j		syscall_fchdir                       # 50
		j		syscall_chroot                       # 51
		j		syscall_fchmod                       # 52
		j		syscall_fchmodat                     # 53
		j		syscall_fchownat                     # 54
		j		syscall_fchown                       # 55
		j		syscall_openat                       # 56
		j		syscall_close                        # 57
		j		syscall_vhangup                      # 58
		j		syscall_pipe2                        # 59
		j		syscall_quotactl                     # 60
		j		syscall_getdents64                   # 61
		j		syscall_unimp                        # 62
		j		syscall_read                         # 63
		j		syscall_write                        # 64
		j		syscall_readv                        # 65
		j		syscall_writev                       # 66
		j		syscall_pread64                      # 67
		j		syscall_pwrite64                     # 68
		j		syscall_preadv                       # 69
		j		syscall_pwritev                      # 70
		j		syscall_sendfile64                   # 71
		j		syscall_signalfd4                    # 72
		j		syscall_unimp                        # 73
		j		syscall_unimp                        # 74
		j		syscall_vmsplice                     # 75
		j		syscall_splice                       # 76
		j		syscall_tee                          # 77
		j		syscall_readlinkat                   # 78
		j		syscall_unimp                        # 79
		j		syscall_unimp                        # 80
		j		syscall_sync                         # 81
		j		syscall_fsync                        # 82
		j		syscall_fdatasync                    # 83
		j		syscall_sync_file_range              # 84
		j		syscall_timerfd_create               # 85
		j		syscall_unimp                        # 86
		j		syscall_unimp                        # 87
		j		syscall_unimp                        # 88
		j		syscall_acct                         # 89
		j		syscall_capget                       # 90
		j		syscall_capset                       # 91
		j		syscall_personality                  # 92
		j		syscall_exit                         # 93
		j		syscall_exit_group                   # 94
		j		syscall_waitid                       # 95
		j		syscall_set_tid_address              # 96
		j		syscall_unshare                      # 97
		j		syscall_unimp                        # 98
		j		syscall_set_robust_list              # 99
		j		syscall_get_robust_list              # 100
		j		syscall_unimp                        # 101
		j		syscall_getitimer                    # 102
		j		syscall_setitimer                    # 103
		j		syscall_kexec_load                   # 104
		j		syscall_init_module                  # 105
		j		syscall_delete_module                # 106
		j		syscall_timer_create                 # 107
		j		syscall_unimp                        # 108
		j		syscall_timer_getoverrun             # 109
		j		syscall_unimp                        # 110
		j		syscall_timer_delete                 # 111
		j		syscall_unimp                        # 112
		j		syscall_unimp                        # 113
		j		syscall_unimp                        # 114
		j		syscall_unimp                        # 115
		j		syscall_syslog                       # 116
		j		syscall_ptrace                       # 117
		j		syscall_sched_setparam               # 118
		j		syscall_sched_setscheduler           # 119
		j		syscall_sched_getscheduler           # 120
		j		syscall_sched_getparam               # 121
		j		syscall_sched_setaffinity            # 122
		j		syscall_sched_getaffinity            # 123
		j		syscall_sched_yield                  # 124
		j		syscall_sched_get_priority_max       # 125
		j		syscall_sched_get_priority_min       # 126
		j		syscall_unimp                        # 127
		j		syscall_restart_syscall              # 128
		j		syscall_kill                         # 129
		j		syscall_tkill                        # 130
		j		syscall_tgkill                       # 131
		j		syscall_sigaltstack                  # 132
		j		syscall_rt_sigsuspend                # 133
		j		syscall_rt_sigaction                 # 134
		j		syscall_rt_sigprocmask               # 135
		j		syscall_rt_sigpending                # 136
		j		syscall_unimp                        # 137
		j		syscall_rt_sigqueueinfo              # 138
		j		syscall_rt_sigreturn                 # 139
		j		syscall_setpriority                  # 140
		j		syscall_getpriority                  # 141
		j		syscall_reboot                       # 142
		j		syscall_setregid                     # 143
		j		syscall_setgid                       # 144
		j		syscall_setreuid                     # 145
		j		syscall_setuid                       # 146
		j		syscall_setresuid                    # 147
		j		syscall_getresuid                    # 148
		j		syscall_setresgid                    # 149
		j		syscall_getresgid                    # 150
		j		syscall_setfsuid                     # 151
		j		syscall_setfsgid                     # 152
		j		syscall_times                        # 153
		j		syscall_setpgid                      # 154
		j		syscall_getpgid                      # 155
		j		syscall_getsid                       # 156
		j		syscall_setsid                       # 157
		j		syscall_getgroups                    # 158
		j		syscall_setgroups                    # 159
		j		syscall_uname                        # 160
		j		syscall_sethostname                  # 161
		j		syscall_setdomainname                # 162
		j		syscall_unimp                        # 163
		j		syscall_unimp                        # 164
		j		syscall_getrusage                    # 165
		j		syscall_umask                        # 166
		j		syscall_prctl                        # 167
		j		syscall_getcpu                       # 168
		j		syscall_unimp                        # 169
		j		syscall_unimp                        # 170
		j		syscall_unimp                        # 171
		j		syscall_getpid                       # 172
		j		syscall_getppid                      # 173
		j		syscall_getuid                       # 174
		j		syscall_geteuid                      # 175
		j		syscall_getgid                       # 176
		j		syscall_getegid                      # 177
		j		syscall_gettid                       # 178
		j		syscall_sysinfo                      # 179
		j		syscall_mq_open                      # 180
		j		syscall_mq_unlink                    # 181
		j		syscall_unimp                        # 182
		j		syscall_unimp                        # 183
		j		syscall_mq_notify                    # 184
		j		syscall_mq_getsetattr                # 185
		j		syscall_msgget                       # 186
		j		syscall_msgctl                       # 187
		j		syscall_msgrcv                       # 188
		j		syscall_msgsnd                       # 189
		j		syscall_semget                       # 190
		j		syscall_semctl                       # 191
		j		syscall_unimp                        # 192
		j		syscall_semop                        # 193
		j		syscall_shmget                       # 194
		j		syscall_shmctl                       # 195
		j		syscall_shmat                        # 196
		j		syscall_shmdt                        # 197
		j		syscall_socket                       # 198
		j		syscall_socketpair                   # 199
		j		syscall_bind                         # 200
		j		syscall_listen                       # 201
		j		syscall_accept                       # 202
		j		syscall_connect                      # 203
		j		syscall_getsockname                  # 204
		j		syscall_getpeername                  # 205
		j		syscall_sendto                       # 206
		j		syscall_recvfrom                     # 207
		j		syscall_setsockopt                   # 208
		j		syscall_getsockopt                   # 209
		j		syscall_shutdown                     # 210
		j		syscall_sendmsg                      # 211
		j		syscall_recvmsg                      # 212
		j		syscall_readahead                    # 213
		j		syscall_brk                          # 214
		j		syscall_munmap                       # 215
		j		syscall_mremap                       # 216
		j		syscall_add_key                      # 217
		j		syscall_request_key                  # 218
		j		syscall_keyctl                       # 219
		j		syscall_clone                        # 220
		j		syscall_execve                       # 221
		j		syscall_mmap2                        # 222
		j		syscall_fadvise64_64                 # 223
		j		syscall_swapon                       # 224
		j		syscall_swapoff                      # 225
		j		syscall_mprotect                     # 226
		j		syscall_msync                        # 227
		j		syscall_mlock                        # 228
		j		syscall_munlock                      # 229
		j		syscall_mlockall                     # 230
		j		syscall_munlockall                   # 231
		j		syscall_mincore                      # 232
		j		syscall_madvise                      # 233
		j		syscall_remap_file_pages             # 234
		j		syscall_mbind                        # 235
		j		syscall_get_mempolicy                # 236
		j		syscall_set_mempolicy                # 237
		j		syscall_migrate_pages                # 238
		j		syscall_move_pages                   # 239
		j		syscall_rt_tgsigqueueinfo            # 240
		j		syscall_perf_event_open              # 241
		j		syscall_accept4                      # 242
		j		syscall_unimp                        # 243
		j		syscall_unimp                        # 244
		j		syscall_unimp                        # 245
		j		syscall_unimp                        # 246
		j		syscall_unimp                        # 247
		j		syscall_unimp                        # 248
		j		syscall_unimp                        # 249
		j		syscall_unimp                        # 250
		j		syscall_unimp                        # 251
		j		syscall_unimp                        # 252
		j		syscall_unimp                        # 253
		j		syscall_unimp                        # 254
		j		syscall_unimp                        # 255
		j		syscall_unimp                        # 256
		j		syscall_unimp                        # 257
		j		syscall_unimp                        # 258
		j		syscall_riscv_flush_icache           # 259
		j		syscall_unimp                        # 260
		j		syscall_prlimit64                    # 261
		j		syscall_fanotify_init                # 262
		j		syscall_fanotify_mark                # 263
		j		syscall_name_to_handle_at            # 264
		j		syscall_open_by_handle_at            # 265
		j		syscall_unimp                        # 266
		j		syscall_syncfs                       # 267
		j		syscall_setns                        # 268
		j		syscall_sendmmsg                     # 269
		j		syscall_process_vm_readv             # 270
		j		syscall_process_vm_writev            # 271
		j		syscall_kcmp                         # 272
		j		syscall_finit_module                 # 273
		j		syscall_sched_setattr                # 274
		j		syscall_sched_getattr                # 275
		j		syscall_renameat2                    # 276
		j		syscall_seccomp                      # 277
		j		syscall_getrandom                    # 278
		j		syscall_memfd_create                 # 279
		j		syscall_bpf                          # 280
		j		syscall_execveat                     # 281
		j		syscall_userfaultfd                  # 282
		j		syscall_membarrier                   # 283
		j		syscall_mlock2                       # 284
		j		syscall_copy_file_range              # 285
		j		syscall_preadv2                      # 286
		j		syscall_pwritev2                     # 287
		j		syscall_pkey_mprotect                # 288
		j		syscall_pkey_alloc                   # 289
		j		syscall_pkey_free                    # 290
		j		syscall_statx                        # 291
		j		syscall_unimp                        # 292
		j		syscall_rseq                         # 293
		j		syscall_kexec_file_load              # 294
		j		syscall_unimp                        # 295
		j		syscall_unimp                        # 296
		j		syscall_unimp                        # 297
		j		syscall_unimp                        # 298
		j		syscall_unimp                        # 299
		j		syscall_unimp                        # 300
		j		syscall_unimp                        # 301
		j		syscall_unimp                        # 302
		j		syscall_unimp                        # 303
		j		syscall_unimp                        # 304
		j		syscall_unimp                        # 305
		j		syscall_unimp                        # 306
		j		syscall_unimp                        # 307
		j		syscall_unimp                        # 308
		j		syscall_unimp                        # 309
		j		syscall_unimp                        # 310
		j		syscall_unimp                        # 311
		j		syscall_unimp                        # 312
		j		syscall_unimp                        # 313
		j		syscall_unimp                        # 314
		j		syscall_unimp                        # 315
		j		syscall_unimp                        # 316
		j		syscall_unimp                        # 317
		j		syscall_unimp                        # 318
		j		syscall_unimp                        # 319
		j		syscall_unimp                        # 320
		j		syscall_unimp                        # 321
		j		syscall_unimp                        # 322
		j		syscall_unimp                        # 323
		j		syscall_unimp                        # 324
		j		syscall_unimp                        # 325
		j		syscall_unimp                        # 326
		j		syscall_unimp                        # 327
		j		syscall_unimp                        # 328
		j		syscall_unimp                        # 329
		j		syscall_unimp                        # 330
		j		syscall_unimp                        # 331
		j		syscall_unimp                        # 332
		j		syscall_unimp                        # 333
		j		syscall_unimp                        # 334
		j		syscall_unimp                        # 335
		j		syscall_unimp                        # 336
		j		syscall_unimp                        # 337
		j		syscall_unimp                        # 338
		j		syscall_unimp                        # 339
		j		syscall_unimp                        # 340
		j		syscall_unimp                        # 341
		j		syscall_unimp                        # 342
		j		syscall_unimp                        # 343
		j		syscall_unimp                        # 344
		j		syscall_unimp                        # 345
		j		syscall_unimp                        # 346
		j		syscall_unimp                        # 347
		j		syscall_unimp                        # 348
		j		syscall_unimp                        # 349
		j		syscall_unimp                        # 350
		j		syscall_unimp                        # 351
		j		syscall_unimp                        # 352
		j		syscall_unimp                        # 353
		j		syscall_unimp                        # 354
		j		syscall_unimp                        # 355
		j		syscall_unimp                        # 356
		j		syscall_unimp                        # 357
		j		syscall_unimp                        # 358
		j		syscall_unimp                        # 359
		j		syscall_unimp                        # 360
		j		syscall_unimp                        # 361
		j		syscall_unimp                        # 362
		j		syscall_unimp                        # 363
		j		syscall_unimp                        # 364
		j		syscall_unimp                        # 365
		j		syscall_unimp                        # 366
		j		syscall_unimp                        # 367
		j		syscall_unimp                        # 368
		j		syscall_unimp                        # 369
		j		syscall_unimp                        # 370
		j		syscall_unimp                        # 371
		j		syscall_unimp                        # 372
		j		syscall_unimp                        # 373
		j		syscall_unimp                        # 374
		j		syscall_unimp                        # 375
		j		syscall_unimp                        # 376
		j		syscall_unimp                        # 377
		j		syscall_unimp                        # 378
		j		syscall_unimp                        # 379
		j		syscall_unimp                        # 380
		j		syscall_unimp                        # 381
		j		syscall_unimp                        # 382
		j		syscall_unimp                        # 383
		j		syscall_unimp                        # 384
		j		syscall_unimp                        # 385
		j		syscall_unimp                        # 386
		j		syscall_unimp                        # 387
		j		syscall_unimp                        # 388
		j		syscall_unimp                        # 389
		j		syscall_unimp                        # 390
		j		syscall_unimp                        # 391
		j		syscall_unimp                        # 392
		j		syscall_unimp                        # 393
		j		syscall_unimp                        # 394
		j		syscall_unimp                        # 395
		j		syscall_unimp                        # 396
		j		syscall_unimp                        # 397
		j		syscall_unimp                        # 398
		j		syscall_unimp                        # 399
		j		syscall_unimp                        # 400
		j		syscall_unimp                        # 401
		j		syscall_unimp                        # 402
		j		syscall_clock_gettime64              # 403
		j		syscall_clock_settime64              # 404
		j		syscall_clock_adjtime64              # 405
		j		syscall_clock_getres_time64          # 406
		j		syscall_clock_nanosleep_time64       # 407
		j		syscall_timer_gettime64              # 408
		j		syscall_timer_settime64              # 409
		j		syscall_timerfd_gettime64            # 410
		j		syscall_timerfd_settime64            # 411
		j		syscall_utimensat_time64             # 412
		j		syscall_pselect6_time64              # 413
		j		syscall_ppoll_time64                 # 414
		j		syscall_unimp                        # 415
		j		syscall_io_pgetevents_time64         # 416
		j		syscall_recvmmsg_time64              # 417
		j		syscall_mq_timedsend_time64          # 418
		j		syscall_mq_timedreceive_time64       # 419
		j		syscall_semtimedop_time64            # 420
		j		syscall_rt_sigtimedwait_time64       # 421
		j		syscall_futex_time64                 # 422
		j		syscall_sched_rr_get_interval_time64 # 423
		j		syscall_pidfd_send_signal            # 424
		j		syscall_io_uring_setup               # 425
		j		syscall_io_uring_enter               # 426
		j		syscall_io_uring_register            # 427
		j		syscall_open_tree                    # 428
		j		syscall_move_mount                   # 429
		j		syscall_fsopen                       # 430
		j		syscall_fsconfig                     # 431
		j		syscall_fsmount                      # 432
		j		syscall_fspick                       # 433
		j		syscall_pidfd_open                   # 434
		j		syscall_unimp                        # 435
		j		syscall_close_range                  # 436
		j		syscall_openat2                      # 437
		j		syscall_pidfd_getfd                  # 438
		j		syscall_faccessat2                   # 439
		j		syscall_process_madvise              # 440
		j		syscall_epoll_pwait2                 # 441
		j		syscall_mount_setattr                # 442
		j		syscall_unimp                        # 443
		j		syscall_landlock_create_ruleset      # 444
		j		syscall_landlock_add_rule            # 445
		j		syscall_landlock_restrict_self       # 446
		ret

syscall_read:	# 63
		li		t0, memio
		add		t0, t0, a0
		li		t2, 0			# t2 = loop counter

1:		beq		t2, a2, 2f
		lb		t1, 0(t0)		# t1 = current char
		sb		t1, 0(a1)
		addi	a1, a1, 1
		addi	t2, t2, 1
		addi	t1, t1, -10		# check for newline
		bnez	t1, 1b
2:		mv		a0, t2			# chars read
		ret

syscall_write:	# 64
		li		t0, memio
		add		t0, t0, a0
		mv		a0, a2			# remaining chars

1:		beqz	a2, 2f
		lb		t1, 0(a1)
		sb		t1, 0(t0)
		addi	a1, a1, 1
		addi	a2, a2, -1
		j		1b
2:		ret

syscall_exit:	# 93
		li		t0, memio_exit
		sb		a0, 0(t0)		# exit code

syscall_io_setup:
syscall_io_destroy:
syscall_io_submit:
syscall_io_cancel:
syscall_setxattr:
syscall_lsetxattr:
syscall_fsetxattr:
syscall_getxattr:
syscall_lgetxattr:
syscall_fgetxattr:
syscall_listxattr:
syscall_llistxattr:
syscall_flistxattr:
syscall_removexattr:
syscall_lremovexattr:
syscall_fremovexattr:
syscall_getcwd:
syscall_lookup_dcookie:
syscall_eventfd2:
syscall_epoll_create1:
syscall_epoll_ctl:
syscall_epoll_pwait:
syscall_dup:
syscall_dup3:
syscall_fcntl64:
syscall_inotify_init1:
syscall_inotify_add_watch:
syscall_inotify_rm_watch:
syscall_ioctl:
syscall_ioprio_set:
syscall_ioprio_get:
syscall_flock:
syscall_mknodat:
syscall_mkdirat:
syscall_unlinkat:
syscall_symlinkat:
syscall_linkat:
syscall_umount2:
syscall_mount:
syscall_pivot_root:
syscall_nfsservctl:
syscall_statfs64:
syscall_fstatfs64:
syscall_truncate64:
syscall_ftruncate64:
syscall_fallocate:
syscall_faccessat:
syscall_chdir:
syscall_fchdir:
syscall_chroot:
syscall_fchmod:
syscall_fchmodat:
syscall_fchownat:
syscall_fchown:
syscall_openat:
syscall_close:
syscall_vhangup:
syscall_pipe2:
syscall_quotactl:
syscall_getdents64:
syscall_readv:
syscall_writev:
syscall_pread64:
syscall_pwrite64:
syscall_preadv:
syscall_pwritev:
syscall_sendfile64:
syscall_signalfd4:
syscall_vmsplice:
syscall_splice:
syscall_tee:
syscall_readlinkat:
syscall_sync:
syscall_fsync:
syscall_fdatasync:
syscall_sync_file_range:
syscall_timerfd_create:
syscall_acct:
syscall_capget:
syscall_capset:
syscall_personality:
syscall_exit_group:
syscall_waitid:
syscall_set_tid_address:
syscall_unshare:
syscall_set_robust_list:
syscall_get_robust_list:
syscall_getitimer:
syscall_setitimer:
syscall_kexec_load:
syscall_init_module:
syscall_delete_module:
syscall_timer_create:
syscall_timer_getoverrun:
syscall_timer_delete:
syscall_syslog:
syscall_ptrace:
syscall_sched_setparam:
syscall_sched_setscheduler:
syscall_sched_getscheduler:
syscall_sched_getparam:
syscall_sched_setaffinity:
syscall_sched_getaffinity:
syscall_sched_yield:
syscall_sched_get_priority_max:
syscall_sched_get_priority_min:
syscall_restart_syscall:
syscall_kill:
syscall_tkill:
syscall_tgkill:
syscall_sigaltstack:
syscall_rt_sigsuspend:
syscall_rt_sigaction:
syscall_rt_sigprocmask:
syscall_rt_sigpending:
syscall_rt_sigqueueinfo:
syscall_rt_sigreturn:
syscall_setpriority:
syscall_getpriority:
syscall_reboot:
syscall_setregid:
syscall_setgid:
syscall_setreuid:
syscall_setuid:
syscall_setresuid:
syscall_getresuid:
syscall_setresgid:
syscall_getresgid:
syscall_setfsuid:
syscall_setfsgid:
syscall_times:
syscall_setpgid:
syscall_getpgid:
syscall_getsid:
syscall_setsid:
syscall_getgroups:
syscall_setgroups:
syscall_uname:
syscall_sethostname:
syscall_setdomainname:
syscall_getrusage:
syscall_umask:
syscall_prctl:
syscall_getcpu:
syscall_getpid:
syscall_getppid:
syscall_getuid:
syscall_geteuid:
syscall_getgid:
syscall_getegid:
syscall_gettid:
syscall_sysinfo:
syscall_mq_open:
syscall_mq_unlink:
syscall_mq_notify:
syscall_mq_getsetattr:
syscall_msgget:
syscall_msgctl:
syscall_msgrcv:
syscall_msgsnd:
syscall_semget:
syscall_semctl:
syscall_semop:
syscall_shmget:
syscall_shmctl:
syscall_shmat:
syscall_shmdt:
syscall_socket:
syscall_socketpair:
syscall_bind:
syscall_listen:
syscall_accept:
syscall_connect:
syscall_getsockname:
syscall_getpeername:
syscall_sendto:
syscall_recvfrom:
syscall_setsockopt:
syscall_getsockopt:
syscall_shutdown:
syscall_sendmsg:
syscall_recvmsg:
syscall_readahead:
syscall_brk:
syscall_munmap:
syscall_mremap:
syscall_add_key:
syscall_request_key:
syscall_keyctl:
syscall_clone:
syscall_execve:
syscall_mmap2:
syscall_fadvise64_64:
syscall_swapon:
syscall_swapoff:
syscall_mprotect:
syscall_msync:
syscall_mlock:
syscall_munlock:
syscall_mlockall:
syscall_munlockall:
syscall_mincore:
syscall_madvise:
syscall_remap_file_pages:
syscall_mbind:
syscall_get_mempolicy:
syscall_set_mempolicy:
syscall_migrate_pages:
syscall_move_pages:
syscall_rt_tgsigqueueinfo:
syscall_perf_event_open:
syscall_accept4:
syscall_riscv_flush_icache:
syscall_prlimit64:
syscall_fanotify_init:
syscall_fanotify_mark:
syscall_name_to_handle_at:
syscall_open_by_handle_at:
syscall_syncfs:
syscall_setns:
syscall_sendmmsg:
syscall_process_vm_readv:
syscall_process_vm_writev:
syscall_kcmp:
syscall_finit_module:
syscall_sched_setattr:
syscall_sched_getattr:
syscall_renameat2:
syscall_seccomp:
syscall_getrandom:
syscall_memfd_create:
syscall_bpf:
syscall_execveat:
syscall_userfaultfd:
syscall_membarrier:
syscall_mlock2:
syscall_copy_file_range:
syscall_preadv2:
syscall_pwritev2:
syscall_pkey_mprotect:
syscall_pkey_alloc:
syscall_pkey_free:
syscall_statx:
syscall_rseq:
syscall_kexec_file_load:
syscall_clock_gettime64:
syscall_clock_settime64:
syscall_clock_adjtime64:
syscall_clock_getres_time64:
syscall_clock_nanosleep_time64:
syscall_timer_gettime64:
syscall_timer_settime64:
syscall_timerfd_gettime64:
syscall_timerfd_settime64:
syscall_utimensat_time64:
syscall_pselect6_time64:
syscall_ppoll_time64:
syscall_io_pgetevents_time64:
syscall_recvmmsg_time64:
syscall_mq_timedsend_time64:
syscall_mq_timedreceive_time64:
syscall_semtimedop_time64:
syscall_rt_sigtimedwait_time64:
syscall_futex_time64:
syscall_sched_rr_get_interval_time64:
syscall_pidfd_send_signal:
syscall_io_uring_setup:
syscall_io_uring_enter:
syscall_io_uring_register:
syscall_open_tree:
syscall_move_mount:
syscall_fsopen:
syscall_fsconfig:
syscall_fsmount:
syscall_fspick:
syscall_pidfd_open:
syscall_close_range:
syscall_openat2:
syscall_pidfd_getfd:
syscall_faccessat2:
syscall_process_madvise:
syscall_epoll_pwait2:
syscall_mount_setattr:
syscall_landlock_create_ruleset:
syscall_landlock_add_rule:
syscall_landlock_restrict_self:
syscall_unimp:

		unimp					# this should never be reached
