# Go Risc-V

Playing around with Risc-V in Go.

## Setup

1. Install the riscv gnu toolchain from github.com/riscv/riscv-gnu-toolchain, follow the instructions for installing Newlib.
2. Clone and build github.com/riscv/riscv-tests under this repository

## Disassembler

```bash
./build.sh
./bin/disassemble
```

This disassembles a hardcoded elf-binary from the riscv-tests.
