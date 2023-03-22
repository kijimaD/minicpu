package cpu

import (
	"github.com/kijimaD/minicpu/pkg/asm"
	"github.com/kijimaD/minicpu/pkg/types"
)

type Regs [8]Register
type Register = uint16

type CPU struct {
	PC   uint16
	Regs Regs
	ROM  [256]uint16
	RAM  [256]uint16
}

// NewCPU is CPU constructor
func NewCPU() *CPU {
	cpu := &CPU{
		PC:   0x00,
		Regs: Regs{},
		ROM:  [256]uint16{},
		RAM:  [256]uint16{},
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
	var inst *inst
	inst = instructions[opcode]
	inst.Execute(cpu, operands)
}

func (cpu *CPU) mov(ra, rb Register) {
	cpu.Regs[ra] = cpu.Regs[rb]
}

func (cpu *CPU) add(ra, rb Register) {
	ra = ra + rb
}

func (cpu *CPU) ldl(r Register, val uint16) {
	cpu.Regs[r] = (r & 0xff00) | (val & 0x00ff)
}

func (cpu *CPU) SetROM() {
	asm := asm.Assembler{}
	cpu.ROM[0] = asm.Ldl(0, 3)
	cpu.ROM[1] = asm.Ldh(0, 1)
	cpu.ROM[2] = asm.Add(0, 0) // to, from
}

// 命令
type inst struct {
	Opcode      byte // なので8種類しかない
	Description string
	Execute     func(cpu *CPU, operands []uint16)
}

var instructions = []*inst{
	&inst{types.MOV, "mov", func(cpu *CPU, operands []uint16) { cpu.mov(operands[0], operands[1]) }},
	&inst{types.ADD, "add", func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
	&inst{types.ADD, "add", func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
	&inst{types.ADD, "add", func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
	&inst{types.ADD, "add", func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
	&inst{types.ADD, "add", func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
	&inst{types.ADD, "add", func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
	&inst{types.ADD, "add", func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
	&inst{types.LDL, "ldl", func(cpu *CPU, operands []uint16) { cpu.ldl(operands[0], operands[2]) }},
}
