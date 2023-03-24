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

func TestMov(t *testing.T) {
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

func TestAdd(t *testing.T) {
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

func TestSub(t *testing.T) {
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

func TestAnd(t *testing.T) {
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

func TestOr(t *testing.T) {
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

func TestSl(t *testing.T) {
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

func TestSr(t *testing.T) {
	// opcode, regA, regB
	// 5, 1, 0
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

func TestLdl(t *testing.T) {
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
