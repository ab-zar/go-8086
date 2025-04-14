package main

import "fmt"

type OpType int

const (
	Op_None OpType = iota

	// Generated from sim86_instruction_table.inl
	//go:generate stringer -type=OpType
	// The list below must match the order of the INST macros in the .inl file
	Op_mov
	Op_push
	Op_pop
	Op_xchg
	Op_in
	Op_out
	Op_xlat
	Op_lea
	Op_lds
	Op_les
	Op_lahf
	Op_sahf
	Op_pushf
	Op_popf
	Op_add
	Op_adc
	Op_inc
	Op_aaa
	Op_daa
	Op_sub
	Op_sbb
	Op_dec
	Op_neg
	Op_cmp
	Op_aas
	Op_das
	Op_mul
	Op_imul
	Op_aam
	Op_div
	Op_idiv
	Op_aad
	Op_cbw
	Op_cwd
	Op_not
	Op_shl
	Op_shr
	Op_sar
	Op_rol
	Op_ror
	Op_rcl
	Op_rcr
	Op_and
	Op_test
	Op_or
	Op_xor
	Op_rep
	Op_movs
	Op_cmps
	Op_scas
	Op_lods
	Op_stos
	Op_call
	Op_jmp
	Op_ret
	Op_je
	Op_jl
	Op_jle
	Op_jb
	Op_jbe
	Op_jp
	Op_jo
	Op_js
	Op_jne
	Op_jnl
	Op_jg
	Op_jnb
	Op_ja
	Op_jnp
	Op_jno
	Op_jns
	Op_loop
	Op_loopz
	Op_loopnz

	Op_Count
)

type InstructionBitsUsage uint8

const (
	Bits_End InstructionBitsUsage = iota
	Bits_Literal
	Bits_D
	Bits_S
	Bits_W
	Bits_V
	Bits_Z
	Bits_MOD
	Bits_REG
	Bits_RM
	Bits_SR
	Bits_Disp
	Bits_Data
	Bits_DispAlwaysW
	Bits_WMakesDataW
	Bits_RMRegAlwaysW
	Bits_RelJMPDisp
	Bits_Count
)

type InstructionBit struct {
	Usage    InstructionBitsUsage
	BitCount uint8
	Shift    uint8
	Value    uint8
}

type InstructionEncoding struct {
	Op   OpType
	Bits [16]InstructionBit
}

type InstructionTable struct {
	Encodings     []InstructionEncoding
	EncodingCount int
}

var instructionTable8086 = []InstructionEncoding{
	// This will be generated from sim86_instruction_table.inl translated into Go.
	// Placeholder for now
}

func Get8086InstructionTable() InstructionTable {
	return InstructionTable{
		Encodings:     instructionTable8086,
		EncodingCount: len(instructionTable8086),
	}
}

func GetMnemonic(op OpType) string {
	return op.String()
}

func (op OpType) String() string {
	switch op {
	case Op_None:
		return ""
	case Op_mov:
		return "mov"
	case Op_push:
		return "push"
	case Op_pop:
		return "pop"
	// TODO: Add all other cases...
	default:
		return fmt.Sprintf("UNKNOWN(%d)", int(op))
	}
}
