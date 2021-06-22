package main

import "fmt"

type instruction struct {
	section *Section
	addr    uint32
	raw     uint32
	sym     string
}

func (i *instruction) Section() *Section {
	return i.section
}

func (i *instruction) Addr() uint32 {
	return i.addr
}

func (i *instruction) Raw() uint32 {
	return i.raw
}

func (i *instruction) Sym() string {
	return i.sym
}

func (i *instruction) nearestSymbol(addr uint32) string {
	return i.section.NearestSymbol(addr)
}

func (i *instruction) target(offset uint32) uint32 {
	return uint32(int64(i.Addr()) + int64(offset))
}

func (i *instruction) decodeBranch() string {
	return i.Func3().Branch()
}

func (i *instruction) decodeArith(imm bool) string {
	return i.Func3().Arith(imm, i.Func7() == 0b0100000 && i.Func3() == func3SRAI)
}

func (i *instruction) decodeLoad() string {
	return i.Func3().Load()
}

func (i *instruction) decodeStore() string {
	return i.Func3().Store()
}

func (i *instruction) decodeSystem() string {
	return i.Func3().System(i.Imm())
}

func (i *instruction) decodeMisc() string {
	return i.Func3().Misc()
}

func (i *instruction) shamt() uint8 {
	return uint8(i.Imm() & 0b11111)
}

func (i *instruction) Text() string {
	if i.raw == 0 {
		return "unimp"
	} else if i.raw == 0x13 {
		// TODO: should be handled as part of the sprinkling of syntactic sugar
		// addi x0,x0,0 == nop
		return "nop"
	}

	switch i.Opcode() {
	case opcodeLUI: // U-type
		return fmt.Sprintf("lui %s, %#x", i.Rd(), i.Imm())
	case opcodeAUIPC: // U-type
		return fmt.Sprintf("auipc %s, %#x", i.Rd(), i.Imm())
	case opcodeJAL: // J-type
		addr := i.target(i.Imm())

		return fmt.Sprintf("jal %08x %s", addr, i.nearestSymbol(addr))
	case opcodeJALR: // I-type
		return fmt.Sprintf("jalr %s, %#x", i.Rs1(), i.Imm())
	case opcodeBRANCH: // B-type
		addr := i.target(i.Imm())

		return fmt.Sprintf("%s %s, %s, %08x %s",
			i.decodeBranch(), i.Rs1(), i.Rs2(), addr, i.nearestSymbol(addr))
	case opcodeLOAD: // I-type
		return fmt.Sprintf("%s %s, %s+%#x", i.decodeLoad(), i.Rd(), i.Rs1(), i.Imm())
	case opcodeSTORE: // S-type
		return fmt.Sprintf("%s %s+%#x, %s", i.decodeStore(), i.Rs1(), i.Imm(), i.Rs2())
	case opcodeOP_IMM: // I-type
		switch i.Func3() {
		case func3SLLI, func3SRLI:
			return fmt.Sprintf("%s %s, %s, %#x", i.decodeArith(true), i.Rd(), i.Rs1(), i.shamt())
		default:
			return fmt.Sprintf("%s %s, %s, %#x", i.decodeArith(true), i.Rd(), i.Rs1(), i.Imm())
		}
	case opcodeOP: // R-type
		return fmt.Sprintf("%s %s, %s, %s",
			i.decodeArith(false), i.Rd(), i.Rs1(), i.Rs2())
	case opcodeSYSTEM: // I-type
		return fmt.Sprintf("%s", i.decodeSystem())
	case opcodeMISC_MEM:
		return fmt.Sprintf("%s", i.decodeMisc())
	}

	panic(fmt.Sprintf("unknown opcode: %v", i.Opcode()))
}

func (i *instruction) Opcode() Opcode {
	return Opcode(i.bits(6, 0))
}

func (i *instruction) Imm() uint32 {
	opcode := i.Opcode()

	return opcode.decodeImm(i.raw)
}

func (i *instruction) Rd() Register {
	return Register(i.bits(11, 7))
}

func (i *instruction) Rs1() Register {
	return Register(i.bits(19, 15))
}

func (i *instruction) Rs2() Register {
	return Register(i.bits(24, 20))
}

func (i *instruction) Func3() Func3 {
	return Func3(i.bits(14, 12))
}

func (i *instruction) Func7() uint8 {
	return uint8(i.bits(31, 25))
}

func (i *instruction) bits(s, e int) uint32 {
	return (i.raw >> e) & ((1 << (s - e + 1)) - 1)
}
