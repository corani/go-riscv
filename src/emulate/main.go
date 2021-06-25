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
		verbose bool
		iter    int
	)

	flag.StringVar(&name, "in", "./riscv-tests/isa/rv32-p-simple", "path to input file")
	flag.BoolVar(&verbose, "v", false, "verbose logging")
	flag.IntVar(&iter, "iter", 500, "maximum number of instructions to run")
	flag.Parse()

	program, err := elf.Load(name)
	verify(err)

	os.Exit(runProgram(program, verbose, iter))
}
