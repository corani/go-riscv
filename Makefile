GO := go
CC := /opt/riscv/bin/riscv64-unknown-elf-gcc
CFLAGS := -march=rv32i -mabi=ilp32 -static -nostdlib -Wl,--no-relax

BIN_DIR := bin
SRC_DIR := src

TARGETS := disassemble emulate
TARGETS := $(TARGETS:%=$(BIN_DIR)/%)

GOFILES = $(shell find $(SRC_DIR)/ -type f -name '*.go')

bin/memio: src/emulate/memio.S src/emulate/memio.ld
	$(CC) $(CFLAGS) -T src/emulate/memio.ld -o $@ src/emulate/memio.S

$(BIN_DIR)/%: $(GOFILES)
	@mkdir -p "$(@D)"
	$(GO) build -o $@ $(patsubst $(BIN_DIR)/%,./$(SRC_DIR)/%/,$@)

all: $(TARGETS) bin/memio

riscv-tests: $(TARGETS)
	@test.sh

clean:
	rm -f $(TARGETS) bin/memio

.PHONY: clean all
