package main

import (
	"github.com/corani/go-riscv/src/riscv"
)

func runProgram(p riscv.Program, verbose bool, gas int64) int {
	emulator := NewEmulator(verbose, p.Entry(), gas)

	for _, s := range p.Sections() {
		emulator.LoadSection(s)
	}

	for emulator.Step() {
	}

	if emulator.done {
		emulator.list.PrintLinef("===== done =====\n")
		emulator.list.PrintLinef("- gas used: %d/%d (left: %d)\n\n",
			gas-emulator.gas, gas, emulator.gas)
	} else {
		emulator.list.PrintLinef("==== terminated ====\n")
		if emulator.gas == 0 {
			emulator.list.PrintLinef("- ran out of gas (%d)\n\n",
				gas)
		} else {
			emulator.list.PrintLinef("- gas used: %d/%d\n\n",
				gas-emulator.gas, gas)
		}

		emulator.exitCode = -1
	}

	emulator.list.PrintLinef(emulator.profile.String())

	return emulator.exitCode
}
