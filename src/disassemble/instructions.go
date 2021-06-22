package main

import "fmt"

type instruction struct {
	section  *Section
	addr     uint32
	raw      uint32
	sym      string
	mnemonic string
}

type Unimp struct{ *instruction }

type Lui struct{ *instruction }

type Auipc struct{ *instruction }

type Jal struct{ *instruction }

type Jalr struct{ *instruction }

type Branch struct{ *instruction }

type Beq struct{ *Branch }

type Bne struct{ *Branch }

type Blt struct{ *Branch }

type Bge struct{ *Branch }

type Bltu struct{ *Branch }

type Bgeu struct{ *Branch }

type Load struct{ *instruction }

type Lb struct{ *Load }

type Lh struct{ *Load }

type Lw struct{ *Load }

type Lbu struct{ *Load }

type Lhu struct{ *Load }

type Store struct{ *instruction }

type Sb struct{ *Store }

type Sh struct{ *Store }

type Sw struct{ *Store }

type Fence struct{ *instruction }

type System struct{ *instruction }

type Ebreak struct{ *System }

type Ecall struct{ *System }

type Csrrw struct{ *System }

type Csrrs struct{ *System }

type Csrrc struct{ *System }

type Csrrwi struct{ *System }

type Csrrsi struct{ *System }

type Csrrci struct{ *System }

type OpReg struct{ *instruction }

type Sub struct{ *OpReg }

type Sra struct{ *OpReg }

type Add struct{ *OpReg }

type Slt struct{ *OpReg }

type Sltu struct{ *OpReg }

type Xor struct{ *OpReg }

type Srl struct{ *OpReg }

type Or struct{ *OpReg }

type And struct{ *OpReg }

type OpImm struct{ *instruction }

type Srai struct{ *OpImm }

type Addi struct{ *OpImm }

type Sltiu struct{ *OpImm }

type Xori struct{ *OpImm }

type Slli struct{ *OpImm }

type Srli struct{ *OpImm }

type Ori struct{ *OpImm }

type Andi struct{ *OpImm }

func (i *Unimp) Text() string {
	if i.raw == 0 {
		return i.Mnemonic()
	}

	return fmt.Sprintf("%s %d",
		i.Mnemonic(), i.Raw())
}

func (i *Lui) Text() string {
	return fmt.Sprintf("%-4s %s, %#x",
		i.Mnemonic(), i.Rd(), i.Imm())
}

func (i *Auipc) Text() string {
	return fmt.Sprintf("%-4s %s, %#x",
		i.Mnemonic(), i.Rd(), i.Imm())
}

func (i *Jal) Text() string {
	addr := i.target(i.Imm())

	// Syntactic Sugar: jal zero, offset == j offset
	if i.Rd() == Register(0) {
		return fmt.Sprintf("j    %08x %s",
			addr, i.nearestSymbol(addr))
	}

	// Syntactic Sugar: jal x1, offset == jal offset
	if i.Rd() == Register(1) {
		return fmt.Sprintf("%-4s %08x %s",
			i.Mnemonic(), addr, i.nearestSymbol(addr))
	}

	return fmt.Sprintf("%-4s %s, %08x %s",
		i.Mnemonic(), i.Rd(), addr, i.nearestSymbol(addr))
}

func (i *Jalr) Text() string {
	if i.Imm() == 0 {
		if i.Rd() == Register(0) && i.Rs1() == Register(1) {
			return "ret"
		}

		if i.Rd() == Register(0) {
			return fmt.Sprintf("jr   %s", i.Rs1())
		}

		if i.Rd() == Register(1) {
			return fmt.Sprintf("%-4s %s", i.Mnemonic(), i.Rs1())
		}
	}

	return fmt.Sprintf("%-4s %s, %d(%s)",
		i.Mnemonic(), i.Rd(), i.Imm(), i.Rs1())
}

func (i *Branch) Text() string {
	addr := i.target(i.Imm())

	return fmt.Sprintf("%-4s %s, %s, %08x %s",
		i.Mnemonic(), i.Rs1(), i.Rs2(), addr, i.nearestSymbol(addr))
}

func (i *Beq) Text() string {
	if i.Rs2() == Register(0) {
		addr := i.target(i.Imm())

		return fmt.Sprintf("beqz %s, %08x %s",
			i.Rs1(), addr, i.nearestSymbol(addr))
	}

	return i.Branch.Text()
}

func (i *Bne) Text() string {
	if i.Rs2() == Register(0) {
		addr := i.target(i.Imm())

		return fmt.Sprintf("bnez %s, %08x %s",
			i.Rs1(), addr, i.nearestSymbol(addr))
	}

	return i.Branch.Text()
}

func (i *Blt) Text() string {
	if i.Rs2() == Register(0) {
		addr := i.target(i.Imm())

		return fmt.Sprintf("bltz %s, %08x %s",
			i.Rs1(), addr, i.nearestSymbol(addr))
	}

	if i.Rs1() == Register(0) {
		addr := i.target(i.Imm())

		return fmt.Sprintf("bgtz %s, %08x %s",
			i.Rs2(), addr, i.nearestSymbol(addr))
	}

	return i.Branch.Text()
}

func (i *Bge) Text() string {
	if i.Rs2() == Register(0) {
		addr := i.target(i.Imm())

		return fmt.Sprintf("bgez %s, %08x %s",
			i.Rs1(), addr, i.nearestSymbol(addr))
	}

	if i.Rs1() == Register(0) {
		addr := i.target(i.Imm())

		return fmt.Sprintf("blez %s, %08x %s",
			i.Rs2(), addr, i.nearestSymbol(addr))
	}

	return i.Branch.Text()
}

func (i *Load) Text() string {
	return fmt.Sprintf("%-4s %s, %s+%#x",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Imm())
}

func (i *Store) Text() string {
	return fmt.Sprintf("%-4s %s, %d(%s)",
		i.Mnemonic(), i.Rs2(), i.Imm(), i.Rs1())
}

func (i *Fence) Text() string {
	return i.Mnemonic()
}

func (i *System) Text() string {
	return i.Mnemonic()
}

func (i *OpReg) Text() string {
	return fmt.Sprintf("%-4s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2())
}

func (i *Sub) Text() string {
	if i.Rs1() == Register(0) {
		return fmt.Sprintf("neg  %s, %s", i.Rd(), i.Rs2())
	}

	return i.OpReg.Text()
}

func (i *Slt) Text() string {
	if i.Rs1() == Register(0) {
		return fmt.Sprintf("sgtz %s, %s", i.Rd(), i.Rs2())
	}

	if i.Rs2() == Register(0) {
		return fmt.Sprintf("sltz %s, %s", i.Rd(), i.Rs1())
	}

	return i.OpReg.Text()
}

func (i *Sltu) Text() string {
	if i.Rs1() == Register(0) {
		return fmt.Sprintf("snez %s, %s", i.Rd(), i.Rs2())
	}

	return i.OpReg.Text()
}

func (i *OpImm) Text() string {
	switch i.Func3() {
	case func3SLLI, func3SRLI:
		return fmt.Sprintf("%-4s %s, %s, %#x",
			i.Mnemonic(), i.Rd(), i.Rs1(), i.shamt())
	default:
		return fmt.Sprintf("%-4s %s, %s, %d",
			i.Mnemonic(), i.Rd(), i.Rs1(), i.Imm())
	}
}

func (i *Xori) Text() string {
	if i.Imm() == -1 {
		return fmt.Sprintf("not  %s, %s", i.Rd(), i.Rs1())
	}

	return i.OpImm.Text()
}

func (i *Addi) Text() string {
	if i.Rd() == Register(0) && i.Rs1() == Register(0) && i.Imm() == 0 {
		return "nop"
	}

	if i.Rs1() == Register(0) {
		return fmt.Sprintf("li   %s, %d", i.Rd(), i.Imm())
	}

	if i.Imm() == 0 {
		return fmt.Sprintf("mv   %s, %s", i.Rd(), i.Rs1())
	}

	return i.OpImm.Text()
}

func (i *Sltiu) Text() string {
	if i.Imm() == 1 {
		return fmt.Sprintf("seqz %s, %s", i.Rd(), i.Rs1())
	}

	return i.OpImm.Text()
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

func (i *instruction) Opcode() Opcode {
	return Opcode(i.bits(6, 0))
}

func (i *instruction) Imm() int32 {
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

func (i *instruction) target(offset int32) uint32 {
	return uint32(int64(i.Addr()) + int64(offset))
}

func (i *instruction) shamt() uint8 {
	return uint8(i.Imm() & 0b11111)
}
