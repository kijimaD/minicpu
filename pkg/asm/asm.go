package asm

import "github.com/kijimaD/minicpu/pkg/types"

type Assembler struct{}

func (asm *Assembler) Add(ra, rb uint16) uint16 {
	return ((types.ADD << 11) | (ra << 8) | (rb << 5))
}

// load immediate value low
func (asm *Assembler) Ldl(ra, val uint16) uint16 {
	return ((types.LDL << 11) | (ra << 8) | (val & 0x00ff))
}
