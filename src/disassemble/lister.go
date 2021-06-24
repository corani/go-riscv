package main

import (
	"fmt"

	"github.com/corani/go-riscv/src/riscv"
)

func listProgram(program riscv.Program) {
	fmt.Printf("\n\n; %v: file format %v-%v-%v\n\n",
		program.Name(), program.Class(), program.Order(), program.Machine())

	for _, s := range program.Sections() {
		fmt.Printf("; Disassembly of section %v (%d instructions)\n", s.Name(), s.Size())

		r := s.Reader()
		for inst := r.Next(); inst != nil; inst = r.Next() {
			if inst.Sym() != "" {
				fmt.Printf("\n%08x <%s>:\n", inst.Addr(), inst.Sym())
			}

			fmt.Printf("%08x:       %08x        %s\n", inst.Addr(), inst.Raw(), inst.Text())
		}
	}
}
