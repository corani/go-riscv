.global _start

_start:	
		li		a7, 64		# syscall write
		li		a0, 1		# 1 = StdOut
		la		a1, prompt	# load address
		li		a2, 17		# length
		ecall

		li		a7, 63		# syscall read
		li		a0, 0		# 0 = StdIn
		la		a1, name	# load address
		li		a2, 100		# length
		ecall
		mv		t1, a0		# store length

		li		a7, 64		# syscall write
		li		a0, 1		# 1 = StdOut
		la		a1, hello	# load address
		li		a2, 6		# length
		ecall

		li		a0, 1		# 1 = StdOut
		la		a1, name	# load address
		mv		a2, t1		# length
		ecall

		li		a0, 1		# 1 = StdOut
		la		a1, nl		# load address
		li		a2, 1		# length
		ecall

		li		a0, 0		# return code
		li		a7, 93		# syscall exit
		ecall

.data

.align 4
prompt:	
		.ascii "Enter your name: "
.align 4
hello:	
		.ascii "Hello "
.align 4
nl:		
		.ascii "\n"
.align 4
name:	
		.zero  100
