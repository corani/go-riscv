package main

type instruction struct {
	addr uint32
	raw  uint32
	sym  string
	text string
}

func (i *instruction) Addr() uint32 {
	return i.addr
}

func (i *instruction) Raw() uint32 {
	return i.raw
}

func (i *instruction) Sym() string {
	return i.sym
}

func (i *instruction) Text() string {
	return i.text
}
