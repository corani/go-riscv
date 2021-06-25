#!/bin/bash
/opt/riscv/bin/riscv64-unknown-elf-gcc -march=rv32i -mabi=ilp32 -static -nostdlib -o hello hello.S
