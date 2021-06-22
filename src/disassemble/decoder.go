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

func decodeInstruction(section *Section, addr, raw uint32, sym string) Instruction {
	i := &instruction{
		section: section,
		addr:    addr,
		sym:     sym,
		raw:     raw,
	}

	if i.raw == 0 {
		return &Unimp{i}
	}

	switch i.Opcode() {
	case opcodeLUI: // U-type
		i.mnemonic = "lui"

		return &Lui{i}
	case opcodeAUIPC: // U-type
		i.mnemonic = "auipc"

		return &Auipc{i}
	case opcodeJAL: // J-type
		i.mnemonic = "jal"

		return &Jal{i}
	case opcodeJALR: // I-type
		i.mnemonic = "jalr"

		return &Jalr{i}
	case opcodeBRANCH: // B-type
		switch i.Func3() {
		case func3BEQ:
			i.mnemonic = "beq"

			return &Beq{&branch{i}}
		case func3BNE:
			i.mnemonic = "bne"

			return &Bne{&branch{i}}
		case func3BLT:
			i.mnemonic = "blt"

			return &Blt{&branch{i}}
		case func3BGE:
			i.mnemonic = "bge"

			return &Bge{&branch{i}}
		case func3BLTU:
			i.mnemonic = "bltu"

			return &Bltu{&branch{i}}
		case func3BGEU:
			i.mnemonic = "bgeu"

			return &Bgeu{&branch{i}}
		}
	case opcodeLOAD: // I-type
		switch i.Func3() {
		case func3LB:
			i.mnemonic = "lb"

			return &Lb{&load{i}}
		case func3LH:
			i.mnemonic = "lh"

			return &Lh{&load{i}}
		case func3LW:
			i.mnemonic = "lw"

			return &Lw{&load{i}}
		case func3LBU:
			i.mnemonic = "lbu"

			return &Lbu{&load{i}}
		case func3LHU:
			i.mnemonic = "lhu"

			return &Lhu{&load{i}}
		}
	case opcodeSTORE: // S-type
		switch i.Func3() {
		case func3SB:
			i.mnemonic = "sb"

			return &Sb{&store{i}}
		case func3SH:
			i.mnemonic = "sh"

			return &Sh{&store{i}}
		case func3SW:
			i.mnemonic = "sw"

			return &Sw{&store{i}}
		}
	case opcodeSYSTEM: // I-type
		switch i.Func3() {
		case func3ECALL: // or func3EBREAK
			if i.Imm()>>20 == 1 {
				i.mnemonic = "ebreak"

				return &Ebreak{&system{i}}
			}

			i.mnemonic = "ecall"

			return &Ecall{&system{i}}
		case func3CSRRW:
			i.mnemonic = "csrrw"

			return &Csrrw{&system{i}}
		case func3CSRRS:
			i.mnemonic = "csrrs"

			return &Csrrs{&system{i}}
		case func3CSRRC:
			i.mnemonic = "scrrc"

			return &Csrrc{&system{i}}
		case func3CSRRWI:
			i.mnemonic = "csrrwi"

			return &Csrrwi{&system{i}}
		case func3CSRRSI:
			i.mnemonic = "csrrsi"

			return &Csrrsi{&system{i}}
		case func3CSRRCI:
			i.mnemonic = "csrrci"

			return &Csrrci{&system{i}}
		}
	case opcodeMISC_MEM:
		switch i.Func3() {
		case func3FENCE:
			i.mnemonic = "fence"

			return &Fence{i}
		}
	case opcodeOP_IMM: // I-type
		if i.Func7() == 0b0100000 && i.Func3() == func3SRAI {
			switch i.Func3() {
			case func3SRAI:
				i.mnemonic = "srai"

				return &Srai{&opImm{i}}
			}
		}

		switch i.Func3() {
		case func3ADDI:
			i.mnemonic = "addi"

			return &Addi{&opImm{i}}
		case func3SLTUI:
			i.mnemonic = "sltui"

			return &Sltui{&opImm{i}}
		case func3XORI:
			i.mnemonic = "xori"

			return &Xori{&opImm{i}}
		case func3SLLI:
			i.mnemonic = "slli"

			return &Slli{&opImm{i}}
		case func3SRLI:
			i.mnemonic = "srli"

			return &Srli{&opImm{i}}
		case func3ORI:
			i.mnemonic = "ori"

			return &Ori{&opImm{i}}
		case func3ANDI:
			i.mnemonic = "andi"

			return &Andi{&opImm{i}}
		}
	case opcodeOP: // R-type
		if i.Func7() == 0b0100000 {
			switch i.Func3() {
			case func3SUB:
				i.mnemonic = "sub"

				return &Sub{&opReg{i}}
			case func3SRA:
				i.mnemonic = "sra"

				return &Sra{&opReg{i}}
			}
		}

		switch i.Func3() {
		case func3ADD: // 000
			i.mnemonic = "add"

			return &Add{&opReg{i}}
		case func3SLT: // 010
			i.mnemonic = "slt"

			return &Slt{&opReg{i}}
		case func3SLTU: // 011
			i.mnemonic = "sltu"

			return &Sltu{&opReg{i}}
		case func3XOR: // 100
			i.mnemonic = "xor"

			return &Xor{&opReg{i}}
		case func3SRL: // 101
			i.mnemonic = "srl"

			return &Srl{&opReg{i}}
		case func3OR: // 110
			i.mnemonic = "or"

			return &Or{&opReg{i}}
		case func3AND: // 111
			i.mnemonic = "and"

			return &And{&opReg{i}}
		}
	}

	return i
}
