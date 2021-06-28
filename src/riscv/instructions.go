package riscv

type Instruction interface {
	Addr() uint32
	Raw() uint32
	Sym() string
	Imm() int32
	Rd() Register
	Rs1() Register
	Rs2() Register
	NearestSymbol(addr uint32) string
	Mnemonic() string
	Visit(InstructionVisitor) bool
	SetRaw(uint32)
}

type instruction struct {
	section  Section
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

type Misc struct{ *instruction }
type Fence struct{ *Misc }
type Fencei struct{ *Misc }

type System struct{ *instruction }
type Ebreak struct{ *System }
type Ecall struct{ *System }
type Uret struct{ *System }
type Sret struct{ *System }
type Mret struct{ *System }
type Wfi struct{ *System }
type Sfence struct{ *System }
type Hfence struct{ *System }
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
type Sll struct{ *OpReg }
type Slt struct{ *OpReg }
type Sltu struct{ *OpReg }
type Xor struct{ *OpReg }
type Srl struct{ *OpReg }
type Or struct{ *OpReg }
type And struct{ *OpReg }

type OpImm struct{ *instruction }
type Srai struct{ *OpImm }
type Addi struct{ *OpImm }
type Slti struct{ *OpImm }
type Sltiu struct{ *OpImm }
type Xori struct{ *OpImm }
type Slli struct{ *OpImm }
type Srli struct{ *OpImm }
type Ori struct{ *OpImm }
type Andi struct{ *OpImm }

func (i *Jal) Target() uint32 {
	return uint32(int64(i.Addr()) + int64(i.Imm()))
}

func (i *Jalr) Target(base uint32) uint32 {
	return uint32(int64(base) + int64(i.Imm()))
}

func (i *Branch) Target() uint32 {
	return uint32(int64(i.Addr()) + int64(i.Imm()))
}

func (i *Load) Mem(base uint32) uint32 {
	return uint32(int64(base) + int64(i.Imm()))
}

func (i *Store) Mem(base uint32) uint32 {
	return uint32(int64(base) + int64(i.Imm()))
}

func (i *OpImm) Shamt() uint32 {
	return uint32(i.Imm()) & 0b11111
}

func (i *System) Csr() string {
	return CsrName(uint32(i.Imm()))
}

func (i *System) Uimm() uint8 {
	return uint8(i.bits(19, 15))
}

func (i *instruction) Section() Section {
	return i.section
}

func (i *instruction) Addr() uint32 {
	return i.addr
}

func (i *instruction) Raw() uint32 {
	return i.raw
}

func (i *instruction) SetRaw(v uint32) {
	i.raw = v
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

func (i *instruction) NearestSymbol(addr uint32) string {
	return i.section.SymbolBefore(addr)
}
