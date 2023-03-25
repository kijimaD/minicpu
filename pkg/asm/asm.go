package asm

import "github.com/kijimaD/minicpu/pkg/types"

type Assembler struct{}

// move
func (asm *Assembler) MOV(ra, rb uint16) uint16 {
	return ((types.MOV << 11) | (ra << 8) | (rb << 5))
}

// addition
func (asm *Assembler) ADD(ra, rb uint16) uint16 {
	return ((types.ADD << 11) | (ra << 8) | (rb << 5))
}

// subtraction
func (asm *Assembler) SUB(ra, rb uint16) uint16 {
	return ((types.SUB << 11) | (ra << 8) | (rb << 5))
}

// and
func (asm *Assembler) AND(ra, rb uint16) uint16 {
	return ((types.AND << 11) | (ra << 8) | (rb << 5))
}

// or
func (asm *Assembler) OR(ra, rb uint16) uint16 {
	return ((types.OR << 11) | (ra << 8) | (rb << 5))
}

// shift left
func (asm *Assembler) SL(ra uint16) uint16 {
	return ((types.SL << 11) | (ra << 8))
}

// shift right
func (asm *Assembler) SR(ra uint16) uint16 {
	return ((types.SR << 11) | (ra << 8))
}

// shift right arithmetic
func (asm *Assembler) SRA(ra uint16) uint16 {
	return ((types.SRA << 11) | (ra << 8))
}

// load immediate value low
func (asm *Assembler) LDL(ra, val uint16) uint16 {
	return ((types.LDL << 11) | (ra << 8) | (val & 0x00ff))
}

// load immediate value high
func (asm *Assembler) LDH(ra, val uint16) uint16 {
	return ((types.LDH << 11) | (ra << 8) | (val & 0x00ff))
}

// compare
func (asm *Assembler) CMP(ra, rb uint16) uint16 {
	return ((types.CMP << 11) | (ra << 8) | (rb << 5))
}

// jump equal
func (asm *Assembler) JE(addr uint16) uint16 {
	return ((types.JE << 11) | (addr & 0x00ff))
}

// jump
func (asm *Assembler) JMP(addr uint16) uint16 {
	return ((types.JMP << 11) | (addr & 0x00ff))
}

// load memory
func (asm *Assembler) LD(ra, addr uint16) uint16 {
	return ((types.LD << 11) | (ra << 8) | (addr & 0x00ff))
}

// store memory
func (asm *Assembler) ST(ra, addr uint16) uint16 {
	return ((types.ST << 11) | (ra << 8) | (addr & 0x00ff))
}

// halt
func (asm *Assembler) HLT() uint16 {
	return (types.HLT << 11)
}

// op_code
func (asm *Assembler) Opcode(ir uint16) uint16 {
	return (ir >> 11)
}

func (asm *Assembler) RegA(ir uint16) uint16 {
	return ((ir >> 8) & 0x0007)
}

func (asm *Assembler) RegB(ir uint16) uint16 {
	return ((ir >> 5) & 0x0007)
}

func (asm *Assembler) Data(ir uint16) uint16 {
	return (ir & 0x00ff)
}

func (asm *Assembler) Addr(ir uint16) uint16 {
	return (ir & 0x00ff)
}
