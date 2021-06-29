package main

import (
	"fmt"
	"os"

	"github.com/corani/go-riscv/src/riscv"
)

type Syscall interface {
	fmt.Stringer

	Execute(*visitor, []uint32) uint
}

type syscall struct {
	id   uint32
	name string
}

type SyscallRead struct{ *syscall }
type SyscallWrite struct{ *syscall }
type SyscallExit struct{ *syscall }
type SyscallKill struct{ *syscall }

var syscalls map[uint32]Syscall

func SyscallByID(id uint32) Syscall {
	if len(syscalls) == 0 {
		syscalls = map[uint32]Syscall{
			63:  &SyscallRead{&syscall{id: 63, name: "read"}},
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

func (s *syscall) Execute(v *visitor, args []uint32) uint {
	v.list.PrintLinef("=> ecall(%d)\n", s.id)

	return 0
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

	gas := call.Execute(v, args)

	v.gas -= int64(gas)

	return true
}

func (s *SyscallExit) Execute(v *visitor, args []uint32) uint {
	code := args[0]

	v.done = true
	v.exitCode = int(code)

	if v.verbose > 1 {
		v.list.PrintLinef("=> exit(%d)\n", code)
	}

	return 0
}

func (s *SyscallKill) Execute(v *visitor, args []uint32) uint {
	code := args[0]

	v.done = true
	v.exitCode = int(code)

	if v.verbose > 1 {
		v.list.PrintLinef("=> kill(%d)\n", code)
	}

	return 0
}

func (s *SyscallRead) Execute(v *visitor, args []uint32) uint {
	fd := args[0]
	buf := args[1]
	count := args[2]

	if v.verbose > 1 {
		v.list.PrintLinef("=> read(%d, %08x, %d)\n", fd, buf, count)
	}

	section := v.SectionFor(buf)
	if section == nil {
		v.list.PrintLinef("!! section not found")

		return 0
	}

	bs := section.MemAt(buf)

	var file *os.File

	if fd == 0 {
		file = os.Stdin
	}

	if file != nil {
		n, err := file.Read(bs[:count])
		if err != nil {
			v.list.PrintLinef("!! read error: %v\n", err)

			n = -1
		} else {
			v.profile.storedBytes += uint(n)
		}

		v.setReg(riscv.RegisterByName("a0"), int32(n))

		if n < 0 {
			n = -n
		}

		return uint(n)
	}

	return uint(count)
}

func (s *SyscallWrite) Execute(v *visitor, args []uint32) uint {
	fd := args[0]
	buf := args[1]
	count := args[2]

	if v.verbose > 1 {
		v.list.PrintLinef("=> write(%d, %08x, %d)\n", fd, buf, count)
	}

	section := v.SectionFor(buf)
	if section == nil {
		return 0
	}

	bs := section.MemAt(buf)

	var file *os.File

	switch fd {
	case 1:
		file = os.Stdout
	case 2:
		file = os.Stderr
	}

	if file != nil {
		n, err := file.Write(bs[0:count])
		if err != nil {
			n = -1
		} else {
			v.profile.loadedBytes += uint(n)
		}

		v.setReg(riscv.RegisterByName("a0"), int32(n))

		if n < 0 {
			n = -n
		}

		return uint(n)
	}

	return uint(count)
}
