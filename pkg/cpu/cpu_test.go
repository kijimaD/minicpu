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

func TestAdd(t *testing.T) {
	// opcode, reg, val
	// 100000000000011
	cpu := setup([]uint16{0x4003})
	cpu.Step()
	assert.Equal(t, uint16(0x03), cpu.Regs[0], "should B equals ...")

	// 100001100000101
	cpu = setup([]uint16{0x4305})
	cpu.Step()
	assert.Equal(t, uint16(0x05), cpu.Regs[3], "should B equals ...")
}
