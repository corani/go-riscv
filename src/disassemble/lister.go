package main

import "fmt"

func listProgram(program *Program) {
	fmt.Printf("; %v: file format %v-%v-%v\n",
		program.name, program.class, program.order, program.machine)

	for _, s := range program.sections {
		fmt.Printf("; Section %v (%v instr, %v syms)\n", s.name, len(s.data), len(s.symbols))

		r := s.Reader()
		for inst := r.Next(); inst != nil; inst = r.Next() {
			if inst.Sym() != "" {
				fmt.Printf("\n%08x <%s>:\n", inst.Addr(), inst.Sym())
			}

			fmt.Printf("%08x:       %08x        %s\n", inst.Addr(), inst.Raw(), inst.Text())
		}
	}
}
