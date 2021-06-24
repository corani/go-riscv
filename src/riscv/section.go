package riscv

import (
	"fmt"
	"math"
)

type Section interface {
	Name() string
	Base() uint32
	Size() uint32
	Data() *[]uint32
	AddSymbol(uint32, string)
	SymbolAtAddr(uint32) (string, bool)
	NearestSymbol(uint32) string
	SymbolForIndex(uint32) (string, bool)
	AddrForIndex(uint32) uint32
	Reader() SectionReader
}

type SectionReader interface {
	Next() Instruction
}

type sectionReader struct {
	section *section
	index   uint32
}

type section struct {
	name    string
	base    uint32
	size    uint32
	symbols map[uint32]string
	data    []uint32
}

func NewSection(name string, base, size uint32) Section {
	return &section{
		name:    name,
		base:    base,
		size:    size,
		data:    make([]uint32, size),
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

func (s *section) Data() *[]uint32 {
	return &s.data
}

func (s *section) AddSymbol(addr uint32, name string) {
	s.symbols[addr] = name
}

func (s *section) SymbolAtAddr(addr uint32) (string, bool) {
	sym, ok := s.symbols[addr]
	return sym, ok
}

func (s *section) NearestSymbol(addr uint32) string {
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

func (s *section) SymbolForIndex(i uint32) (string, bool) {
	addr := s.AddrForIndex(i)

	sym, ok := s.symbols[addr]
	return sym, ok
}

func (s *section) AddrForIndex(i uint32) uint32 {
	return s.base + i*4
}

func (s *section) Reader() SectionReader {
	return &sectionReader{
		section: s,
		index:   0,
	}
}

func (r *sectionReader) Next() Instruction {
	if int(r.index) >= len(r.section.data) {
		return nil
	}

	sym, _ := r.section.SymbolForIndex(r.index)
	addr := r.section.AddrForIndex(r.index)
	raw := r.section.data[r.index]

	inst := decodeInstruction(r.section, addr, raw, sym)

	r.index++

	return inst
}
