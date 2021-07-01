.global _start

_start:	li		a0, 1				# 1 = StdOut
		la		a1, helloworld		# load address
		li		a2, 13				# length
		jal		write

		li		a0, 0				# return code
		jal		exit

# memio functions
.equ	memio, 0x10000000

write:  li		t0, memio			# t0 = memio
		add		t0, t0, a0
		mv		a0, a2				# a2 = remaining chars

loop:	beqz	a2, done
		lb		t1, 0(a1)			# t1 = current char
		sb		t1, 0(t0)
		addi	a1, a1, 1
		addi	a2, a2, -1
		j		loop
done:	ret

exit:	li		a7, 93				# syscall exit
		ecall

.data
helloworld:	.ascii "Hello World!\n..."

