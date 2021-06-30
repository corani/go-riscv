GO := go

BIN_DIR := bin
SRC_DIR := src

TARGETS := disassemble emulate
TARGETS := $(TARGETS:%=$(BIN_DIR)/%)

GOFILES = $(shell find $(SRC_DIR)/ -type f -name '*.go')

$(BIN_DIR)/%: $(GOFILES)
	@mkdir -p "$(@D)"
	$(GO) build -o $@ $(patsubst $(BIN_DIR)/%,./$(SRC_DIR)/%/,$@)

all: $(TARGETS)

riscv-tests: $(TARGETS)
	@test.sh

clean:
	rm -f $(TARGETS)

.PHONY: clean all
