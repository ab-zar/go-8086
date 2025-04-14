package main

import (
	"fmt"
	"os"
)

const (
	MemorySize       = 1024 * 1024
	MemoryAccessMask = 0xfffff
)

type Memory struct {
	Bytes [MemorySize]u8
}

type SegmentedAccess struct {
	SegmentBase   u16
	SegmentOffset u16
}

func GetAbsoluteAddressOf(segmentBase u16, segmentOffset u16, additionalOffset u16) u32 {
	return (u32(segmentBase)<<4 + u32(segmentOffset+additionalOffset)) & MemoryAccessMask
}

func GetAbsoluteAddress(access SegmentedAccess, additionalOffset u16) u32 {
	return GetAbsoluteAddressOf(access.SegmentBase, access.SegmentOffset, additionalOffset)
}

func ReadMemory(memory *Memory, absoluteAddress u32) u8 {
	if absoluteAddress >= MemorySize {
		panic(fmt.Sprintf("ReadMemory: address out of bounds: %d", absoluteAddress))
	}
	return memory.Bytes[absoluteAddress]
}

func LoadMemoryFromFile(fileName string, memory *Memory, atOffset u32) u32 {
	if atOffset >= MemorySize {
		return 0
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Unable to open %s.\n", fileName)
		return 0
	}
	defer file.Close()

	n, err := file.Read(memory.Bytes[atOffset:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Reading %s failed: %v\n", fileName, err)
		return 0
	}

	return u32(n)
}
