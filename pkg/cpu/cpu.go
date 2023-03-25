package cpu

import (
	"github.com/kijimaD/minicpu/pkg/asm"
	"github.com/kijimaD/minicpu/pkg/types"
)

type Regs [8]Register
type Register = uint16

type CPU struct {
	PC     uint16
	IR     uint16
	EQFlag bool
	Regs   Regs
	ROM    [256]uint16
	RAM    [256]uint16
}

// NewCPU is CPU constructor
func NewCPU() *CPU {
	cpu := &CPU{
		PC:     0x00,
		IR:     0x00,
		EQFlag: false,
		Regs:   Regs{},
		ROM:    [256]uint16{},
		RAM:    [256]uint16{},
	}
	return cpu
}

func (cpu *CPU) fetch() uint16 {
	d := cpu.ROM[cpu.PC]
	cpu.PC++
	return cpu.opcode(d)
}

func (cpu *CPU) fetchOperands() []uint16 {
	d := cpu.ROM[cpu.PC]
	operands := []uint16{}
	var result uint16
	{
		result = cpu.opregA(d)
		operands = append(operands, result)
	}
	{
		result = cpu.opregB(d)
		operands = append(operands, result)
	}
	{
		result = cpu.opdata(d)
		operands = append(operands, result)
	}

	return operands
}

func (cpu *CPU) opcode(line uint16) uint16 {
	return (line >> 11)
}

func (cpu *CPU) opregA(line uint16) uint16 {
	return ((line >> 8) & 0x0007)
}

func (cpu *CPU) opregB(line uint16) uint16 {
	return ((line >> 5) & 0x0007)
}

func (cpu *CPU) opdata(line uint16) uint16 {
	return (line & 0x00ff)
}

func (cpu *CPU) Step() {
	operands := cpu.fetchOperands()
	opcode := cpu.fetch()
	cpu.IR = opcode

	var inst *inst
	inst = instructions[opcode]
	inst.Execute(cpu, operands)
}

func (cpu *CPU) mov(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[rb]
}

func (cpu *CPU) add(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[ra] + cpu.Regs[rb]
}

func (cpu *CPU) sub(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[ra] - cpu.Regs[rb]
}

func (cpu *CPU) and(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[ra] & cpu.Regs[rb]
}

func (cpu *CPU) or(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[ra] | cpu.Regs[rb]
}

func (cpu *CPU) sl(ra Register) {
	cpu.Regs[ra] = cpu.Regs[ra] << 1
}

func (cpu *CPU) sr(ra Register) {
	cpu.Regs[ra] = cpu.Regs[ra] >> 1
}

func (cpu *CPU) sra(ra Register) {
	cpu.Regs[ra] = (cpu.Regs[ra] & 0x8000) | (cpu.Regs[ra] >> 1)
}

func (cpu *CPU) ldl(r Register, val uint16) {
	cpu.Regs[r] = (r & 0xff00) | (val & 0x00ff)
}

func (cpu *CPU) ldh(ra Register, val uint16) {
	cpu.Regs[ra] = (val << 8) | (cpu.Regs[ra] & 0x00ff)
}

func (cpu *CPU) cmp(ra, rb Register) {
	if cpu.Regs[ra] == cpu.Regs[rb] {
		cpu.EQFlag = true
	} else {
		cpu.EQFlag = false
	}
}

func (cpu *CPU) je(val uint16) {
	if cpu.EQFlag == true {
		cpu.PC = val
	}
}

func (cpu *CPU) jmp(val uint16) {
	cpu.PC = val
}

func (cpu *CPU) ld(ra Register, val uint16) {
	cpu.Regs[ra] = cpu.RAM[val]
}

func (cpu *CPU) st(ra Register, val uint16) {
	cpu.RAM[val] = cpu.Regs[ra]
}

func (cpu *CPU) SetROM() {
	asm := asm.Assembler{}
	cpu.ROM[0] = asm.Ldl(0, 3)
	cpu.ROM[1] = asm.Ldh(0, 1)
	cpu.ROM[2] = asm.Add(0, 0) // to, from
}

// 命令
type inst struct {
	Opcode      byte // byteなので8種類しかない
	Description string
	Execute     func(cpu *CPU, operands []uint16)
}

var instructions = []*inst{
	&inst{types.MOV, "mov", func(cpu *CPU, operands []uint16) { cpu.mov(operands[0], operands[1]) }},
	&inst{types.SUB, "add", func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
	&inst{types.ADD, "sub", func(cpu *CPU, operands []uint16) { cpu.sub(operands[0], operands[1]) }},
	&inst{types.ADD, "and", func(cpu *CPU, operands []uint16) { cpu.and(operands[0], operands[1]) }},
	&inst{types.OR, "or", func(cpu *CPU, operands []uint16) { cpu.or(operands[0], operands[1]) }},
	&inst{types.SL, "sl", func(cpu *CPU, operands []uint16) { cpu.sl(operands[0]) }},
	&inst{types.SR, "sr", func(cpu *CPU, operands []uint16) { cpu.sr(operands[0]) }},
	&inst{types.SRA, "sra", func(cpu *CPU, operands []uint16) { cpu.sra(operands[0]) }},
	&inst{types.LDL, "ldl", func(cpu *CPU, operands []uint16) { cpu.ldl(operands[0], operands[2]) }},
	&inst{types.LDH, "ldh", func(cpu *CPU, operands []uint16) { cpu.ldh(operands[0], operands[2]) }},
	&inst{types.CMP, "cmp", func(cpu *CPU, operands []uint16) { cpu.cmp(operands[0], operands[1]) }},
	&inst{types.JE, "je", func(cpu *CPU, operands []uint16) { cpu.je(operands[2]) }},
	&inst{types.JMP, "jmp", func(cpu *CPU, operands []uint16) { cpu.jmp(operands[2]) }},
	&inst{types.LD, "ld", func(cpu *CPU, operands []uint16) { cpu.ld(operands[0], operands[2]) }},
	&inst{types.ST, "st", func(cpu *CPU, operands []uint16) { cpu.st(operands[0], operands[2]) }},
}
