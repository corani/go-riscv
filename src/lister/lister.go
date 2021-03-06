package lister

import (
	"fmt"

	"github.com/corani/go-riscv/src/riscv"
)

type Printer interface {
	PrintLinef(text string, args ...interface{})
	PrintInstruction(i riscv.Instruction)
}

func NewPrinter() Printer {
	return &visitor{}
}

type visitor struct{}

func (v *visitor) PrintLinef(text string, args ...interface{}) {
	fmt.Printf(text, args...)
}

func (v *visitor) PrintInstruction(i riscv.Instruction) {
	_ = i.Visit(v)
}

func (v *visitor) printInstr(i riscv.Instruction, text string) bool {
	if i.Sym() != "" {
		fmt.Printf("\n%08x <%s>:\n", i.Addr(), i.Sym())
	}

	fmt.Printf("%08x:       %08x        %s\n", i.Addr(), i.Raw(), text)

	return true
}

func (v *visitor) Unimp(i *riscv.Unimp) bool {
	if i.Raw() == 0 {
		return v.printInstr(i, i.Mnemonic())
	}

	return v.printInstr(i, fmt.Sprintf("%-8s # %d",
		i.Mnemonic(), i.Raw()))
}

func (v *visitor) Lui(i *riscv.Lui) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %#x",
		i.Mnemonic(), i.Rd(), i.Imm()))
}

func (v *visitor) Auipc(i *riscv.Auipc) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %#x",
		i.Mnemonic(), i.Rd(), i.Imm()))
}

func (v *visitor) Jal(i *riscv.Jal) bool {
	addr := i.Target()

	// Syntactic Sugar: jal zero, offset == j offset
	if i.Rd() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("j        %08x %s",
			addr, i.NearestSymbol(addr)))
	}

	// Syntactic Sugar: jal x1, offset == jal offset
	if i.Rd() == riscv.Register(1) {
		return v.printInstr(i, fmt.Sprintf("%-8s %08x %s",
			i.Mnemonic(), addr, i.NearestSymbol(addr)))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %08x %s",
		i.Mnemonic(), i.Rd(), addr, i.NearestSymbol(addr)))

}

func (v *visitor) Jalr(i *riscv.Jalr) bool {
	if i.Imm() == 0 {
		if i.Rd() == riscv.Register(0) && i.Rs1() == riscv.Register(1) {
			return v.printInstr(i, "ret")
		}

		if i.Rd() == riscv.Register(0) {
			return v.printInstr(i, fmt.Sprintf("jr       %s",
				i.Rs1()))
		}

		if i.Rd() == riscv.Register(1) {
			return v.printInstr(i, fmt.Sprintf("%-8s %s",
				i.Mnemonic(), i.Rs1()))
		}

		return v.printInstr(i, fmt.Sprintf("%-8s %s, %s",
			i.Mnemonic(), i.Rd(), i.Rs1()))
	}

	if i.Rd() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("jr        %d(%s)",
			i.Imm(), i.Rs1()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %d(%s)",
		i.Mnemonic(), i.Rd(), i.Imm(), i.Rs1()))
}

func (v *visitor) Beq(i *riscv.Beq) bool {
	addr := i.Target()

	if i.Rs2() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("beqz     %s, %08x %s",
			i.Rs1(), addr, i.NearestSymbol(addr)))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %08x %s",
		i.Mnemonic(), i.Rs1(), i.Rs2(), addr, i.NearestSymbol(addr)))
}

func (v *visitor) Bne(i *riscv.Bne) bool {
	addr := i.Target()

	if i.Rs2() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("bnez     %s, %08x %s",
			i.Rs1(), addr, i.NearestSymbol(addr)))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %08x %s",
		i.Mnemonic(), i.Rs1(), i.Rs2(), addr, i.NearestSymbol(addr)))
}

func (v *visitor) Blt(i *riscv.Blt) bool {
	addr := i.Target()

	if i.Rs2() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("bltz     %s, %08x %s",
			i.Rs1(), addr, i.NearestSymbol(addr)))
	}

	if i.Rs1() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("bgtz     %s, %08x %s",
			i.Rs2(), addr, i.NearestSymbol(addr)))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %08x %s",
		i.Mnemonic(), i.Rs1(), i.Rs2(), addr, i.NearestSymbol(addr)))
}

func (v *visitor) Bge(i *riscv.Bge) bool {
	addr := i.Target()

	if i.Rs2() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("bgez     %s, %08x %s",
			i.Rs1(), addr, i.NearestSymbol(addr)))
	}

	if i.Rs1() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("blez     %s, %08x %s",
			i.Rs2(), addr, i.NearestSymbol(addr)))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %08x %s",
		i.Mnemonic(), i.Rs1(), i.Rs2(), addr, i.NearestSymbol(addr)))
}

func (v *visitor) Bltu(i *riscv.Bltu) bool {
	addr := i.Target()

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %08x %s",
		i.Mnemonic(), i.Rs1(), i.Rs2(), addr, i.NearestSymbol(addr)))
}

func (v *visitor) Bgeu(i *riscv.Bgeu) bool {
	addr := i.Target()

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %08x %s",
		i.Mnemonic(), i.Rs1(), i.Rs2(), addr, i.NearestSymbol(addr)))
}

func (v *visitor) Lb(i *riscv.Lb) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %d(%s)",
		i.Mnemonic(), i.Rd(), i.Imm(), i.Rs1()))
}

func (v *visitor) Lh(i *riscv.Lh) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %d(%s)",
		i.Mnemonic(), i.Rd(), i.Imm(), i.Rs1()))
}

func (v *visitor) Lw(i *riscv.Lw) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %d(%s)",
		i.Mnemonic(), i.Rd(), i.Imm(), i.Rs1()))
}

func (v *visitor) Lbu(i *riscv.Lbu) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %d(%s)",
		i.Mnemonic(), i.Rd(), i.Imm(), i.Rs1()))
}

func (v *visitor) Lhu(i *riscv.Lhu) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %d(%s)",
		i.Mnemonic(), i.Rd(), i.Imm(), i.Rs1()))
}

func (v *visitor) Sb(i *riscv.Sb) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %d(%s)",
		i.Mnemonic(), i.Rs2(), i.Imm(), i.Rs1()))
}

func (v *visitor) Sh(i *riscv.Sh) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %d(%s)",
		i.Mnemonic(), i.Rs2(), i.Imm(), i.Rs1()))
}

func (v *visitor) Sw(i *riscv.Sw) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %d(%s)",
		i.Mnemonic(), i.Rs2(), i.Imm(), i.Rs1()))
}

func (v *visitor) Fence(i *riscv.Fence) bool {
	return v.printInstr(i, i.Mnemonic())
}

func (v *visitor) Fencei(i *riscv.Fencei) bool {
	return v.printInstr(i, i.Mnemonic())
}

func (v *visitor) Ebreak(i *riscv.Ebreak) bool {
	return v.printInstr(i, i.Mnemonic())
}

func (v *visitor) Ecall(i *riscv.Ecall) bool {
	return v.printInstr(i, i.Mnemonic())
}

func (v *visitor) Uret(i *riscv.Uret) bool {
	return v.printInstr(i, i.Mnemonic())
}

func (v *visitor) Sret(i *riscv.Sret) bool {
	return v.printInstr(i, i.Mnemonic())
}

func (v *visitor) Mret(i *riscv.Mret) bool {
	return v.printInstr(i, i.Mnemonic())
}

func (v *visitor) Wfi(i *riscv.Wfi) bool {
	return v.printInstr(i, i.Mnemonic())
}

func (v *visitor) Sfence(i *riscv.Sfence) bool {
	if i.Rs2() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("%-8s %s",
			i.Mnemonic(), i.Rs1()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s",
		i.Mnemonic(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Hfence(i *riscv.Hfence) bool {
	if i.Rs2() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("%-8s %s",
			i.Mnemonic(), i.Rs1()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s",
		i.Mnemonic(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Csrrs(i *riscv.Csrrs) bool {
	if i.Rs1() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("csrr     %s, %s",
			i.Rd(), i.Csr()))
	}

	if i.Rd() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("csrs     %s, %s",
			i.Csr(), i.Rs1()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Csr(), i.Rs1()))
}

func (v *visitor) Csrrw(i *riscv.Csrrw) bool {
	if i.Rd() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("csrw     %s, %s",
			i.Csr(), i.Rs1()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Csr(), i.Rs1()))
}

func (v *visitor) Csrrc(i *riscv.Csrrc) bool {
	if i.Rd() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("csrc     %s, %s",
			i.Csr(), i.Rs1()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Csr(), i.Rs1()))
}

func (v *visitor) Csrrsi(i *riscv.Csrrsi) bool {
	if i.Rd() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("csrsi    %s, %d",
			i.Csr(), i.Uimm()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %d",
		i.Mnemonic(), i.Rd(), i.Csr(), i.Uimm()))
}

func (v *visitor) Csrrwi(i *riscv.Csrrwi) bool {
	if i.Rd() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("csrwi    %s, %d",
			i.Csr(), i.Uimm()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %d",
		i.Mnemonic(), i.Rd(), i.Csr(), i.Uimm()))
}

func (v *visitor) Csrrci(i *riscv.Csrrci) bool {
	if i.Rd() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("csrci    %s, %d",
			i.Csr(), i.Uimm()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %d",
		i.Mnemonic(), i.Rd(), i.Csr(), i.Uimm()))
}

func (v *visitor) Sub(i *riscv.Sub) bool {
	if i.Rs1() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("neg      %s, %s",
			i.Rd(), i.Rs2()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Sra(i *riscv.Sra) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Add(i *riscv.Add) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Sll(i *riscv.Sll) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Slt(i *riscv.Slt) bool {
	if i.Rs2() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("sltz     %s, %s",
			i.Rd(), i.Rs1()))
	}

	if i.Rs1() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("sgtz     %s, %s",
			i.Rd(), i.Rs2()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Sltu(i *riscv.Sltu) bool {
	if i.Rs1() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("snez     %s, %s",
			i.Rd(), i.Rs2()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Xor(i *riscv.Xor) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Srl(i *riscv.Srl) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Or(i *riscv.Or) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2()))
}

func (v *visitor) And(i *riscv.And) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %s",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Rs2()))
}

func (v *visitor) Srai(i *riscv.Srai) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %d",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Shamt()))
}

func (v *visitor) Addi(i *riscv.Addi) bool {
	if i.Rd() == riscv.Register(0) && i.Rs1() == riscv.Register(0) && i.Imm() == 0 {
		return v.printInstr(i, "nop")
	}

	if i.Rs1() == riscv.Register(0) {
		return v.printInstr(i, fmt.Sprintf("li       %s, %d",
			i.Rd(), i.Imm()))
	}

	if i.Imm() == 0 {
		return v.printInstr(i, fmt.Sprintf("mv       %s, %s",
			i.Rd(), i.Rs1()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %d",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Imm()))
}

func (v *visitor) Slti(i *riscv.Slti) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %d",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Imm()))
}

func (v *visitor) Sltiu(i *riscv.Sltiu) bool {
	if i.Imm() == 1 {
		return v.printInstr(i, fmt.Sprintf("seqz     %s, %s",
			i.Rd(), i.Rs1()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %d",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Imm()))
}

func (v *visitor) Xori(i *riscv.Xori) bool {
	if i.Imm() == -1 {
		return v.printInstr(i, fmt.Sprintf("not      %s, %s",
			i.Rd(), i.Rs1()))
	}

	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %d",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Imm()))
}

func (v *visitor) Slli(i *riscv.Slli) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %#x",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Shamt()))
}

func (v *visitor) Srli(i *riscv.Srli) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %#x",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Shamt()))
}

func (v *visitor) Ori(i *riscv.Ori) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %d",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Imm()))
}

func (v *visitor) Andi(i *riscv.Andi) bool {
	return v.printInstr(i, fmt.Sprintf("%-8s %s, %s, %d",
		i.Mnemonic(), i.Rd(), i.Rs1(), i.Imm()))
}
