#!/bin/bash
mkdir -p bin
go build -o bin/disassemble ./src/disassemble/
go build -o bin/emulate     ./src/emulate/
