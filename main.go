package main

import (
	"fmt"
	"os"
)

const (
	MOV uint8 = 0b100010
)

var (
	opcode uint8
	d      direction
	w      wordByte
	sp     uint8
)

type direction uint8

// direction
const (
	fromRegister direction = iota
	toRegister
)

type wordByte uint8

// word or byte
const (
	byteOP wordByte = iota
	wordOp
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
		opcode = (b >> 2) & 0x3F // Extract the first 6 bits

		switch opcode {
		case MOV:
			fmt.Println("MOV")
		default:
			fmt.Println("Unknown opcode:", opcode)
		}
		d := direction((b >> 1) & 0b1) //Extract direction
		w := b & 0b1

		fmt.Printf("Byte %d: 0x%02X -> opcode: %06b\n", i, b, opcode)
		fmt.Printf("Byte %d: 0x%02X -> d: %1b\n", i, b, d)
		fmt.Printf("Byte %d: 0x%02X -> w: %01b\n", i, b, w)
	}
}
