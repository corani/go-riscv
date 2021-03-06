package riscv

import (
	"fmt"
	"strconv"
)

// TODO: no globals :(
var csrMapping map[uint32]string

func csrGetMapping() map[uint32]string {
	if csrMapping == nil {
		// NOTE: list taken from "The RISC-V Instruction Set Manual - Volume II"
		// (version 20190608-Priv-MSU-Ratified)
		csrMapping = map[uint32]string{
			// user trap setup
			0x000: "ustatus",
			0x004: "uie",
			0x005: "utvec",
			// user trap handling
			0x040: "uscratch",
			0x041: "uepc",
			0x042: "ucause",
			0x043: "utval",
			0x044: "uip",
			// user floating point csrs
			0x001: "fflags",
			0x002: "frm",
			0x003: "fcsr",
			// user counters / timers
			0xc00: "cycle",
			0xc01: "time",
			0xc02: "instret",
			0xc03: "hpmcounter3",
			0xc04: "hpmcounter4",
			0xc05: "hpmcounter5",
			0xc06: "hpmcounter6",
			0xc07: "hpmcounter7",
			0xc08: "hpmcounter8",
			0xc09: "hpmcounter9",
			0xc0a: "hpmcounter10",
			0xc0b: "hpmcounter11",
			0xc0c: "hpmcounter12",
			0xc0d: "hpmcounter13",
			0xc0e: "hpmcounter14",
			0xc0f: "hpmcounter15",
			0xc10: "hpmcounter16",
			0xc11: "hpmcounter17",
			0xc12: "hpmcounter18",
			0xc13: "hpmcounter19",
			0xc14: "hpmcounter20",
			0xc15: "hpmcounter21",
			0xc16: "hpmcounter22",
			0xc17: "hpmcounter23",
			0xc18: "hpmcounter24",
			0xc19: "hpmcounter25",
			0xc1a: "hpmcounter26",
			0xc1b: "hpmcounter27",
			0xc1c: "hpmcounter28",
			0xc1d: "hpmcounter29",
			0xc1e: "hpmcounter30",
			0xc1f: "hpmcounter31",
			0xc80: "cycleh",
			0xc81: "timeh",
			0xc82: "instreth",
			0xc83: "hpmcounter3h",
			0xc84: "hpmcounter4h",
			0xc85: "hpmcounter5h",
			0xc86: "hpmcounter6h",
			0xc87: "hpmcounter7h",
			0xc88: "hpmcounter8h",
			0xc89: "hpmcounter9h",
			0xc8a: "hpmcounter10h",
			0xc8b: "hpmcounter11h",
			0xc8c: "hpmcounter12h",
			0xc8d: "hpmcounter13h",
			0xc8e: "hpmcounter14h",
			0xc8f: "hpmcounter15h",
			0xc90: "hpmcounter16h",
			0xc91: "hpmcounter17h",
			0xc92: "hpmcounter18h",
			0xc93: "hpmcounter19h",
			0xc94: "hpmcounter20h",
			0xc95: "hpmcounter21h",
			0xc96: "hpmcounter22h",
			0xc97: "hpmcounter23h",
			0xc98: "hpmcounter24h",
			0xc99: "hpmcounter25h",
			0xc9a: "hpmcounter26h",
			0xc9b: "hpmcounter27h",
			0xc9c: "hpmcounter28h",
			0xc9d: "hpmcounter29h",
			0xc9e: "hpmcounter30h",
			0xc9f: "hpmcounter31h",

			// supervisor trap setup
			0x100: "sstatus",
			0x102: "sedeleg",
			0x103: "sideleg",
			0x104: "sie",
			0x105: "stvec",
			0x106: "scounteren",
			// supervisor trap handling
			0x140: "sscratch",
			0x141: "sepc",
			0x142: "scause",
			0x143: "stval",
			0x144: "sip",
			// supervisor protection and translation
			0x180: "satp",

			// machine trap setup
			0x300: "mstatus",
			0x301: "misa",
			0x302: "medeleg",
			0x303: "mideleg",
			0x304: "mie",
			0x305: "mtvec",
			0x306: "mcounteren",
			// machine trap handling
			0x340: "mscratch",
			0x341: "mepc",
			0x342: "mcause",
			0x343: "mtval",
			0x344: "mip",
			// machine memory protection
			0x3a0: "pmpcfg0",
			0x3a1: "pmpcfg1",
			0x3a2: "pmpcfg2",
			0x3a3: "pmpcfg3",
			0x3b0: "pmpaddr0",
			0x3b1: "pmpaddr1",
			0x3b2: "pmpaddr2",
			0x3b3: "pmpaddr3",
			0x3b4: "pmpaddr4",
			0x3b5: "pmpaddr5",
			0x3b6: "pmpaddr6",
			0x3b7: "pmpaddr7",
			0x3b8: "pmpaddr8",
			0x3b9: "pmpaddr9",
			0x3ba: "pmpaddr10",
			0x3bb: "pmpaddr11",
			0x3bc: "pmpaddr12",
			0x3bd: "pmpaddr13",
			0x3be: "pmpaddr14",
			0x3bf: "pmpaddr15",

			// machine counter/timers
			0xb00: "mcycle",
			0xb01: "mtime",
			0xb02: "minstret",
			0xb03: "mhpmcounter3",
			0xb04: "mhpmcounter4",
			0xb05: "mhpmcounter5",
			0xb06: "mhpmcounter6",
			0xb07: "mhpmcounter7",
			0xb08: "mhpmcounter8",
			0xb09: "mhpmcounter9",
			0xb0a: "mhpmcounter10",
			0xb0b: "mhpmcounter11",
			0xb0c: "mhpmcounter12",
			0xb0d: "mhpmcounter13",
			0xb0e: "mhpmcounter14",
			0xb0f: "mhpmcounter15",
			0xb10: "mhpmcounter16",
			0xb11: "mhpmcounter17",
			0xb12: "mhpmcounter18",
			0xb13: "mhpmcounter19",
			0xb14: "mhpmcounter20",
			0xb15: "mhpmcounter21",
			0xb16: "mhpmcounter22",
			0xb17: "mhpmcounter23",
			0xb18: "mhpmcounter24",
			0xb19: "mhpmcounter25",
			0xb1a: "mhpmcounter26",
			0xb1b: "mhpmcounter27",
			0xb1c: "mhpmcounter28",
			0xb1d: "mhpmcounter29",
			0xb1e: "mhpmcounter30",
			0xb1f: "mhpmcounter31",
			0xb80: "mcycleh",
			0xb81: "mtimeh",
			0xb82: "minstreth",
			0xb83: "mhpmcounter3h",
			0xb84: "mhpmcounter4h",
			0xb85: "mhpmcounter5h",
			0xb86: "mhpmcounter6h",
			0xb87: "mhpmcounter7h",
			0xb88: "mhpmcounter8h",
			0xb89: "mhpmcounter9h",
			0xb8a: "mhpmcounter10h",
			0xb8b: "mhpmcounter11h",
			0xb8c: "mhpmcounter12h",
			0xb8d: "mhpmcounter13h",
			0xb8e: "mhpmcounter14h",
			0xb8f: "mhpmcounter15h",
			0xb90: "mhpmcounter16h",
			0xb91: "mhpmcounter17h",
			0xb92: "mhpmcounter18h",
			0xb93: "mhpmcounter19h",
			0xb94: "mhpmcounter20h",
			0xb95: "mhpmcounter21h",
			0xb96: "mhpmcounter22h",
			0xb97: "mhpmcounter23h",
			0xb98: "mhpmcounter24h",
			0xb99: "mhpmcounter25h",
			0xb9a: "mhpmcounter26h",
			0xb9b: "mhpmcounter27h",
			0xb9c: "mhpmcounter28h",
			0xb9d: "mhpmcounter29h",
			0xb9e: "mhpmcounter30h",
			0xb9f: "mhpmcounter31h",

			// machine counter setup
			0x320: "mcountinhibit",
			0x323: "mhpevent3",
			0x324: "mhpevent4",
			0x325: "mhpevent5",
			0x326: "mhpevent6",
			0x327: "mhpevent7",
			0x328: "mhpevent8",
			0x329: "mhpevent9",
			0x32a: "mhpevent10",
			0x32b: "mhpevent11",
			0x32c: "mhpevent12",
			0x32d: "mhpevent13",
			0x32e: "mhpevent14",
			0x32f: "mhpevent15",
			0x330: "mhpevent16",
			0x331: "mhpevent17",
			0x332: "mhpevent18",
			0x333: "mhpevent19",
			0x334: "mhpevent20",
			0x335: "mhpevent21",
			0x336: "mhpevent22",
			0x337: "mhpevent23",
			0x338: "mhpevent24",
			0x339: "mhpevent25",
			0x33a: "mhpevent26",
			0x33b: "mhpevent27",
			0x33c: "mhpevent28",
			0x33d: "mhpevent29",
			0x33e: "mhpevent30",
			0x33f: "mhpevent31",

			// debug/trace registers (shared with debug mode)
			0x7a0: "tselect",
			0x7a1: "tdata1",
			0x7a2: "tdata2",
			0x7a3: "tdata3",

			// debug mode registers
			0x7b0: "dcsr",
			0x7b1: "dpc",
			0x7b2: "dscratch0",
			0x7b3: "dscratch1",

			// machine information
			0xf11: "mvendorid",
			0xf12: "marchid",
			0xf13: "mimpid",
			0xf14: "mhartid",
		}
	}

	return csrMapping
}

func CsrName(csr uint32) string {
	names := csrGetMapping()

	if name, ok := names[csr&0xfff]; ok {
		return name
	}

	return fmt.Sprintf("%#x", csr&0xfff)
}

func CsrRegister(name string) uint32 {
	names := csrGetMapping()

	// TODO: this is pretty inefficient!
	for k, v := range names {
		if v == name {
			return k
		}
	}

	// NOTE: in case we didn't find the csr register by name, try to parse the name
	// as a hex-string.
	if v, err := strconv.ParseUint(name, 16, 32); err != nil {
		return uint32(v & 0xfff)
	}

	return 0
}
