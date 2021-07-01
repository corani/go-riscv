.global _start

_start:	la		a1, helloworld		# load address
		li		a2, 13				# length
		li		a7, 0x10000001		# StdOut

loop:	beqz	a2, done
		lb		a3, 0(a1)
		sb		a3, 0(a7)
		addi	a1, a1, 1
		addi	a2, a2, -1
		j		loop

done:	li		a0, 0				# return code
		li		a7, 93				# syscall exit
		ecall

.data
helloworld:	.ascii "Hello World!\n..."

