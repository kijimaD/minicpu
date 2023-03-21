package cpu_test

import (
	"testing"

	"github.com/kijimaD/minicpu/pkg/cpu"
	"github.com/stretchr/testify/assert"
)

func setup(data []uint16) *cpu.CPU {
	cpu := &cpu.CPU{
		PC:   0,
		Regs: cpu.Regs{0},
		ROM:  [256]uint16{0},
		RAM:  [256]uint16{0},
	}

	for i, d := range data {
		cpu.ROM[i] = d
	}

	return cpu
}

func TestAdd(t *testing.T) {
	// opcode, reg, val
	cpu := setup([]uint16{0x08, 0x01, 0x02})
	cpu.Step()
	assert.Equal(t, uint16(0x02), cpu.Regs[1], "should B equals ...")
}
