package main

import "fmt"

type instruction struct {
	section  *Section
	addr     uint32
	raw      uint32
	sym      string
	mnemonic string
}

type Unimp struct {
	*instruction
}

type Lui struct {
	*instruction
}

type Auipc struct {
	*instruction
}

type Jal struct {
	*instruction
}

type Jalr struct {
	*instruction
}

type branch struct {
	*instruction
}

type Beq struct {
	*branch
}

type Bne struct {
	*branch
}

type Blt struct {
	*branch
}

type Bge struct {
	*branch
}

type Bltu struct {
	*branch
}

type Bgeu struct {
	*branch
}

type load struct {
	*instruction
}

type Lb struct {
	*load
}

type Lh struct {
	*load
}

type Lw struct {
	*load
}

type Lbu struct {
	*load
}

type Lhu struct {
	*load
}

type store struct {
	*instruction
}

type Sb struct {
	*store
}

type Sh struct {
	*store
}

type Sw struct {
	*store
}

type Fence struct {
	*instruction
}

type system struct {
	*instruction
}

type Ebreak struct {
	*system
}

type Ecall struct {
	*system
}

type Csrrw struct {
	*system
}

type Csrrs struct {
	*system
}

type Csrrc struct {
	*system
}

type Csrrwi struct {
	*system
}

type Csrrsi struct {
	*system
}

type Csrrci struct {
	*system
}

type opReg struct {
	*instruction
}

type Sub struct {
	*opReg
}

type Sra struct {
	*opReg
}

type Add struct {
	*opReg
}

type Slt struct {
	*opReg
}

type Sltu struct {
	*opReg
}

type Xor struct {
	*opReg
}

type Srl struct {
	*opReg
}

type Or struct {
	*opReg
}

type And struct {
	*opReg
}

type opImm struct {
	*instruction
}

type Srai struct {
	*opImm
}

type Addi struct {
	*opImm
}

type Sltui struct {
	*opImm
}

type Xori struct {
	*opImm
}

type Slli struct {
	*opImm
}

type Srli struct {
	*opImm
}

type Ori struct {
	*opImm
}

type Andi struct {
	*opImm
}

func (i *Unimp) Text() string {
	return "unimp"
}

func (i *Lui) Text() string {
	return fmt.Sprintf("lui %s, %#x",
		i.Rd(), i.Imm())
}

func (i *Auipc) Text() string {
	return fmt.Sprintf("auipc %s, %#x",
		i.Rd(), i.Imm())
}

func (i *Jal) Text() string {
	addr := i.target(i.Imm())

	return fmt.Sprintf("jal %08x %s",
		addr, i.nearestSymbol(addr))
}

func (i *Jalr) Text() string {
	return fmt.Sprintf("jalr %s, %#x",
		i.Rs1(), i.Imm())
}

func (i *branch) Text() string {
	addr := i.target(i.Imm())

	return fmt.Sprintf("%s %s, %s, %08x %s",
		i.Mnemonic(), i.Rs1(), i.Rs2(), addr, i.nearestSymbol(addr))
}

func (i *load) Text() string {
	return fmt.Sprintf("%s %s, %s+%#x",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Imm())
}

func (i *store) Text() string {
	return fmt.Sprintf("%s %s+%#x, %s",
		i.Mnemonic(), i.Rs1(), i.Imm(), i.Rs2())
}

func (i *Fence) Text() string {
	return i.Mnemonic()
}

func (i *system) Text() string {
	return i.Mnemonic()
}

func (i *opReg) Text() string {
	return fmt.Sprintf("%s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2())
}

func (i *opImm) Text() string {
	switch i.Func3() {
	case func3SLLI, func3SRLI:
		return fmt.Sprintf("%s %s, %s, %#x",
			i.Mnemonic(), i.Rd(), i.Rs1(), i.shamt())
	default:
		return fmt.Sprintf("%s %s, %s, %#x",
			i.Mnemonic(), i.Rd(), i.Rs1(), i.Imm())
	}
}

func (i *Addi) Text() string {
	// Syntax sugar: addi zero, zero, 0 == nop
	if i.Rd() == Register(0) && i.Rs1() == Register(0) && i.Imm() == 0 {
		return "nop"
	}

	return i.opImm.Text()
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

func (i *instruction) Mnemonic() string {
	return i.mnemonic
}

func (i *instruction) Text() string {
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

func (i *instruction) nearestSymbol(addr uint32) string {
	return i.section.NearestSymbol(addr)
}

func (i *instruction) target(offset uint32) uint32 {
	return uint32(int64(i.Addr()) + int64(offset))
}

func (i *instruction) shamt() uint8 {
	return uint8(i.Imm() & 0b11111)
}
