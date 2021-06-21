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

type Func3 uint8

const (
	// Arith
	func3ADD  Func3 = 0b000
	func3SUB  Func3 = 0b000
	func3ADDI Func3 = 0b000

	func3SLLI  Func3 = 0b001
	func3SLT   Func3 = 0b010
	func3SLTU  Func3 = 0b011
	func3SLTIU Func3 = 0b011

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

	// Misc/System
	func3ECALL  Func3 = 0b000
	func3CSRRW  Func3 = 0b001
	func3CSRRS  Func3 = 0b010
	func3CSRRC  Func3 = 0b011
	func3CSRRWI Func3 = 0b101
	func3CSRRSI Func3 = 0b110
	func3CSRRCI Func3 = 0b111
)

func (f3 Func3) Branch() string {
	switch f3 {
	case 0b000:
		return "beq"
	case 0b001:
		return "bne"
	case 0b100:
		return "blt"
	case 0b101:
		return "bge"
	case 0b110:
		return "bltu"
	case 0b111:
		return "bgeu"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

func (f3 Func3) Load() string {
	switch f3 {
	case 0b000:
		return "lb"
	case 0b001:
		return "lh"
	case 0b010:
		return "lw"
	case 0b100:
		return "lbu"
	case 0b101:
		return "lhu"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

func (f3 Func3) Store() string {
	switch f3 {
	case 0b000:
		return "sb"
	case 0b001:
		return "sh"
	case 0b010:
		return "sw"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

func (f3 Func3) Arith(imm, alt bool) string {
	switch f3 {
	case 0b000:
		if alt {
			return "sub"
		}

		if imm {
			return "addi"
		}

		return "add"
	case 0b001:
		return "slli"
	case 0b010:
		return "slt"
	case 0b011:
		if imm {
			return "sltui"
		}

		return "sltu"
	case 0b100:
		if imm {
			return "xori"
		}

		return "xor"
	case 0b101:
		if alt {
			if imm {
				return "srai"
			}

			return "sra"
		}

		if imm {
			return "srli"
		}

		return "srl"
	case 0b110:
		if imm {
			return "ori"
		}

		return "or"
	case 0b111:
		if imm {
			return "andi"
		}

		return "and"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

func (f3 Func3) Misc() string {
	switch f3 {
	case 0b000:
		return "fence"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

func (f3 Func3) System(imm uint32) string {
	switch f3 {
	case 0b000:
		if imm>>20 == 1 {
			return "ebreak"
		}

		return "ecall"
	case 0b001:
		return "csrrw"
	case 0b010:
		return "csrrs"
	case 0b011:
		return "scrrc"
	case 0b101:
		return "csrrwi"
	case 0b110:
		return "csrrsi"
	case 0b111:
		return "csrrci"
	}

	panic(fmt.Sprintf("f3 out of range: %#x", f3))
}

type Register uint8

func (r Register) String() string {
	switch r {
	case 0:
		return "zero"
	case 1:
		return "ra"
	case 2:
		return "sp"
	case 3:
		return "gp"
	case 4:
		return "tp"
	case 5:
		return "t0"
	case 6:
		return "t1"
	case 7:
		return "t2"
	case 8:
		return "s0"
	case 9:
		return "s1"
	case 10:
		return "a0"
	case 11:
		return "a1"
	case 12:
		return "a2"
	case 13:
		return "a3"
	case 14:
		return "a4"
	case 15:
		return "a5"
	case 16:
		return "a6"
	case 17:
		return "a7"
	case 18:
		return "s2"
	case 19:
		return "s3"
	case 20:
		return "s4"
	case 21:
		return "s5"
	case 22:
		return "s6"
	case 23:
		return "s7"
	case 24:
		return "s8"
	case 25:
		return "s9"
	case 26:
		return "s10"
	case 27:
		return "s11"
	case 28:
		return "t3"
	case 29:
		return "t4"
	case 30:
		return "t5"
	case 31:
		return "t6"
	case 32:
		return "t7"
	}

	panic(fmt.Sprintf("illegal register %#x", uint8(r)))
}

func decode(s *Section, addr, raw uint32) string {
	if raw == 0 {
		return "unimp"
	}

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

	opcode := Opcode(bits(6, 0))
	imm_b := sign_extend(bits(32, 31)<<12|bits(30, 25)<<5|bits(11, 8)<<1|bits(8, 7)<<11, 13)
	imm_i := sign_extend(bits(31, 20), 12)
	imm_j := sign_extend(bits(32, 31)<<20|bits(30, 21)<<1|bits(21, 20)<<11|bits(19, 12)<<12, 21)
	imm_s := sign_extend(bits(31, 25)<<5|bits(11, 7), 12)
	imm_u := sign_extend(bits(31, 12), 32)

	// registers
	rd := Register(bits(11, 7))
	rs1 := Register(bits(19, 15))
	rs2 := Register(bits(24, 20))

	fn3 := Func3(bits(14, 12))
	fn7 := bits(31, 25)

	switch opcode {
	case opcodeLUI: // U
		return fmt.Sprintf("lui %s, %#x",
			rd, imm_u)
	case opcodeAUIPC: // U
		return fmt.Sprintf("auipc %s, %#x",
			rd, imm_u)
	case opcodeJAL: // J
		addr += imm_j

		return fmt.Sprintf("jal %08x %s",
			addr, s.NearestSymbol(addr))
	case opcodeJALR: // I
		return fmt.Sprintf("jalr %s, %#x",
			rs1, imm_i)
	case opcodeBRANCH: // B
		addr += imm_b

		return fmt.Sprintf("%s %s, %s, %08x %s",
			fn3.Branch(), rs1, rs2, addr, s.NearestSymbol(addr))
	case opcodeLOAD: // I
		return fmt.Sprintf("%s %s, %s+%#x",
			fn3.Load(), rd, rs1, imm_i)
	case opcodeSTORE: // S
		return fmt.Sprintf("%s %s+%#x, %s",
			fn3.Store(), rs1, imm_s, rs2)
	case opcodeOP_IMM: // I
		return fmt.Sprintf("%s %s, %s, %#x",
			fn3.Arith(true, fn7 == 0b0100000 && fn3 == func3SRAI),
			rd, rs1, imm_i)
	case opcodeOP: // R
		return fmt.Sprintf("%s %s, %s, %s",
			fn3.Arith(false, fn7 == 0b0100000), rd, rs1, rs2)
	case opcodeMISC_MEM:
		return fmt.Sprintf("%s",
			fn3.Misc())
	case opcodeSYSTEM: // I
		return fmt.Sprintf("%s",
			fn3.System(imm_i))
	}

	panic(fmt.Sprintf("unknown opcode: %v", opcode))
}
