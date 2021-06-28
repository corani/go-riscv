package riscv

import (
	"encoding/binary"
	"fmt"
	"math"
)

type Section interface {
	Name() string
	Base() uint32
	Size() uint32
	MemAt(uint32) []byte
	AddSymbol(uint32, string)
	SymbolAt(uint32) (string, bool)
	SymbolBefore(uint32) string
	InstructionAt(uint32) Instruction
	Reader() SectionReader
}

type SectionReader interface {
	Next() Instruction
}

type sectionReader struct {
	section *section
	pc      uint32
}

type section struct {
	name    string
	base    uint32
	size    uint32
	symbols map[uint32]string
	raw     []byte
}

func NewSection(name string, base, size uint32) Section {
	return &section{
		name:    name,
		base:    base,
		size:    size,
		raw:     make([]byte, size),
		symbols: make(map[uint32]string),
	}
}

func (s *section) Name() string {
	return s.name
}

func (s *section) Base() uint32 {
	return s.base
}

func (s *section) Size() uint32 {
	return s.size
}

func (s *section) MemAt(addr uint32) []byte {
	return s.raw[addr-s.base:]
}

func (s *section) AddSymbol(addr uint32, name string) {
	s.symbols[addr] = name
}

func (s *section) SymbolAt(addr uint32) (string, bool) {
	sym, ok := s.symbols[addr]

	return sym, ok
}

func (s *section) InstructionAt(addr uint32) Instruction {
	sym, _ := s.SymbolAt(addr)
	raw := binary.LittleEndian.Uint32(s.raw[(addr - s.base):])

	return decodeInstruction(s, addr, raw, sym)
}

func (s *section) SymbolBefore(addr uint32) string {
	offset := uint32(math.MaxUint32)
	sym := ""

	for k, v := range s.symbols {
		if k <= addr && (addr-k) < offset {
			offset = addr - k
			sym = v
		}
	}

	if offset == 0 {
		return fmt.Sprintf("<%s>", sym)
	}

	return fmt.Sprintf("<%s+%#x>", sym, offset)
}

func (s *section) Reader() SectionReader {
	return &sectionReader{section: s}
}

func (r *sectionReader) Next() Instruction {
	if int(r.pc) >= len(r.section.raw) {
		return nil
	}

	inst := r.section.InstructionAt(r.pc + r.section.base)

	r.pc += 4

	return inst
}
