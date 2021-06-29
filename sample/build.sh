#!/bin/bash
/opt/riscv/bin/riscv64-unknown-elf-gcc -march=rv32i -mabi=ilp32 -static -nostdlib -Wl,--no-relax -o hello hello.S
/opt/riscv/bin/riscv64-unknown-elf-gcc -march=rv32i -mabi=ilp32 -static -nostdlib -Wl,--no-relax -o greet greet.S
