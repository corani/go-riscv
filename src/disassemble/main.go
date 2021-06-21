package main

func verify(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	program, err := loadElf("./riscv-tests/isa/rv32ui-p-add")
	verify(err)

	listProgram(program)
}
