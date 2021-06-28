package main

import (
	"fmt"
	"sort"

	"github.com/corani/go-riscv/src/riscv"
)

type profile struct {
	inst        map[string]uint
	ecall       map[string]uint
	mem         map[string]uint
	loadedBytes uint
	storedBytes uint
}

func newProfile() *profile {
	return &profile{
		inst:  make(map[string]uint),
		ecall: make(map[string]uint),
		mem:   make(map[string]uint),
	}
}

func (p *profile) recordInstruction(i riscv.Instruction) {
	p.inst[i.Mnemonic()]++

	switch i.(type) {
	case *riscv.Lb, *riscv.Lbu:
		p.loadedBytes += 1
	case *riscv.Lh, *riscv.Lhu:
		p.loadedBytes += 2
	case *riscv.Lw:
		p.loadedBytes += 4
	case *riscv.Sb:
		p.storedBytes += 1
	case *riscv.Sh:
		p.storedBytes += 2
	case *riscv.Sw:
		p.storedBytes += 4
	}
}

func (p *profile) recordSyscall(s Syscall) {
	p.ecall[s.String()]++
}

func (p *profile) String() string {
	type pair struct {
		mnemonic string
		count    uint
	}

	sortProfile := func(prof map[string]uint) []pair {
		var pairs []pair

		for k, c := range prof {
			pairs = append(pairs, pair{k, c})
		}

		sort.Slice(pairs, func(i, j int) bool {
			switch {
			case pairs[i].count > pairs[j].count:
				return true
			case pairs[i].count < pairs[j].count:
				return false
			default:
				return pairs[i].mnemonic < pairs[j].mnemonic
			}
		})

		return pairs
	}

	res := "===== profile =====\n"

	if len(p.inst) > 0 {
		res += "instructions:\n"

		for _, pair := range sortProfile(p.inst) {
			res += fmt.Sprintf(" - %-08s: %4d\n", pair.mnemonic, pair.count)
		}
	}

	if len(p.ecall) > 0 {
		res += "\nsyscalls:\n"

		for _, pair := range sortProfile(p.ecall) {
			res += fmt.Sprintf(" - %-08s: %4d\n", pair.mnemonic, pair.count)
		}
	}

	res += "\nmemory:\n"
	res += fmt.Sprintf("- loaded: %d bytes\n", p.loadedBytes)
	res += fmt.Sprintf("- stored: %d bytes\n", p.storedBytes)

	return res
}
