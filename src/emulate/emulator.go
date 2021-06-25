package main

import (
	"fmt"

	"github.com/corani/go-riscv/src/riscv"
)

func runProgram(p riscv.Program, iter int) int {
	emulator := NewEmulator(false)

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
