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

	flag.StringVar(&name, "in", "./riscv-tests/isa/rv32ui-p-add", "path to input file")
	flag.Parse()

	program, err := elf.Load(name)
	verify(err)

	listProgram(program)
}
