package main

import "fmt"

func main() {
	x := 0x87654321

	// A.
	fmt.Printf("A: %x\n", x&0xFF)
	// B.
	fmt.Printf("B: %x\n", x^0xFFFFFF00)
	// C.
	fmt.Printf("C: %x\n", x|0xFF)
}
