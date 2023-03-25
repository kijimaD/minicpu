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
	Halted bool
	Curop  uint16
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
		Halted: false,
		Curop:  0x00,
	}
	return cpu
}

func (cpu *CPU) fetch() uint16 {
	d := cpu.ROM[cpu.PC]
	cpu.IR = d
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
	cpu.Curop = opcode

	var inst *inst
	inst = Instructions[opcode]
	inst.Execute(cpu, operands)
}

func (cpu *CPU) SetROM() {
	// 1+2+...+10=55
	asm := asm.Assembler{}
	cpu.ROM[0] = asm.LDH(types.REG0, 0)
	cpu.ROM[1] = asm.LDL(types.REG0, 0)
	cpu.ROM[2] = asm.LDH(types.REG1, 0)
	cpu.ROM[3] = asm.LDL(types.REG1, 1)
	cpu.ROM[4] = asm.LDH(types.REG2, 0)
	cpu.ROM[5] = asm.LDL(types.REG2, 0)
	cpu.ROM[6] = asm.LDH(types.REG3, 0)
	cpu.ROM[7] = asm.LDL(types.REG3, 10)
	cpu.ROM[8] = asm.ADD(types.REG2, types.REG1)
	cpu.ROM[9] = asm.ADD(types.REG0, types.REG2)
	cpu.ROM[10] = asm.ST(types.REG0, 64)
	cpu.ROM[11] = asm.CMP(types.REG2, types.REG3)
	cpu.ROM[12] = asm.JE(14)
	cpu.ROM[13] = asm.JMP(8)
	cpu.ROM[14] = asm.HLT()
}
