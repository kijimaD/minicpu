package cpu

import (
	"github.com/kijimaD/minicpu/pkg/types"
)

// move
func (cpu *CPU) mov(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[rb]
}

// addition
func (cpu *CPU) add(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[ra] + cpu.Regs[rb]
}

// subtraction
func (cpu *CPU) sub(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[ra] - cpu.Regs[rb]
}

// logical and
func (cpu *CPU) and(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[ra] & cpu.Regs[rb]
}

// logical or
func (cpu *CPU) or(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[ra] | cpu.Regs[rb]
}

// shift left
func (cpu *CPU) sl(ra Register) {
	cpu.Regs[ra] = cpu.Regs[ra] << 1
}

// shift right
func (cpu *CPU) sr(ra Register) {
	cpu.Regs[ra] = cpu.Regs[ra] >> 1
}

// shift right arithmetic
func (cpu *CPU) sra(ra Register) {
	cpu.Regs[ra] = (cpu.Regs[ra] & 0x8000) | (cpu.Regs[ra] >> 1)
}

// load immediate value low
func (cpu *CPU) ldl(r Register, val uint16) {
	cpu.Regs[r] = (r & 0xff00) | (val & 0x00ff)
}

// load immediate value high
func (cpu *CPU) ldh(ra Register, val uint16) {
	cpu.Regs[ra] = (val << 8) | (cpu.Regs[ra] & 0x00ff)
}

// compare
func (cpu *CPU) cmp(ra, rb Register) {
	if cpu.Regs[ra] == cpu.Regs[rb] {
		cpu.EQFlag = true
	} else {
		cpu.EQFlag = false
	}
}

// jump equal
func (cpu *CPU) je(val uint16) {
	if cpu.EQFlag == true {
		cpu.PC = val
	}
}

// jump
func (cpu *CPU) jmp(val uint16) {
	cpu.PC = val
}

// load memory
func (cpu *CPU) ld(ra Register, val uint16) {
	cpu.Regs[ra] = cpu.RAM[val]
}

// store memory
func (cpu *CPU) st(ra Register, val uint16) {
	cpu.RAM[val] = cpu.Regs[ra]
}

// halt
func (cpu *CPU) hlt() {
	cpu.Halted = true
}

// 命令
type inst struct {
	Opcode      byte // byteなので8種類しかない
	Description string
	Execute     func(cpu *CPU, operands []uint16)
}

var Instructions = []*inst{
	&inst{types.MOV, "MOV", func(cpu *CPU, operands []uint16) { cpu.mov(operands[0], operands[1]) }},
	&inst{types.SUB, "ADD", func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
	&inst{types.ADD, "SUB", func(cpu *CPU, operands []uint16) { cpu.sub(operands[0], operands[1]) }},
	&inst{types.ADD, "AND", func(cpu *CPU, operands []uint16) { cpu.and(operands[0], operands[1]) }},
	&inst{types.OR, "OR", func(cpu *CPU, operands []uint16) { cpu.or(operands[0], operands[1]) }},
	&inst{types.SL, "SL", func(cpu *CPU, operands []uint16) { cpu.sl(operands[0]) }},
	&inst{types.SR, "SR", func(cpu *CPU, operands []uint16) { cpu.sr(operands[0]) }},
	&inst{types.SRA, "SRA", func(cpu *CPU, operands []uint16) { cpu.sra(operands[0]) }},
	&inst{types.LDL, "LDL", func(cpu *CPU, operands []uint16) { cpu.ldl(operands[0], operands[2]) }},
	&inst{types.LDH, "LDH", func(cpu *CPU, operands []uint16) { cpu.ldh(operands[0], operands[2]) }},
	&inst{types.CMP, "CMP", func(cpu *CPU, operands []uint16) { cpu.cmp(operands[0], operands[1]) }},
	&inst{types.JE, "JE", func(cpu *CPU, operands []uint16) { cpu.je(operands[2]) }},
	&inst{types.JMP, "JMP", func(cpu *CPU, operands []uint16) { cpu.jmp(operands[2]) }},
	&inst{types.LD, "LD", func(cpu *CPU, operands []uint16) { cpu.ld(operands[0], operands[2]) }},
	&inst{types.ST, "ST", func(cpu *CPU, operands []uint16) { cpu.st(operands[0], operands[2]) }},
	&inst{types.HLT, "HLT", func(cpu *CPU, operands []uint16) { cpu.hlt() }},
}
