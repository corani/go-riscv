package elf

import (
	"debug/elf"
	"encoding/binary"
	"fmt"

	"github.com/corani/go-riscv/src/riscv"
)

func Load(name string) (riscv.Program, error) {
	f, err := elf.Open(name)
	if err != nil {
		return nil, err
	}

	if f.Machine != elf.EM_RISCV {
		return nil, fmt.Errorf("machine is not riscv: %v", f.Machine)
	}

	if f.Class != elf.ELFCLASS32 {
		return nil, fmt.Errorf("class is not 32-bit: %v", f.Class)
	}

	if f.Data != elf.ELFDATA2LSB {
		return nil, fmt.Errorf("data is not little-endian: %v", f.Data)
	}

	program := riscv.NewProgram(name, "riscv", "elf32", "little", uint32(f.Entry))

	syms, err := f.Symbols()
	if err != nil {
		return nil, err
	}

	symbols := map[uint32]string{}

	for _, s := range syms {
		symbols[uint32(s.Value)] = s.Name
	}

	for _, s := range f.Sections {
		if s.Addr == 0 {
			continue
		}

		// no program code
		if s.Type&elf.SHT_PROGBITS == 0 {
			continue
		}

		// not executable
		if s.Flags&elf.SHF_EXECINSTR == 0 {
			// data
		}

		section := riscv.NewSection(s.Name, uint32(s.Addr), uint32(s.Size/4))

		for k, v := range symbols {
			if uint64(k) >= s.Addr && uint64(k) <= s.Addr+s.Size {
				section.AddSymbol(k, v)
			}
		}

		if err := binary.Read(s.Open(), binary.LittleEndian, section.Data()); err != nil {
			return nil, err
		}

		program.AddSection(section)
	}

	return program, nil
}
