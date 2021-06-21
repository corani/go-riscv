package main

import "fmt"

type Opcode uint8

const (
	opcodeLUI    Opcode = 0b00110111
	opcodeLOAD   Opcode = 0b00000011
	opcodeSTORE  Opcode = 0b00100011
	opcodeAUIPC  Opcode = 0b00010111
	opcodeBRANCH Opcode = 0b01100011
	opcodeJAL    Opcode = 0b01101111
	opcodeJALR   Opcode = 0b01100111
	opcodeIMM    Opcode = 0b00010011
	opcodeOP     Opcode = 0b00110011
	opcodeMISC   Opcode = 0b00001111
	opcodeSYSTEM Opcode = 0b01110011
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

func (c Opcode) String() string {
	switch c {
	case opcodeLUI:
		return "lui"
	case opcodeLOAD:
		return "load"
	case opcodeSTORE:
		return "store"
	case opcodeAUIPC:
		return "auipc"
	case opcodeBRANCH:
		return "branch"
	case opcodeJAL:
		return "jal"
	case opcodeJALR:
		return "jalr"
	case opcodeIMM:
		return "imm"
	case opcodeOP:
		return "op"
	case opcodeMISC:
		return "misc"
	case opcodeSYSTEM:
		return "system"
	}

	return fmt.Sprintf("?? %#x", c)
}

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

func (f3 Func3) System() string {
	switch f3 {
	case 0b000:
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

func (f3 Func3) String() string {
	switch f3 {
	case 0b000:
		return "add/sub/addi/beq/lb/sb/ecall"
	case 0b001:
		return "slli/bne/lh/sh/csrrw"
	case 0b010:
		return "slt/lw/sw/csrrs"
	case 0b011:
		return "sltu/sltui/scrrc"
	case 0b100:
		return "xor/xori/blt/lbu"
	case 0b101:
		return "srl/srli/sra/srai/bge/lhu/csrrwi"
	case 0b110:
		return "or/ori/bltu/csrrsi"
	case 0b111:
		return "and/andi/bgeu/csrrci"
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

	return fmt.Sprintf("?? %x", uint8(r))
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
	imm_u := sign_extend(bits(31, 12), 32)

	// registers
	d := Register(bits(11, 7))
	s1 := Register(bits(19, 15))
	s2 := Register(bits(24, 20))
	_ = s2

	func3 := Func3(bits(14, 12))
	func7 := bits(31, 25)

	switch opcode {
	case opcodeLUI: // U
		return fmt.Sprintf("lui %s, %#x", d, imm_u)
	case opcodeLOAD: // I
		return fmt.Sprintf("load")
	case opcodeSTORE: // S
		return fmt.Sprintf("store")
	case opcodeAUIPC: // U
		return fmt.Sprintf("auipc %s, %#x", d, imm_u)
	case opcodeBRANCH: // B
		addr += imm_b

		sym := ""
		if s.symbols[addr] != "" {
			sym = "<" + s.symbols[addr] + ">"
		}

		return fmt.Sprintf("%s %s, %s, %08x %s",
			func3.Branch(), s1, s2, addr, sym)
	case opcodeJAL: // J
		addr += imm_j

		sym := ""
		if s.symbols[addr] != "" {
			sym = "<" + s.symbols[addr] + ">"
		}

		return fmt.Sprintf("jal %08x %s", addr, sym)
	case opcodeJALR: // I
		return fmt.Sprintf("jalr %s, %#x", s1, imm_i)
	case opcodeIMM: // I
		return fmt.Sprintf("%s %s, %s, %#x", func3.Arith(true, func7 == 0b0100000 && func3 == func3SRAI),
			d, s1, imm_i)
	case opcodeOP: // R
		return fmt.Sprintf("%s %s, %s, %s", func3.Arith(false, func7 == 0b0100000), d, s1, s2)
	case opcodeMISC:
		return fmt.Sprintf("%s", func3.Misc())
	case opcodeSYSTEM: // I
		return fmt.Sprintf("%s", func3.System())
	}

	return fmt.Sprintf("%s %s %#b", opcode, func3, func7)
}
