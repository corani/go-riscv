package main

import (
	"github.com/corani/go-riscv/src/riscv"
)

func NewEmulator() *visitor {
	result := &visitor{
		registers: make(map[riscv.Register]uint32),
		inst:      make(map[uint32]riscv.Instruction),
		pc:        0x80000000,
	}

	for i := 0; i < 32; i++ {
		result.registers[riscv.Register(0)] = 0
	}

	return result
}

func (v *visitor) LoadSection(s riscv.Section) {
	r := s.Reader()
	for i := r.Next(); i != nil; i = r.Next() {
		v.inst[i.Addr()] = i
	}
}

func (v *visitor) PC() uint32 {
	return v.pc
}

func (v *visitor) Current() riscv.Instruction {
	if v, ok := v.inst[v.pc]; ok {
		return v
	}

	return nil
}

func (v *visitor) Step() bool {
	i := v.Current()

	if i.Visit(v) {
		v.pc += 4
	}

	return true
}

type visitor struct {
	pc        uint32
	registers map[riscv.Register]uint32
	inst      map[uint32]riscv.Instruction
}

func (v *visitor) Unimp(i *riscv.Unimp) bool {
	return true
}

func (v *visitor) Lui(i *riscv.Lui) bool {
	return true
}

func (v *visitor) Auipc(i *riscv.Auipc) bool {
	return true
}

func (v *visitor) Jal(i *riscv.Jal) bool {
	// return address
	if i.Rd() != riscv.Register(0) {
		v.setRegu(i.Rd(), v.pc+4)
	}

	return v.jump(i.Target())
}

func (v *visitor) Jalr(i *riscv.Jalr) bool {
	// return address
	if i.Rd() != riscv.Register(0) {
		v.setRegu(i.Rd(), v.pc+4)
	}

	return v.jump(i.Target(v.getRegu(i.Rs1())))
}

func (v *visitor) Beq(i *riscv.Beq) bool {
	if v.getReg(i.Rs1()) == v.getReg(i.Rs2()) {
		return v.jump(i.Target())
	}

	return true
}

func (v *visitor) Bne(i *riscv.Bne) bool {
	if v.getReg(i.Rs1()) != v.getReg(i.Rs2()) {
		return v.jump(i.Target())
	}

	return true
}

func (v *visitor) Blt(i *riscv.Blt) bool {
	if v.getReg(i.Rs1()) < v.getReg(i.Rs2()) {
		return v.jump(i.Target())
	}

	return true
}

func (v *visitor) Bge(i *riscv.Bge) bool {
	if v.getReg(i.Rs1()) >= v.getReg(i.Rs2()) {
		return v.jump(i.Target())
	}

	return true
}

func (v *visitor) Bltu(i *riscv.Bltu) bool {
	if v.getRegu(i.Rs1()) < v.getRegu(i.Rs2()) {
		return v.jump(i.Target())
	}

	return true
}

func (v *visitor) Bgeu(i *riscv.Bgeu) bool {
	if v.getRegu(i.Rs1()) >= v.getRegu(i.Rs2()) {
		return v.jump(i.Target())
	}

	return true
}

func (v *visitor) Lb(i *riscv.Lb) bool {
	return true
}

func (v *visitor) Lh(i *riscv.Lh) bool {
	return true
}

func (v *visitor) Lw(i *riscv.Lw) bool {
	return true
}

func (v *visitor) Lbu(i *riscv.Lbu) bool {
	return true
}

func (v *visitor) Lhu(i *riscv.Lhu) bool {
	return true
}

func (v *visitor) Sb(i *riscv.Sb) bool {
	return true
}

func (v *visitor) Sh(i *riscv.Sh) bool {
	return true
}

func (v *visitor) Sw(i *riscv.Sw) bool {
	return true
}

func (v *visitor) Fence(i *riscv.Fence) bool {
	return true
}

func (v *visitor) Fencei(i *riscv.Fencei) bool {
	return true
}

func (v *visitor) Ebreak(i *riscv.Ebreak) bool {
	return true
}

func (v *visitor) Ecall(i *riscv.Ecall) bool {
	return true
}

func (v *visitor) Uret(i *riscv.Uret) bool {
	return true
}

func (v *visitor) Sret(i *riscv.Sret) bool {
	return true
}

func (v *visitor) Mret(i *riscv.Mret) bool {
	return true
}

func (v *visitor) Wfi(i *riscv.Wfi) bool {
	return true
}

func (v *visitor) Sfence(i *riscv.Sfence) bool {
	return true
}

func (v *visitor) Hfence(i *riscv.Hfence) bool {
	return true
}

func (v *visitor) Csrrs(i *riscv.Csrrs) bool {
	return true
}

func (v *visitor) Csrrw(i *riscv.Csrrw) bool {
	return true
}

func (v *visitor) Csrrc(i *riscv.Csrrc) bool {
	return true
}

func (v *visitor) Csrrsi(i *riscv.Csrrsi) bool {
	return true
}

func (v *visitor) Csrrwi(i *riscv.Csrrwi) bool {
	return true
}

func (v *visitor) Csrrci(i *riscv.Csrrci) bool {
	return true
}

func (v *visitor) Sub(i *riscv.Sub) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())-v.getReg(i.Rs2()))

	return true
}

func (v *visitor) Sra(i *riscv.Sra) bool {
	return true
}

func (v *visitor) Add(i *riscv.Add) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())+v.getReg(i.Rs2()))

	return true
}

func (v *visitor) Sll(i *riscv.Sll) bool {
	return true
}

func (v *visitor) Slt(i *riscv.Slt) bool {
	return true
}

func (v *visitor) Sltu(i *riscv.Sltu) bool {
	return true
}

func (v *visitor) Xor(i *riscv.Xor) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())^v.getReg(i.Rs2()))

	return true
}

func (v *visitor) Srl(i *riscv.Srl) bool {
	return true
}

func (v *visitor) Or(i *riscv.Or) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())|v.getReg(i.Rs2()))

	return true
}

func (v *visitor) And(i *riscv.And) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())&v.getReg(i.Rs2()))

	return true
}

func (v *visitor) Srai(i *riscv.Srai) bool {
	return true
}

func (v *visitor) Addi(i *riscv.Addi) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())+i.Imm())

	return true
}

func (v *visitor) Slti(i *riscv.Slti) bool {
	return true
}

func (v *visitor) Sltiu(i *riscv.Sltiu) bool {
	return true
}

func (v *visitor) Xori(i *riscv.Xori) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())^i.Imm())

	return true
}

func (v *visitor) Slli(i *riscv.Slli) bool {
	return true
}

func (v *visitor) Srli(i *riscv.Srli) bool {
	return true
}

func (v *visitor) Ori(i *riscv.Ori) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())|i.Imm())

	return true
}

func (v *visitor) Andi(i *riscv.Andi) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())&i.Imm())

	return true
}

func (v *visitor) jump(addr uint32) bool {
	v.pc = addr

	return false
}

func (v *visitor) setRegu(r riscv.Register, val uint32) {
	if r == riscv.Register(0) {
		return
	}

	v.registers[r] = val
}

func (v *visitor) getRegu(r riscv.Register) uint32 {
	if r == riscv.Register(0) && v.registers[r] != 0 {
		panic("getReg(0) was not zero")
	}

	return v.registers[r]
}

func (v *visitor) setReg(r riscv.Register, val int32) {
	if r == riscv.Register(0) {
		return
	}

	v.registers[r] = uint32(val)
}
func (v *visitor) getReg(r riscv.Register) int32 {
	if r == riscv.Register(0) {
		return 0
	}

	return int32(v.registers[r])
}
