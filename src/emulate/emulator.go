package main

import (
	"fmt"

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
			return emulator.exitCode
		}

		i++
	}

	fmt.Println("==== terminated ====")

	return -1
}
