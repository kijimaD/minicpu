package main

import (
	"fmt"

	"github.com/kijimaD/minicpu/pkg/asm"
)

type Regs [8]Register
type Register = uint16

type CPU struct {
	PC   int
	Regs Regs
	ROM  [256]uint16
	RAM  [256]uint16
}

// NewCPU is CPU constructor
func NewCPU() *CPU {
	cpu := &CPU{
		PC:   0x0100, // INFO: Skip
		Regs: Regs{0x0000, 0x0000, 0x0000, 0x0000, 0x0000, 0x0000, 0x0000, 0x0000},
	}
	return cpu
}

func (cpu *CPU) fetch() uint16 {
	d := cpu.ROM[cpu.PC]
	cpu.PC++
	return d
}

func (cpu *CPU) fetchOperands(size uint) []uint16 {
	operands := []uint16{}
	switch size {
	case 1:
		operands = append(operands, cpu.fetch())
	case 2:
		operands = append(operands, cpu.fetch())
		operands = append(operands, cpu.fetch())
	}
	return operands
}

func (cpu *CPU) Step() {
	opcode := cpu.fetch()
	var inst *inst
	inst = instructions[opcode]

	operands := cpu.fetchOperands(inst.OperandsSize)
	inst.Execute(cpu, operands)
}

// 命令
type inst struct {
	Opcode       byte
	Description  string
	OperandsSize uint
	Execute      func(cpu *CPU, operands []uint16)
}

var instructions = []*inst{
	&inst{0x0, "add", 3, func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
}

func (cpu *CPU) add(ra, rb Register) {
	ra = ra | rb
}

// ここには生の値が入っている(byte)
func (cpu *CPU) setROM() {
	asm := asm.Assembler{}
	cpu.ROM[0] = asm.Ldl(cpu.Regs[0], 3)
	cpu.ROM[1] = asm.Add(cpu.Regs[1], cpu.Regs[0]) // to, from
}

func main() {
	cpu := NewCPU()
	cpu.setROM()
	fmt.Printf("%#v", cpu)
}
