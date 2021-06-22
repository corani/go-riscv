package main

import "fmt"

type Register uint8

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
