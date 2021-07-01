.global _start

_start:	
		li		a0, 1		# 1 = StdOut
		la		a1, prompt	# load address
		li		a2, 17		# length
		jal		write

		li		a0, 0		# 0 = StdIn
		la		a1, name	# load address
		li		a2, 100		# length
		jal		read
		mv		t1, a0		# store length

		li		a0, 1		# 1 = StdOut
		la		a1, hello	# load address
		li		a2, 6		# length
		jal		write

		li		a0, 1		# 1 = StdOut
		la		a1, name	# load address
		mv		a2, t1		# length
		jal		write

		li		a0, 1		# 1 = StdOut
		la		a1, nl		# load address
		li		a2, 1		# length
		jal		write

		li		a0, 0		# return code
		jal		exit

# memio functions
.equ	memio, 0x10000000

read:	li		t0, memio			# t0 = memio
		add		t0, t0, a0
		li		t2, 0				# t2 = loop counter

rloop:	beq		t2, a2, rdone
		lb		t1, 0(t0)			# t1 = current char
		sb		t1, 0(a1)
		addi	a1, a1, 1
		addi	t2, t2, 1
		addi	t1, t1, -10			# check for newline
		bnez	t1, rloop
rdone:	mv		a0, t2
		ret

write:  li		t0, memio			# t0 = memio
		add		t0, t0, a0
		mv		a0, a2				# a2 = remaining chars

wloop:	beqz	a2, wdone
		lb		t1, 0(a1)			# t1 = current char
		sb		t1, 0(t0)
		addi	a1, a1, 1
		addi	a2, a2, -1
		j		wloop
wdone:	ret

exit:	li		a7, 93				# syscall exit
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
