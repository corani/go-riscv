package riscv

type InstructionVisitor interface {
	Unimp(i *Unimp) bool
	Lui(i *Lui) bool
	Auipc(i *Auipc) bool
	Jal(i *Jal) bool
	Jalr(i *Jalr) bool
	Beq(i *Beq) bool
	Bne(i *Bne) bool
	Blt(i *Blt) bool
	Bge(i *Bge) bool
	Bltu(i *Bltu) bool
	Bgeu(i *Bgeu) bool
	Lb(i *Lb) bool
	Lh(i *Lh) bool
	Lw(i *Lw) bool
	Lbu(i *Lbu) bool
	Lhu(i *Lhu) bool
	Sb(i *Sb) bool
	Sh(i *Sh) bool
	Sw(i *Sw) bool
	Fence(i *Fence) bool
	Fencei(i *Fencei) bool
	Ebreak(i *Ebreak) bool
	Ecall(i *Ecall) bool
	Uret(i *Uret) bool
	Sret(i *Sret) bool
	Mret(i *Mret) bool
	Wfi(i *Wfi) bool
	Sfence(i *Sfence) bool
	Hfence(i *Hfence) bool
	Csrrs(i *Csrrs) bool
	Csrrw(i *Csrrw) bool
	Csrrc(i *Csrrc) bool
	Csrrsi(i *Csrrsi) bool
	Csrrwi(i *Csrrwi) bool
	Csrrci(i *Csrrci) bool
	Sub(i *Sub) bool
	Sra(i *Sra) bool
	Add(i *Add) bool
	Sll(i *Sll) bool
	Slt(i *Slt) bool
	Sltu(i *Sltu) bool
	Xor(i *Xor) bool
	Srl(i *Srl) bool
	Or(i *Or) bool
	And(i *And) bool
	Srai(i *Srai) bool
	Addi(i *Addi) bool
	Slti(i *Slti) bool
	Sltiu(i *Sltiu) bool
	Xori(i *Xori) bool
	Slli(i *Slli) bool
	Srli(i *Srli) bool
	Ori(i *Ori) bool
	Andi(i *Andi) bool
}

func (i *Unimp) Visit(v InstructionVisitor) bool {
	return v.Unimp(i)
}
func (i *Lui) Visit(v InstructionVisitor) bool {
	return v.Lui(i)
}
func (i *Auipc) Visit(v InstructionVisitor) bool {
	return v.Auipc(i)
}
func (i *Jal) Visit(v InstructionVisitor) bool {
	return v.Jal(i)
}
func (i *Jalr) Visit(v InstructionVisitor) bool {
	return v.Jalr(i)
}
func (i *Beq) Visit(v InstructionVisitor) bool {
	return v.Beq(i)
}
func (i *Bne) Visit(v InstructionVisitor) bool {
	return v.Bne(i)
}
func (i *Blt) Visit(v InstructionVisitor) bool {
	return v.Blt(i)
}
func (i *Bge) Visit(v InstructionVisitor) bool {
	return v.Bge(i)
}
func (i *Bltu) Visit(v InstructionVisitor) bool {
	return v.Bltu(i)
}
func (i *Bgeu) Visit(v InstructionVisitor) bool {
	return v.Bgeu(i)
}
func (i *Lb) Visit(v InstructionVisitor) bool {
	return v.Lb(i)
}
func (i *Lh) Visit(v InstructionVisitor) bool {
	return v.Lh(i)
}
func (i *Lw) Visit(v InstructionVisitor) bool {
	return v.Lw(i)
}
func (i *Lbu) Visit(v InstructionVisitor) bool {
	return v.Lbu(i)
}
func (i *Lhu) Visit(v InstructionVisitor) bool {
	return v.Lhu(i)
}
func (i *Sb) Visit(v InstructionVisitor) bool {
	return v.Sb(i)
}
func (i *Sh) Visit(v InstructionVisitor) bool {
	return v.Sh(i)
}
func (i *Sw) Visit(v InstructionVisitor) bool {
	return v.Sw(i)
}
func (i *Fence) Visit(v InstructionVisitor) bool {
	return v.Fence(i)
}
func (i *Fencei) Visit(v InstructionVisitor) bool {
	return v.Fencei(i)
}
func (i *Ebreak) Visit(v InstructionVisitor) bool {
	return v.Ebreak(i)
}
func (i *Ecall) Visit(v InstructionVisitor) bool {
	return v.Ecall(i)
}
func (i *Uret) Visit(v InstructionVisitor) bool {
	return v.Uret(i)
}
func (i *Sret) Visit(v InstructionVisitor) bool {
	return v.Sret(i)
}
func (i *Mret) Visit(v InstructionVisitor) bool {
	return v.Mret(i)
}
func (i *Wfi) Visit(v InstructionVisitor) bool {
	return v.Wfi(i)
}
func (i *Sfence) Visit(v InstructionVisitor) bool {
	return v.Sfence(i)
}
func (i *Hfence) Visit(v InstructionVisitor) bool {
	return v.Hfence(i)
}
func (i *Csrrw) Visit(v InstructionVisitor) bool {
	return v.Csrrw(i)
}
func (i *Csrrs) Visit(v InstructionVisitor) bool {
	return v.Csrrs(i)
}
func (i *Csrrc) Visit(v InstructionVisitor) bool {
	return v.Csrrc(i)
}
func (i *Csrrwi) Visit(v InstructionVisitor) bool {
	return v.Csrrwi(i)
}
func (i *Csrrsi) Visit(v InstructionVisitor) bool {
	return v.Csrrsi(i)
}
func (i *Csrrci) Visit(v InstructionVisitor) bool {
	return v.Csrrci(i)
}
func (i *Sub) Visit(v InstructionVisitor) bool {
	return v.Sub(i)
}
func (i *Sra) Visit(v InstructionVisitor) bool {
	return v.Sra(i)
}
func (i *Add) Visit(v InstructionVisitor) bool {
	return v.Add(i)
}
func (i *Sll) Visit(v InstructionVisitor) bool {
	return v.Sll(i)
}
func (i *Slt) Visit(v InstructionVisitor) bool {
	return v.Slt(i)
}
func (i *Sltu) Visit(v InstructionVisitor) bool {
	return v.Sltu(i)
}
func (i *Xor) Visit(v InstructionVisitor) bool {
	return v.Xor(i)
}
func (i *Srl) Visit(v InstructionVisitor) bool {
	return v.Srl(i)
}
func (i *Or) Visit(v InstructionVisitor) bool {
	return v.Or(i)
}
func (i *And) Visit(v InstructionVisitor) bool {
	return v.And(i)
}
func (i *Srai) Visit(v InstructionVisitor) bool {
	return v.Srai(i)
}
func (i *Addi) Visit(v InstructionVisitor) bool {
	return v.Addi(i)
}
func (i *Slti) Visit(v InstructionVisitor) bool {
	return v.Slti(i)
}
func (i *Sltiu) Visit(v InstructionVisitor) bool {
	return v.Sltiu(i)
}
func (i *Xori) Visit(v InstructionVisitor) bool {
	return v.Xori(i)
}
func (i *Slli) Visit(v InstructionVisitor) bool {
	return v.Slli(i)
}
func (i *Srli) Visit(v InstructionVisitor) bool {
	return v.Srli(i)
}
func (i *Ori) Visit(v InstructionVisitor) bool {
	return v.Ori(i)
}
func (i *Andi) Visit(v InstructionVisitor) bool {
	return v.Andi(i)
}
