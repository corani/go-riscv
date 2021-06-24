package main

import (
	"fmt"

	"github.com/corani/go-riscv/src/riscv"
)

func listProgram(program riscv.Program) {
	printer := &visitor{}

	printer.header(fmt.Sprintf("\n\n; %v: file format %v-%v-%v\n\n",
		program.Name(), program.Class(), program.Order(), program.Machine()))

	for _, s := range program.Sections() {
		printer.header(fmt.Sprintf("; Disassembly of section %v (%d instructions)\n",
			s.Name(), s.Size()))

		r := s.Reader()
		for inst := r.Next(); inst != nil; inst = r.Next() {
			inst.Visit(printer)
		}
	}
}
