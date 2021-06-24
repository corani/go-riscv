package riscv

import (
	"fmt"
	"math"
)

type Program interface {
	Name() string
	Machine() string
	Class() string
	Order() string
	Sections() []Section
	AddSection(s Section)
}

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
	Reader() *SectionReader
}

type program struct {
	name     string
	machine  string
	class    string
	order    string
	sections []Section
}

type section struct {
	name    string
	base    uint32
	size    uint32
	symbols map[uint32]string
	data    []uint32
}

type SectionReader struct {
	section Section
	index   uint32
}

type Instruction interface {
	Addr() uint32
	Raw() uint32
	Sym() string
	Text() string
}

func NewProgram(name, machine, class, order string) Program {
	return &program{
		name:    name,
		machine: machine,
		class:   class,
		order:   order,
	}
}

func (p *program) Name() string {
	return p.name
}

func (p *program) Machine() string {
	return p.machine
}

func (p *program) Class() string {
	return p.class
}

func (p *program) Order() string {
	return p.order
}

func (p *program) Sections() []Section {
	return p.sections
}

func (p *program) AddSection(s Section) {
	p.sections = append(p.sections, s)
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

func (s *section) Reader() *SectionReader {
	return &SectionReader{
		section: s,
		index:   0,
	}
}

func (r *SectionReader) Next() Instruction {
	s := r.section.(*section)

	if int(r.index) >= len(s.data) {
		return nil
	}

	sym, _ := r.section.SymbolForIndex(r.index)
	addr := r.section.AddrForIndex(r.index)
	raw := s.data[r.index]

	inst := decodeInstruction(r.section, addr, raw, sym)

	r.index++

	return inst
}
