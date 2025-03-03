package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage go run main.go <asm_file>")
		return
	}

	fliename := os.Args[1]

	data, err := os.ReadFile(fliename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("First 6 bits of each byte:")
	for i, b := range data {
		first6Bits := (b >> 2) & 0x3F // Extract the first 6 bits
		fmt.Printf("Byte %d: 0x%02X -> First 6 bits: %06b\n", i, b, first6Bits)
	}
}
