package main

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

type Instruction struct {
	addr uint32
	raw  uint32
	sym  string
	text string
}

func (s Section) SymbolAtAddr(addr uint32) (string, bool) {
	sym, ok := s.symbols[addr]
	return sym, ok
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

func (r *SectionReader) Next() *Instruction {
	if int(r.index) >= len(r.section.data) {
		return nil
	}

	sym, _ := r.section.SymbolForIndex(r.index)
	addr := r.section.AddrForIndex(r.index)
	raw := r.section.data[r.index]

	inst := Instruction{
		addr: addr,
		sym:  sym,
		raw:  raw,
		text: decode(r.section, addr, raw),
	}

	r.index++

	return &inst
}
