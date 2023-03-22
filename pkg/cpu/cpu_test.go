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
	c := setup([]uint16{0x0100, 0x0150})
	c.Regs = cpu.Regs{0x3, 0x5}
	c.Step()
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x3), c.Regs[1])
	c.Step()
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0x0), c.Regs[1])
}

func TestAdd(t *testing.T) {
	// opcode, regA, regB
	// 1, 1, 2
	c := setup([]uint16{0x1140})
	c.Regs = cpu.Regs{0x3, 0x5, 0x7}
	c.Step()
	assert.Equal(t, uint16(0x3), c.Regs[0])
	assert.Equal(t, uint16(0xc), c.Regs[1])
	assert.Equal(t, uint16(0x7), c.Regs[2])
}

func TestLdl(t *testing.T) {
	// opcode, reg, val
	// 8, 0, 3
	// 8, 3, 5
	c := setup([]uint16{0x4003, 0x4305})
	c.Step()
	assert.Equal(t, uint16(0x03), c.Regs[0])
	c.Step()
	assert.Equal(t, uint16(0x05), c.Regs[3])
}
