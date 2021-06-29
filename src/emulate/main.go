package main

import (
	"flag"
	"os"

	"github.com/corani/go-riscv/src/elf"
)

func verify(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		name    string
		verbose int
		gas     int64
	)

	flag.StringVar(&name, "in", "./riscv-tests/isa/rv32-p-simple", "path to input file")
	flag.IntVar(&verbose, "v", 0, "verbose logging")
	flag.Int64Var(&gas, "gas", 500, "maximum amount of gas to use")
	flag.Parse()

	program, err := elf.Load(name)
	verify(err)

	os.Exit(runProgram(program, verbose, gas))
}
