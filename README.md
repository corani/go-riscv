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
./bin/disassemble
```

This disassembles a hardcoded elf-binary from the riscv-tests.

## Future

- Don't hardcode the test elf-binary for disassembler
- Disassemble all riscv-tests and validate output
- Create an emulator
- Create an assembler
- Create a debugger
- Support rv32m, rv32f?
