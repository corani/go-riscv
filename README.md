# Go Risc-V

Playing around with Risc-V in Go.

## Setup

1. Install the riscv gnu toolchain from github.com/riscv/riscv-gnu-toolchain, follow the
   instructions for installing Newlib.
2. Clone and build github.com/riscv/riscv-tests under this repository

## Disassembler

Disassembles rv32i with zifenci and zicsr extensions.

```bash
./build.sh
./bin/disassemble -in ./riscv-tests/isa/rv32ui-p-simple
```

## Future

- Create an emulator
- Create an assembler
- Create a debugger
- Support rv32g: rv32m (mul/div) + rv32a (atomic) + rv32f (float) + rv32d (double)?
