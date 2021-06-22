package main

import (
	"fmt"
)

type Opcode uint8

const (
	opcodeLUI      Opcode = 0b00110111
	opcodeLOAD     Opcode = 0b00000011
	opcodeSTORE    Opcode = 0b00100011
	opcodeAUIPC    Opcode = 0b00010111
	opcodeBRANCH   Opcode = 0b01100011
	opcodeJAL      Opcode = 0b01101111
	opcodeJALR     Opcode = 0b01100111
	opcodeOP_IMM   Opcode = 0b00010011
	opcodeOP       Opcode = 0b00110011
	opcodeMISC_MEM Opcode = 0b00001111
	opcodeSYSTEM   Opcode = 0b01110011
)

func (o Opcode) decodeImm(raw uint32) uint32 {
	// bits gets the bits [s..e] from raw (shifted down)
	bits := func(s, e int) uint32 {
		return (raw >> e) & ((1 << (s - e + 1)) - 1)
	}

	sign_extend := func(x uint32, l int) uint32 {
		if x>>(l-1) == 1 {
			return -((1 << l) - x)
		}

		return x
	}

	imm_b := sign_extend(bits(32, 31)<<12|bits(30, 25)<<5|bits(11, 8)<<1|bits(8, 7)<<11, 13)
	imm_i := sign_extend(bits(31, 20), 12)
	imm_j := sign_extend(bits(32, 31)<<20|bits(30, 21)<<1|bits(21, 20)<<11|bits(19, 12)<<12, 21)
	imm_s := sign_extend(bits(31, 25)<<5|bits(11, 7), 12)
	imm_u := sign_extend(bits(31, 12), 32)

	switch o {
	case opcodeLUI, opcodeAUIPC:
		// U-type
		return imm_u
	case opcodeJAL:
		// J-type
		return imm_j
	case opcodeBRANCH:
		// B-type
		return imm_b
	case opcodeSTORE:
		// S-type
		return imm_s
	case opcodeJALR, opcodeLOAD, opcodeOP_IMM, opcodeSYSTEM:
		// I-type
		return imm_i
	case opcodeOP, opcodeMISC_MEM:
		// No immediate value
		return 0
	}

	panic(fmt.Sprintf("unknown opcode: %v", o))
}

type Func3 uint8

const (
	// Arith
	func3ADD  Func3 = 0b000
	func3SUB  Func3 = 0b000
	func3ADDI Func3 = 0b000

	func3SLLI  Func3 = 0b001
	func3SLT   Func3 = 0b010
	func3SLTU  Func3 = 0b011
	func3SLTUI Func3 = 0b011

	func3XOR  Func3 = 0b100
	func3XORI Func3 = 0b100
	func3SRL  Func3 = 0b101
	func3SRLI Func3 = 0b101
	func3SRA  Func3 = 0b101
	func3SRAI Func3 = 0b101
	func3OR   Func3 = 0b110
	func3ORI  Func3 = 0b110
	func3AND  Func3 = 0b111
	func3ANDI Func3 = 0b111

	// Branch
	func3BEQ  Func3 = 0b000
	func3BNE  Func3 = 0b001
	func3BLT  Func3 = 0b100
	func3BGE  Func3 = 0b101
	func3BLTU Func3 = 0b110
	func3BGEU Func3 = 0b111

	// Load/Store
	func3LB  Func3 = 0b000
	func3SB  Func3 = 0b000
	func3LH  Func3 = 0b001
	func3SH  Func3 = 0b001
	func3LW  Func3 = 0b010
	func3SW  Func3 = 0b010
	func3LBU Func3 = 0b100
	func3LHU Func3 = 0b101

	// Misc
	func3ECALL  Func3 = 0b000
	func3EBREAK Func3 = 0b000

	// System
	func3FENCE  Func3 = 0b000
	func3CSRRW  Func3 = 0b001
	func3CSRRS  Func3 = 0b010
	func3CSRRC  Func3 = 0b011
	func3CSRRWI Func3 = 0b101
	func3CSRRSI Func3 = 0b110
	func3CSRRCI Func3 = 0b111
)

func (f3 Func3) Branch() string {
	switch f3 {
	case func3BEQ:
		return "beq"
	case func3BNE:
		return "bne"
	case func3BLT:
		return "blt"
	case func3BGE:
		return "bge"
	case func3BLTU:
		return "bltu"
	case func3BGEU:
		return "bgeu"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

func (f3 Func3) Load() string {
	switch f3 {
	case func3LB:
		return "lb"
	case func3LH:
		return "lh"
	case func3LW:
		return "lw"
	case func3LBU:
		return "lbu"
	case func3LHU:
		return "lhu"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

func (f3 Func3) Store() string {
	switch f3 {
	case func3SB:
		return "sb"
	case func3SH:
		return "sh"
	case func3SW:
		return "sw"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

func (f3 Func3) Arith(imm, alt bool) string {
	switch {
	case alt && imm:
		switch f3 {
		case func3SRAI:
			return "srai"
		}
	case alt:
		switch f3 {
		case func3SUB:
			return "sub"
		case func3SRA:
			return "sra"
		}
	case imm:
		switch f3 {
		case func3ADDI:
			return "addi"
		case func3SLTUI:
			return "sltui"
		case func3XORI:
			return "xori"
		case func3SLLI:
			return "slli"
		case func3SRLI:
			return "srli"
		case func3ORI:
			return "ori"
		case func3ANDI:
			return "andi"
		}
	default:
		switch f3 {
		case func3ADD: // 000
			return "add"
		case func3SLT: // 010
			return "slt"
		case func3SLTU: // 011
			return "sltu"
		case func3XOR: // 100
			return "xor"
		case func3SRL: // 101
			return "srl"
		case func3OR: // 110
			return "or"
		case func3AND: // 111
			return "and"
		}
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

func (f3 Func3) Misc() string {
	switch f3 {
	case func3FENCE:
		return "fence"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

func (f3 Func3) System(imm uint32) string {
	switch f3 {
	case func3ECALL: // or func3EBREAK
		if imm>>20 == 1 {
			return "ebreak"
		}

		return "ecall"
	case func3CSRRW:
		return "csrrw"
	case func3CSRRS:
		return "csrrs"
	case func3CSRRC:
		return "scrrc"
	case func3CSRRWI:
		return "csrrwi"
	case func3CSRRSI:
		return "csrrsi"
	case func3CSRRCI:
		return "csrrci"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

type Register uint8

func (r Register) String() string {
	names := map[Register]string{
		0: "zero",
		1: "ra", 2: "sp", 3: "gp", 4: "tp",
		5: "t0", 6: "t1", 7: "t2", 28: "t3",
		29: "t4", 30: "t5", 31: "t6", 32: "t7",
		8: "s0", 9: "s1", 18: "s2", 19: "s3",
		20: "s4", 21: "s5", 22: "s6", 23: "s7",
		24: "s8", 25: "s9", 26: "s10", 27: "s11",
		10: "a0", 11: "a1", 12: "a2", 13: "a3",
		14: "a4", 15: "a5", 16: "a6", 17: "a7",
	}

	if v, ok := names[r]; ok {
		return v
	}

	panic(fmt.Sprintf("illegal register %#x", uint8(r)))
}
