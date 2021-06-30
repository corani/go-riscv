CC := /opt/riscv/bin/riscv64-unknown-elf-gcc
CFLAGS := -march=rv32i -mabi=ilp32 -static -nostdlib -Wl,--no-relax

BIN_DIR := bin
SRC_DIR := src

# Generate a target for each source file
TARGETS := $(wildcard $(SRC_DIR)/*.S)
TARGETS := $(TARGETS:$(SRC_DIR)/%.S=$(BIN_DIR)/%)

$(BIN_DIR)/%: $(SRC_DIR)/%.S
	@mkdir -p $(dir $@)
	$(CC) $(CFLAGS) -o $@ $^

all: $(TARGETS)

clean:
	rm -f $(TARGETS)

.PHONY: clean all
