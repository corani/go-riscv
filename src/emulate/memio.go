package main

import (
	"os"

	"github.com/corani/go-riscv/src/riscv"
)

// implements interface riscv.Section

const (
	memioReadStdin uint32 = iota
	memioWriteStdout
	memioWriteStderr
	memioWriteExit
)

func NewMemIO(base uint32, v *visitor) riscv.Section {
	return &memio{
		base: base,
		size: 4,
		v:    v,
	}
}

type memio struct {
	base uint32
	size uint32
	v    *visitor
}

func (m *memio) Name() string {
	return "<memio>"
}

func (m *memio) Base() uint32 {
	return m.base
}

func (m *memio) Size() uint32 {
	return m.size
}

func (m *memio) Read(addr, size uint32) []byte {
	addr -= m.base

	if addr > m.size || addr+size > m.size {
		return nil
	}

	if size != 1 {
		return nil
	}

	switch addr {
	case memioReadStdin:
		buf := make([]byte, 1)

		os.Stdin.Read(buf)

		return buf
	}

	return nil
}

func (m *memio) Write(addr uint32, data []byte) {
	addr -= m.base

	if addr > m.size || addr+uint32(len(data)) > m.size {
		return
	}

	if len(data) != 1 {
		return
	}

	switch addr {
	case memioWriteStdout:
		os.Stdout.Write(data)
	case memioWriteStderr:
		os.Stderr.Write(data)
	case memioWriteExit:
		m.v.done = true
		m.v.exitCode = int(data[0])
	}
}

func (m *memio) AddSymbol(uint32, string) {

}

func (m *memio) SymbolAt(uint32) (string, bool) {
	return "", false
}

func (m *memio) SymbolBefore(uint32) string {
	return ""
}

func (m *memio) InstructionAt(uint32) riscv.Instruction {
	return nil
}

func (m *memio) Reader() riscv.SectionReader {
	return nil
}
