package main

import (
	"fmt"

	"github.com/corani/go-riscv/src/riscv"
)

func runProgram(p riscv.Program) {
	emulator := NewEmulator()

	for _, s := range p.Sections() {
		emulator.LoadSection(s)
	}

	i := 0

	for i < 10 {
		fmt.Printf("%8x    %s\n", emulator.Current().Addr(), emulator.Current().Mnemonic())

		if !emulator.Step() {
			break
		}

		i++
	}
}
