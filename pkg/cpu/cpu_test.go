package cpu_test

import (
	"testing"

	"github.com/kijimaD/minicpu/pkg/cpu"
	"github.com/stretchr/testify/assert"
)

func setup(data []uint16) *cpu.CPU {
	cpu := &cpu.CPU{
		PC:   0,
		Regs: cpu.Regs{},
		ROM:  [256]uint16{},
		RAM:  [256]uint16{},
	}

	for i, d := range data {
		cpu.ROM[i] = d
	}

	return cpu
}

func TestMOV(t *testing.T) {
	// opcode, regA, regB
	// 0, 1, 0
	// 0, 1, 5
	c := setup([]uint16{0b0000_001_000_00000, 0b0000_001_101_00000})
	c.Regs = cpu.Regs{0x3, 0x5}
	c.Step()
	assert.Equal(t, uint16(0x0), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x3), c.Regs[1])
	c.Step()
	assert.Equal(t, uint16(0x0), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x0), c.Regs[1])
}

func TestADD(t *testing.T) {
	// opcode, regA, regB
	// 1, 1, 2
	c := setup([]uint16{0b0001_001_010_00000})
	c.Regs = cpu.Regs{0x3, 0x1, 0x2}
	c.Step()

	assert.Equal(t, uint16(0x1), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x3), c.Regs[1])
	assert.Equal(t, uint16(0x2), c.Regs[2])
}

func TestSUB(t *testing.T) {
	// opcode, regA, regB
	// 2, 1, 2
	c := setup([]uint16{0b0010_001_010_00000})
	c.Regs = cpu.Regs{0x1, 0x7, 0x5}
	c.Step()
	assert.Equal(t, uint16(0x2), c.IR)
	assert.Equal(t, uint16(0x1), c.Regs[0])
	assert.Equal(t, uint16(0x2), c.Regs[1])
	assert.Equal(t, uint16(0x5), c.Regs[2])
}

func TestAND(t *testing.T) {
	// opcode, regA, regB
	// 3, 1, 2
	c := setup([]uint16{0b0011_001_010_00000})
	c.Regs = cpu.Regs{0x3, 0x1, 0x2}
	c.Step()
	assert.Equal(t, uint16(0x3), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x0), c.Regs[1])
	assert.Equal(t, uint16(0x2), c.Regs[2])
	// 0b01
	// 0b10
	// 0b00 => 0x0
}

func TestOR(t *testing.T) {
	// opcode, regA, regB
	// 4, 1, 2
	c := setup([]uint16{0b0100_001_010_00000})
	c.Regs = cpu.Regs{0x3, 0x1, 0x2}
	c.Step()
	assert.Equal(t, uint16(0x4), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x3), c.Regs[1])
	assert.Equal(t, uint16(0x2), c.Regs[2])
	// 0b01
	// 0b10
	// 0b11 => 0x3
}

func TestSL(t *testing.T) {
	// opcode, regA, regB
	// 5, 1, 0
	c := setup([]uint16{0b0101_001_000_00000})
	c.Regs = cpu.Regs{0x3, 0x2, 0x2}
	c.Step()
	assert.Equal(t, uint16(0x5), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x4), c.Regs[1])
	assert.Equal(t, uint16(0x2), c.Regs[2])
	// 0b010
	// 0b100 => 0x4
}

func TestSR(t *testing.T) {
	// opcode, regA, regB
	// 6, 1, 0
	c := setup([]uint16{0b0110_001_000_00000})
	c.Regs = cpu.Regs{0x3, 0x2, 0x2}
	c.Step()
	assert.Equal(t, uint16(0x6), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x1), c.Regs[1])
	assert.Equal(t, uint16(0x2), c.Regs[2])
	// 0b010
	// 0b001 => 0x1
}

func TestSRA(t *testing.T) {
	// opcode, regA, regB
	// 7, 1, 0
	c := setup([]uint16{0b0111_001_000_00000})
	c.Regs = cpu.Regs{0x3, 0x2, 0x2}
	c.Step()
	assert.Equal(t, uint16(0x7), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x1), c.Regs[1])
	assert.Equal(t, uint16(0x2), c.Regs[2])
	// 0b010
	// 0b001 => 0x1
}

func TestLDL(t *testing.T) {
	// opcode, reg, val
	// 8, 0, 3
	// 8, 3, 4
	c := setup([]uint16{0b1000_000_000_00011, 0b1000_011_000_00100})
	c.Step()
	assert.Equal(t, uint16(0x8), c.IR)
	assert.Equal(t, uint16(0x03), c.Regs[0])
	assert.Equal(t, uint16(0x00), c.Regs[1])
	assert.Equal(t, uint16(0x00), c.Regs[2])
	assert.Equal(t, uint16(0x00), c.Regs[3])
	c.Step()
	assert.Equal(t, uint16(0x8), c.IR)
	assert.Equal(t, uint16(0x03), c.Regs[0])
	assert.Equal(t, uint16(0x00), c.Regs[1])
	assert.Equal(t, uint16(0x00), c.Regs[2])
	assert.Equal(t, uint16(0x04), c.Regs[3])
}

func TestLDH(t *testing.T) {
	// opcode, regA, val
	// 9, 1, 0
	c := setup([]uint16{0b1001_001_000_00001})
	c.Regs = cpu.Regs{0x3, 0x2, 0x2}
	c.Step()
	assert.Equal(t, uint16(0x9), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x102), c.Regs[1])
	assert.Equal(t, uint16(0x2), c.Regs[2])
	// 0b0000_0001 // operand 2
	// 0b0000_0000_0000_0010 // RegA
	// 0b0000_0001_0000_0010 // result
}

func TestCMP(t *testing.T) {
	// opcode, regA, regB
	// 10, 1, 0
	c := setup([]uint16{0b1010_001_001_00000, 0b1010_000_001_00000})
	c.Regs = cpu.Regs{0x3, 0x2}
	c.Step()
	assert.Equal(t, uint16(0xa), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x2), c.Regs[1])
	assert.Equal(t, true, c.EQFlag)
	c.Step()
	assert.Equal(t, uint16(0xa), c.IR)
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x2), c.Regs[1])
	assert.Equal(t, false, c.EQFlag)
}

func TestJE(t *testing.T) {
	// opcode, val
	// 11, 0
	c := setup([]uint16{0b1010_001_001_00000, 0b1011_000_000_00000})
	c.Regs = cpu.Regs{0x3, 0x2}
	c.Step()
	assert.Equal(t, uint16(0xa), c.IR)
	assert.Equal(t, true, c.EQFlag)

	c.Step()
	assert.Equal(t, uint16(0xb), c.IR)
	assert.Equal(t, uint16(0x0), c.PC)

	c.Step()
	assert.Equal(t, uint16(0xa), c.IR)
}

func TestJMP(t *testing.T) {
	// opcode, val
	// 12, 0
	c := setup([]uint16{0b1010_001_001_00000, 0b1100_000_000_00000})
	c.Regs = cpu.Regs{0x3, 0x2}
	c.Step()
	assert.Equal(t, uint16(0xa), c.IR)

	c.Step()
	assert.Equal(t, uint16(0xc), c.IR)
	assert.Equal(t, uint16(0x0), c.PC)
}

func TestLD(t *testing.T) {
	// opcode, regA, val
	// 13, 1, 2
	c := setup([]uint16{0b1101_001_000_00000})
	c.Regs = cpu.Regs{0x3, 0x2}
	c.RAM[0] = uint16(0xa)
	c.Step()
	assert.Equal(t, uint16(0xd), c.IR)
	assert.Equal(t, uint16(0xa), c.Regs[1])
}

func TestST(t *testing.T) {
	// opcode, regA, val
	// 14, 1, 2
	c := setup([]uint16{0b1110_001_000_00001})
	c.Regs = cpu.Regs{0x3, 0x4}
	c.Step()
	assert.Equal(t, uint16(0xe), c.IR)
	assert.Equal(t, uint16(0x4), c.RAM[1])
}

func TestHLT(t *testing.T) {
	// opcode
	// 15
	c := setup([]uint16{0b1111_000_000_00000})
	c.Step()
	assert.Equal(t, uint16(0xf), c.IR)
	assert.Equal(t, true, c.Halted)
}
