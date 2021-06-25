package main

import (
	"fmt"

	"github.com/corani/go-riscv/src/lister"
	"github.com/corani/go-riscv/src/riscv"
)

func NewEmulator(verbose bool, entry uint32) *visitor {
	result := &visitor{
		registers: make(map[riscv.Register]uint32),
		inst:      make(map[uint32]riscv.Instruction),
		pc:        entry,
		list:      lister.NewPrinter(),
		verbose:   verbose,
	}

	result.syscall = &syscall{v: result}

	for i := 0; i < 32; i++ {
		result.registers[riscv.Register(0)] = 0
	}

	return result
}

type visitor struct {
	pc        uint32
	sections  []riscv.Section
	registers map[riscv.Register]uint32
	inst      map[uint32]riscv.Instruction
	list      lister.Printer
	syscall   *syscall
	verbose   bool
	count     uint64
	done      bool
	exitCode  int
}

func (v *visitor) LoadSection(s riscv.Section) {
	if v.verbose {
		v.list.PrintLinef("\n; Disassembly of section %s (base=%08x, size=%d)\n",
			s.Name(), s.Base(), s.Size())
	}

	v.sections = append(v.sections, s)

	r := s.Reader()
	for i := r.Next(); i != nil; i = r.Next() {
		if v.verbose {
			v.list.PrintInstruction(i)
		}

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

func (v *visitor) SectionFor(addr uint32) riscv.Section {
	for _, s := range v.sections {
		if s.Base() <= addr && s.Base()+s.Size() >= addr {
			return s
		}
	}

	return nil
}

func (v *visitor) Step() bool {
	i := v.Current()
	v.count++

	v.list.PrintLinef("===== %04d =====\n", v.count)
	v.list.PrintInstruction(i)

	if i.Visit(v) {
		v.pc += 4
	}

	if v.verbose {
		status := fmt.Sprintf("|\n|\t pc=%#8x", v.pc)

		for r := 1; r < 32; r++ {
			if r%8 == 0 {
				status += "\n|"
			}

			status += fmt.Sprintf("\t%3s=%#08x", riscv.Register(r), v.registers[riscv.Register(r)])
		}

		status += "\n|\n"

		v.list.PrintLinef(status)
	}

	if v.done {
		v.list.PrintLinef("===== done =====\n\n")

		return false
	}

	return true
}

func (v *visitor) Unimp(i *riscv.Unimp) bool {
	return true
}

func (v *visitor) Lui(i *riscv.Lui) bool {
	v.setRegu(i.Rd(), uint32(i.Imm())<<12)

	return true
}

func (v *visitor) Auipc(i *riscv.Auipc) bool {
	v.setRegu(i.Rd(), i.Addr()+uint32(i.Imm())<<12)

	return true
}

func (v *visitor) Jal(i *riscv.Jal) bool {
	target := i.Target()

	// return address
	if i.Rd() != riscv.Register(0) {
		v.setRegu(i.Rd(), v.pc+4)
	}

	return v.jump(target)
}

func (v *visitor) Jalr(i *riscv.Jalr) bool {
	target := i.Target(v.getRegu(i.Rs1()))

	// return address
	if i.Rd() != riscv.Register(0) {
		v.setRegu(i.Rd(), v.pc+4)
	}

	return v.jump(target)
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

func (v *visitor) Load(addr uint32) uint32 {
	// align to 4 bytes
	base := addr & 0xfffffffc

	if data, ok := v.inst[base]; ok {
		return data.Raw()
	} else {
		panic(fmt.Sprintf("load mem out of range: %08x", base))
	}
}

func (v *visitor) Store(addr uint32, val uint32) {
	// align to 4 bytes
	base := addr & 0xfffffffc

	if _, ok := v.inst[base]; ok {
		v.inst[base].SetRaw(val)
	} else {
		panic(fmt.Sprintf("store mem out of range: %08x", base))
	}
}

func (v *visitor) SignExtend(x uint32, l int) uint32 {
	if x>>(l-1) == 1 {
		return -((1 << l) - x)
	}

	return x
}

func (v *visitor) Lb(i *riscv.Lb) bool {
	addr := i.Mem(v.getRegu(i.Rs1()))

	data := v.Load(addr)

	switch addr & 0x3 {
	case 0:
	case 1:
		data >>= 8
	case 2:
		data >>= 16
	case 3:
		data >>= 24
	}

	data &= 0xff
	data = v.SignExtend(data, 8)

	v.setRegu(i.Rd(), data)

	return true
}

func (v *visitor) Lh(i *riscv.Lh) bool {
	addr := i.Mem(v.getRegu(i.Rs1()))

	data := v.Load(addr)

	switch addr & 0x3 {
	case 0:
	case 2:
		data >>= 16
	}

	data &= 0xffff
	data = v.SignExtend(data, 16)

	v.setRegu(i.Rd(), data)

	return true
}

func (v *visitor) Lw(i *riscv.Lw) bool {
	addr := i.Mem(v.getRegu(i.Rs1()))

	v.setRegu(i.Rd(), v.Load(addr))

	return true
}

func (v *visitor) Lbu(i *riscv.Lbu) bool {
	addr := i.Mem(v.getRegu(i.Rs1()))

	data := v.Load(addr)

	switch addr & 0x3 {
	case 0:
	case 1:
		data >>= 8
	case 2:
		data >>= 16
	case 3:
		data >>= 24
	}

	data &= 0xff

	v.setRegu(i.Rd(), data)

	return true
}

func (v *visitor) Lhu(i *riscv.Lhu) bool {
	addr := i.Mem(v.getRegu(i.Rs1()))

	data := v.Load(addr)

	switch addr & 0x3 {
	case 0:
	case 2:
		data >>= 16
	}

	data &= 0xffff

	v.setRegu(i.Rd(), data)

	return true
}

func (v *visitor) Sb(i *riscv.Sb) bool {
	addr := i.Mem(v.getRegu(i.Rs1()))

	old := v.Load(addr)
	data := v.getRegu(i.Rs2()) & 0xff

	switch addr & 0x3 {
	case 0:
		old = (old & 0xffffff00) | data
	case 1:
		old = (old & 0xffff00ff) | (data << 8)
	case 2:
		old = (old & 0xff00ffff) | (data << 16)
	case 3:
		old = (old & 0x00ffffff) | (data << 24)
	}

	v.Store(addr, old)

	return true
}

func (v *visitor) Sh(i *riscv.Sh) bool {
	addr := i.Mem(v.getRegu(i.Rs1()))

	old := v.Load(addr)
	data := v.getRegu(i.Rs2()) & 0xffff

	switch addr & 0x3 {
	case 0:
		old = (old & 0xffff0000) | data
	case 2:
		old = (old & 0x0000ffff) | (data << 16)
	}

	v.Store(addr, old)

	return true
}

func (v *visitor) Sw(i *riscv.Sw) bool {
	addr := i.Mem(v.getRegu(i.Rs1()))

	v.Store(addr, v.getRegu(i.Rs2()))

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
	shamt := v.getRegu(i.Rs2()) & 0b11111

	v.setReg(i.Rd(), v.getReg(i.Rs1())>>shamt)

	return true
}

func (v *visitor) Add(i *riscv.Add) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())+v.getReg(i.Rs2()))

	return true
}

func (v *visitor) Sll(i *riscv.Sll) bool {
	shamt := v.getRegu(i.Rs2()) & 0b11111

	v.setRegu(i.Rd(), v.getRegu(i.Rs1())<<shamt)

	return true
}

func (v *visitor) Slt(i *riscv.Slt) bool {
	if v.getReg(i.Rs1()) < v.getReg(i.Rs2()) {
		v.setReg(i.Rd(), 1)
	} else {
		v.setReg(i.Rd(), 0)
	}

	return true
}

func (v *visitor) Sltu(i *riscv.Sltu) bool {
	if v.getRegu(i.Rs1()) < v.getRegu(i.Rs2()) {
		v.setReg(i.Rd(), 1)
	} else {
		v.setReg(i.Rd(), 0)
	}

	return true
}

func (v *visitor) Xor(i *riscv.Xor) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())^v.getReg(i.Rs2()))

	return true
}

func (v *visitor) Srl(i *riscv.Srl) bool {
	shamt := v.getRegu(i.Rs2()) & 0b11111

	v.setRegu(i.Rd(), v.getRegu(i.Rs1())>>shamt)

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
	v.setReg(i.Rd(), v.getReg(i.Rs1())>>i.Shamt())

	return true
}

func (v *visitor) Addi(i *riscv.Addi) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())+i.Imm())

	return true
}

func (v *visitor) Slti(i *riscv.Slti) bool {
	if v.getReg(i.Rs1()) < i.Imm() {
		v.setReg(i.Rd(), 1)
	} else {
		v.setReg(i.Rd(), 0)
	}

	return true
}

func (v *visitor) Sltiu(i *riscv.Sltiu) bool {
	if v.getRegu(i.Rs1()) < uint32(i.Imm()) {
		v.setReg(i.Rd(), 1)
	} else {
		v.setReg(i.Rd(), 0)
	}

	return true
}

func (v *visitor) Xori(i *riscv.Xori) bool {
	v.setReg(i.Rd(), v.getReg(i.Rs1())^i.Imm())

	return true
}

func (v *visitor) Slli(i *riscv.Slli) bool {
	v.setRegu(i.Rd(), v.getRegu(i.Rs1())<<i.Shamt())

	return true
}

func (v *visitor) Srli(i *riscv.Srli) bool {
	v.setRegu(i.Rd(), v.getRegu(i.Rs1())>>i.Shamt())

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
	v.list.PrintLinef("| => took branch\n")

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
	if r == riscv.Register(0) {
		return 0
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
