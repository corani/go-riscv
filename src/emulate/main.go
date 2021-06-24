package main

import (
	"flag"

	"github.com/corani/go-riscv/src/elf"
)

func verify(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var name string

	flag.StringVar(&name, "in", "./riscv-tests/isa/rv32-p-simple", "path to input file")
	flag.Parse()

	program, err := elf.Load(name)
	verify(err)

	runProgram(program)
}
