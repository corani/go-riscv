package main

import (
	"fmt"
	"math"
)

type Program struct {
	name     string
	machine  string
	class    string
	order    string
	sections []Section
}

type Section struct {
	name    string
	base    uint32
	symbols map[uint32]string
	data    []uint32
}

type SectionReader struct {
	section *Section
	index   uint32
}

type Instruction interface {
	Addr() uint32
	Raw() uint32
	Sym() string
	Text() string
}

func (s Section) SymbolAtAddr(addr uint32) (string, bool) {
	sym, ok := s.symbols[addr]
	return sym, ok
}

func (s Section) NearestSymbol(addr uint32) string {
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

func (s Section) SymbolForIndex(i uint32) (string, bool) {
	addr := s.AddrForIndex(i)

	sym, ok := s.symbols[addr]
	return sym, ok
}

func (s Section) AddrForIndex(i uint32) uint32 {
	return s.base + i*4
}

func (s Section) Reader() *SectionReader {
	return &SectionReader{
		section: &s,
		index:   0,
	}
}

func (r *SectionReader) Next() Instruction {
	if int(r.index) >= len(r.section.data) {
		return nil
	}

	sym, _ := r.section.SymbolForIndex(r.index)
	addr := r.section.AddrForIndex(r.index)
	raw := r.section.data[r.index]

	inst := &instruction{
		section: r.section,
		addr:    addr,
		sym:     sym,
		raw:     raw,
	}

	r.index++

	return inst
}
