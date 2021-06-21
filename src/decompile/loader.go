package main

import (
	"debug/elf"
	"encoding/binary"
	"fmt"
)

func loadElf(name string) (*Program, error) {
	f, err := elf.Open(name)
	if err != nil {
		return nil, err
	}

	program := Program{
		name: name,
	}

	if f.Machine != elf.EM_RISCV {
		return nil, fmt.Errorf("machine is not riscv: %v", f.Machine)
	}

	program.machine = "riscv"

	if f.Class != elf.ELFCLASS32 {
		return nil, fmt.Errorf("class is not 32-bit: %v", f.Class)
	}

	program.class = "elf32"

	if f.Data != elf.ELFDATA2LSB {
		return nil, fmt.Errorf("data is not little-endian: %v", f.Data)
	}

	program.order = "little"

	syms, err := f.Symbols()
	if err != nil {
		return nil, err
	}

	symbols := map[uint32]string{}

	for _, s := range syms {
		symbols[uint32(s.Value)] = s.Name
	}

	for _, s := range f.Sections {
		// no program code
		if s.Type&elf.SHT_PROGBITS == 0 {
			continue
		}

		// not executable
		if s.Flags&elf.SHF_EXECINSTR == 0 {
			continue
		}

		section := Section{
			name:    s.Name,
			base:    uint32(s.Addr),
			data:    make([]uint32, s.Size/4),
			symbols: make(map[uint32]string),
		}

		for k, v := range symbols {
			if uint64(k) >= s.Addr && uint64(k) <= s.Addr+s.Size {
				section.symbols[k] = v
			}
		}

		if err := binary.Read(s.Open(), binary.LittleEndian, &section.data); err != nil {
			return nil, err
		}

		program.sections = append(program.sections, section)
	}

	return &program, nil
}
