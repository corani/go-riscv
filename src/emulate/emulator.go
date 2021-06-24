package main

import (
	"github.com/corani/go-riscv/src/riscv"
)

func runProgram(p riscv.Program) {
	emulator := NewEmulator(true)

	for _, s := range p.Sections() {
		emulator.LoadSection(s)
	}

	i := 0

	for i < 500 {
		if !emulator.Step() {
			break
		}

		i++
	}
}
