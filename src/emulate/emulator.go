package main

import (
	"github.com/corani/go-riscv/src/riscv"
)

func runProgram(p riscv.Program, verbose bool, iter int) int {
	emulator := NewEmulator(verbose, p.Entry())

	for _, s := range p.Sections() {
		emulator.LoadSection(s)
	}

	i := 0

	for i < iter {
		if !emulator.Step() {
			break
		}

		i++
	}

	if i < iter {
		emulator.list.PrintLinef("===== done =====\n\n")
	} else {
		emulator.list.PrintLinef("==== terminated ====\n\n")

		emulator.exitCode = -1
	}

	emulator.printProfile()

	return emulator.exitCode
}
