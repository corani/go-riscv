package riscv

import "fmt"

type Register uint8

func RegisterByName(name string) Register {
	names := map[string]Register{
		"zero": 0,
		"ra":   1, "sp": 2, "gp": 3, "tp": 4,
		"t0": 5, "t1": 6, "t2": 7, "t3": 28,
		"t4": 29, "t5": 30, "t6": 31, "t7": 32,
		"s0": 8, "s1": 9, "s2": 18, "s3": 19,
		"s4": 20, "s5": 21, "s6": 22, "s7": 23,
		"s8": 24, "s9": 25, "s10": 26, "s11": 27,
		"a0": 10, "a1": 11, "a2": 12, "a3": 13,
		"a4": 14, "a5": 15, "a6": 16, "a7": 17,
	}

	if v, ok := names[name]; ok {
		return v
	}

	panic(fmt.Sprintf("illegal register %q", name))
}

func (r Register) String() string {
	names := map[Register]string{
		0: "zero",
		1: "ra", 2: "sp", 3: "gp", 4: "tp",
		5: "t0", 6: "t1", 7: "t2", 28: "t3",
		29: "t4", 30: "t5", 31: "t6", 32: "t7",
		8: "s0", 9: "s1", 18: "s2", 19: "s3",
		20: "s4", 21: "s5", 22: "s6", 23: "s7",
		24: "s8", 25: "s9", 26: "s10", 27: "s11",
		10: "a0", 11: "a1", 12: "a2", 13: "a3",
		14: "a4", 15: "a5", 16: "a6", 17: "a7",
	}

	if v, ok := names[r]; ok {
		return v
	}

	panic(fmt.Sprintf("illegal register %#x", uint8(r)))
}

func (r Register) Index() uint8 {
	return uint8(r)
}
