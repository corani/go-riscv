# Go Risc-V

Playing around with Risc-V in Go.

## Setup

1. Install the riscv gnu toolchain from github.com/riscv/riscv-gnu-toolchain, follow the
   instructions for installing Newlib.
2. Clone and build github.com/riscv/riscv-tests under this repository

## Disassembler

Disassembles rv32i with zifenci and zicsr extensions.

```bash
make
./bin/disassemble -in ./riscv-tests/isa/rv32ui-p-simple
```

## Emulator

The emulator needs to be charged with gas before running, each instruction and memory access uses
gas and the emulator is killed when we run out. By default 500 gas is charged, which is sufficient
for (very) simple programs. See the command line arguments below if you need to charge more.

Emulates rv32i with zifenci and zicsr extensions.

```bash
make

./bin/emulate -in ./riscv-tests/isa/rv32ui-p-add
```

To run the emulator for all the riscv-tests:

```bash
make riscv-tests
```

To run the emulator for the samples:

```bash
make
make -C sample

./bin/emulate -in ./sample/bin/hello
```

Additional command line arguments:

- `-gas N`  charge the emulator with `N` gas (default=500)
- `-v N`    verbose logging
  - 1       print a profile after completion
  - 2       print the disassembly before starting
  - 3       print each ecalls before executing
  - 4       print each instruction before executing
  - 5       print all registers after executing each instruction

## Future

- Emulator
  - Memory mapped I/O
  - Proper trapping of ecalls
  - Dynamic memory
- Create an assembler
- Create a debugger
- Support rv32g: rv32m (mul/div) + rv32a (atomic) + rv32f (float) + rv32d (double)?
