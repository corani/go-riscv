.global _start

_start:	li		a0, 1				# 1 = StdOut
		la		a1, helloworld		# load address
		li		a2, 13				# length
		li		a7, 64				# syscall write
		ecall

		li		a0, 0				# return code
		li		a7, 93				# syscall exit
		ecall

.data
helloworld:	.ascii "Hello World!\n..."
