package main

import "fmt"

const ADD uint16 = 1

const (
	MOV = 0
	// ADD  = 1
	SUB  = 2
	OR   = 4
	SL   = 5
	SR   = 6
	SRA  = 7
	LDL  = 8
	LDH  = 9
	CMP  = 10
	JE   = 11
	JMP  = 12
	LD   = 13
	ST   = 14
	HLT  = 15
	REG0 = 0
	REG1 = 1
	REG2 = 2
	REG3 = 3
	REG4 = 4
	REG5 = 5
	REG6 = 6
	REG7 = 7
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
	// Cycles       uint
	Execute func(cpu *CPU, operands []uint16)
}

var instructions = []*inst{
	&inst{0x0, "add", 3, func(cpu *CPU, operands []uint16) { cpu.add(operands[0], operands[1]) }},
}

func (cpu *CPU) mov(ra, rb Register) uint16 {
	return ((MOV << 11) | (ra << 8) | (rb << 5))
}

func (cpu *CPU) add(ra, rb Register) {
	ra = ra | rb
}

// asm ================

type Assembler struct{}

func (asm *Assembler) add(ra, rb uint16) uint16 {
	return ((ADD << 11) | (ra << 8) | (rb << 5))
}

// RUN ================

// ここには生の値が入っている(byte)
func (cpu *CPU) setROM() {
	asm := Assembler{}
	cpu.ROM[0] = asm.add(cpu.Regs[2], cpu.Regs[1])
}

func main() {
	cpu := NewCPU()
	cpu.setROM()
	fmt.Println(cpu)
}
