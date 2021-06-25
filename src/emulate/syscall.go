package main

import (
	"fmt"

	"github.com/corani/go-riscv/src/riscv"
)

type syscall struct {
	v *visitor
}

func getArgRegs() []riscv.Register {
	names := []string{"a0", "a1", "a2", "a3", "a4", "a5", "a6", "a7"}

	regs := make([]riscv.Register, len(names))

	for i, v := range names {
		regs[i] = riscv.RegisterByName(v)
	}

	return regs
}

func (v *visitor) Ecall(i *riscv.Ecall) bool {
	args := getArgRegs()

	id := v.registers[args[7]]

	switch id {
	case 64: // write
		fd := v.registers[args[0]]
		buf := v.registers[args[1]]
		count := v.registers[args[2]]

		v.syscall.Write(fd, buf, count)
	case 93: // exit
		code := v.registers[args[0]]

		v.syscall.Exit(code)
	case 129: // kill
		code := v.registers[args[0]]

		v.syscall.Kill(code)
	default:
		v.list.PrintLinef("=> ecall(%d)\n", id)
	}

	return true
}

func (s *syscall) Exit(code uint32) {
	s.v.done = true
	s.v.exitCode = int(code)

	s.v.list.PrintLinef("=> exit(%d)\n", code)
}

func (s *syscall) Kill(code uint32) {
	s.v.done = true
	s.v.exitCode = int(code)

	s.v.list.PrintLinef("=> kill(%d)\n", code)
}

func (s *syscall) Write(fd, buf, count uint32) {
	section := s.v.SectionFor(buf)
	if section == nil {
		return
	}

	buf -= section.Base()
	bs := section.GetBytes(buf, count)

	fmt.Print(string(bs))
}
