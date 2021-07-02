package main

import (
	"fmt"
	"sort"

	"github.com/corani/go-riscv/src/riscv"
)

type profile struct {
	inst        map[string]uint
	mem         map[string]uint
	loadedBytes uint
	storedBytes uint
}

func newProfile() *profile {
	return &profile{
		inst: make(map[string]uint),
		mem:  make(map[string]uint),
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

	separator := func(i int) string {
		if i == 0 {
			return "\n| "
		} else if i%4 == 0 {
			return " |\n| "
		} else {
			return " | "
		}
	}

	res := "===== profile =====\n"

	if len(p.inst) > 0 {
		res += "instructions:"

		for i, pair := range sortProfile(p.inst) {
			res += separator(i)
			res += fmt.Sprintf("%-08s: %4d", pair.mnemonic, pair.count)
		}

		res += " |\n"
	}

	res += "\nmemory:\n"
	res += fmt.Sprintf("- loaded: %d bytes\n", p.loadedBytes)
	res += fmt.Sprintf("- stored: %d bytes\n", p.storedBytes)

	return res
}
