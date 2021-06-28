package main

import (
	"fmt"

	"github.com/corani/go-riscv/src/riscv"
)

type Syscall interface {
	fmt.Stringer

	Execute(*visitor, []uint32)
}

type syscall struct {
	id   uint32
	name string
}

type SyscallWrite struct{ *syscall }
type SyscallExit struct{ *syscall }
type SyscallKill struct{ *syscall }

var syscalls map[uint32]Syscall

func SyscallByID(id uint32) Syscall {
	if len(syscalls) == 0 {
		syscalls = map[uint32]Syscall{
			64:  &SyscallWrite{&syscall{id: 64, name: "write"}},
			93:  &SyscallExit{&syscall{id: 93, name: "exit"}},
			129: &SyscallKill{&syscall{id: 129, name: "kill"}},
		}
	}

	if v, ok := syscalls[id]; ok {
		return v
	}

	return &syscall{id: id, name: fmt.Sprintf("?%d?", id)}
}

func (s *syscall) String() string {
	return s.name
}

func (s *syscall) Execute(v *visitor, args []uint32) {
	v.list.PrintLinef("=> ecall(%d)\n", s.id)
}

func (v *visitor) getEcallArgs() []uint32 {
	names := []string{"a0", "a1", "a2", "a3", "a4", "a5", "a6", "a7"}

	args := make([]uint32, len(names))

	for i, n := range names {
		args[i] = v.registers[riscv.RegisterByName(n)]
	}

	return args
}

func (v *visitor) Ecall(i *riscv.Ecall) bool {
	args := v.getEcallArgs()

	call := SyscallByID(args[7])

	v.profile.recordSyscall(call)

	call.Execute(v, args)

	return true
}

func (s *SyscallExit) Execute(v *visitor, args []uint32) {
	code := args[0]

	v.done = true
	v.exitCode = int(code)

	v.list.PrintLinef("=> exit(%d)\n", code)
}

func (s *SyscallKill) Execute(v *visitor, args []uint32) {
	code := args[0]

	v.done = true
	v.exitCode = int(code)

	v.list.PrintLinef("=> kill(%d)\n", code)
}

func (s *SyscallWrite) Execute(v *visitor, args []uint32) {
	fd := args[0]
	buf := args[1]
	count := args[2]

	section := v.SectionFor(buf)
	if section == nil {
		return
	}

	buf -= section.Base()
	bs := section.GetBytes(buf, count)

	_ = fd

	fmt.Print(string(bs))
}
