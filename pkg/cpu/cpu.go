package cpu

import "github.com/kijimaD/minicpu/pkg/asm"

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

func (cpu *CPU) SetROM() {
	asm := asm.Assembler{}
	cpu.ROM[0] = asm.Ldl(0, 3)
	cpu.ROM[1] = asm.Ldh(0, 1)
	cpu.ROM[2] = asm.Add(0, 0) // to, from
}
