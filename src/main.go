package main

import (
	"fmt"
	"os"
)

type (
	u8  = uint8
	u16 = uint16
	u32 = uint32
	u64 = uint64

	s8  = int8
	s16 = int16
	s32 = int32
	s64 = int64

	b32 = int32
)

func DisAsm8086(memory *Memory, disasmByteCount u32, disasmStart SegmentedAccess) {
	at := disasmStart
	var context DecodeContext
	table := Get8086InstructionTable()

	count := disasmByteCount
	for count > 0 {
		instruction := DecodeInstruction(&context, table, memory, &at)
		if instruction.Op != 0 {
			if count >= instruction.Size {
				count -= instruction.Size
			} else {
				fmt.Fprintln(os.Stderr, "ERROR: Instruction extends outside disassembly region")
				break
			}

			if IsPrintable(instruction) {
				PrintInstruction(instruction, os.Stdout)
				fmt.Println()
			}
		} else {
			fmt.Fprintln(os.Stderr, "ERROR: Unrecognized binary in instruction stream.")
			break
		}
	}
}

func main() {
	memory := &Memory{}

	if len(os.Args) > 1 {
		for _, fileName := range os.Args[1:] {
			bytesRead := LoadMemoryFromFile(fileName, memory, 0)

			fmt.Printf("; %s disassembly:\n", fileName)
			fmt.Println("bits 16")
			DisAsm8086(memory, bytesRead, SegmentedAccess{})
		}
	} else {
		fmt.Fprintf(os.Stderr, "USAGE: %s [8086 machine code file] ...\n", os.Args[0])
	}
}
