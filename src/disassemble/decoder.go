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

func (o Opcode) decodeImm(raw uint32) int32 {
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

	switch o {
	case opcodeLUI, opcodeAUIPC:
		// U-type
		return int32(sign_extend(bits(31, 12), 32))
	case opcodeJAL:
		// J-type
		return int32(sign_extend(bits(32, 31)<<20|bits(30, 21)<<1|bits(21, 20)<<11|bits(19, 12)<<12, 21))
	case opcodeBRANCH:
		// B-type
		return int32(sign_extend(bits(32, 31)<<12|bits(30, 25)<<5|bits(11, 8)<<1|bits(8, 7)<<11, 13))
	case opcodeSTORE:
		// S-type
		return int32(sign_extend(bits(31, 25)<<5|bits(11, 7), 12))
	case opcodeJALR, opcodeLOAD, opcodeOP_IMM, opcodeSYSTEM:
		// I-type
		return int32(sign_extend(bits(31, 20), 12))
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
	func3FENCE Func3 = 0b000
	// System: Zifencei extension
	func3FENCEI Func3 = 0b001
	// System: Zicsr extension
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

	set := func(i *instruction, m string) *instruction {
		i.mnemonic = m

		return i
	}

	branch := func(i *instruction, m string) *Branch { return &Branch{set(i, m)} }
	load := func(i *instruction, m string) *Load { return &Load{set(i, m)} }
	store := func(i *instruction, m string) *Store { return &Store{set(i, m)} }
	misc := func(i *instruction, m string) *Misc { return &Misc{set(i, m)} }
	system := func(i *instruction, m string) *System { return &System{set(i, m)} }
	opImm := func(i *instruction, m string) *OpImm { return &OpImm{set(i, m)} }
	opReg := func(i *instruction, m string) *OpReg { return &OpReg{set(i, m)} }

	switch i.Opcode() {
	case opcodeLUI:
		return &Lui{set(i, "lui")}
	case opcodeAUIPC:
		return &Auipc{set(i, "auipc")}
	case opcodeJAL:
		return &Jal{set(i, "jal")}
	case opcodeJALR:
		return &Jalr{set(i, "jalr")}
	case opcodeBRANCH:
		switch i.Func3() {
		case func3BEQ:
			return &Beq{branch(i, "beq")}
		case func3BNE:
			return &Bne{branch(i, "bne")}
		case func3BLT:
			return &Blt{branch(i, "blt")}
		case func3BGE:
			return &Bge{branch(i, "bge")}
		case func3BLTU:
			return &Bltu{branch(i, "bltu")}
		case func3BGEU:
			return &Bgeu{branch(i, "bgeu")}
		}
	case opcodeLOAD:
		switch i.Func3() {
		case func3LB:
			return &Lb{load(i, "lb")}
		case func3LH:
			return &Lh{load(i, "lh")}
		case func3LW:
			return &Lw{load(i, "lw")}
		case func3LBU:
			return &Lbu{load(i, "lbu")}
		case func3LHU:
			return &Lhu{load(i, "lhu")}
		}
	case opcodeSTORE:
		switch i.Func3() {
		case func3SB:
			return &Sb{store(i, "sb")}
		case func3SH:
			return &Sh{store(i, "sh")}
		case func3SW:
			return &Sw{store(i, "sw")}
		}
	case opcodeSYSTEM:
		switch i.Func3() {
		case func3ECALL:
			if i.bits(31, 20) == 1 {
				return &Ebreak{system(i, "ebreak")}
			}

			return &Ecall{system(i, "ecall")}
		case func3CSRRW:
			return &Csrrw{system(i, "csrrw")}
		case func3CSRRS:
			return &Csrrs{system(i, "csrrs")}
		case func3CSRRC:
			return &Csrrc{system(i, "csrrc")}
		case func3CSRRWI:
			return &Csrrwi{system(i, "csrrwi")}
		case func3CSRRSI:
			return &Csrrsi{system(i, "csrrsi")}
		case func3CSRRCI:
			return &Csrrci{system(i, "csrrci")}
		}
	case opcodeMISC_MEM:
		switch i.Func3() {
		case func3FENCE:
			return &Fence{misc(i, "fence")}
		case func3FENCEI:
			return &Fencei{misc(i, "fence.i")}
		}
	case opcodeOP_IMM:
		if i.Func7() == 0b0100000 && i.Func3() == func3SRAI {
			switch i.Func3() {
			case func3SRAI:
				return &Srai{opImm(i, "srai")}
			}
		}

		switch i.Func3() {
		case func3ADDI:
			return &Addi{opImm(i, "addi")}
		case func3SLTUI:
			return &Sltiu{opImm(i, "sltiu")}
		case func3XORI:
			return &Xori{opImm(i, "xori")}
		case func3SLLI:
			return &Slli{opImm(i, "slli")}
		case func3SRLI:
			return &Srli{opImm(i, "srli")}
		case func3ORI:
			return &Ori{opImm(i, "ori")}
		case func3ANDI:
			return &Andi{opImm(i, "andi")}
		}
	case opcodeOP:
		if i.Func7() == 0b0100000 {
			switch i.Func3() {
			case func3SUB:
				return &Sub{opReg(i, "sub")}
			case func3SRA:
				return &Sra{opReg(i, "sra")}
			}
		}

		switch i.Func3() {
		case func3ADD:
			return &Add{opReg(i, "add")}
		case func3SLT:
			return &Slt{opReg(i, "slt")}
		case func3SLTU:
			return &Sltu{opReg(i, "sltu")}
		case func3XOR:
			return &Xor{opReg(i, "xor")}
		case func3SRL:
			return &Srl{opReg(i, "srl")}
		case func3OR:
			return &Or{opReg(i, "or")}
		case func3AND:
			return &And{opReg(i, "and")}
		}
	}

	return &Unimp{set(i, "unimp")}
}
