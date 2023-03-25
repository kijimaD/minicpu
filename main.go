package main

import (
	"fmt"

	"github.com/kijimaD/minicpu/pkg/cpu"
)

func main() {
	c := cpu.NewCPU()
	c.SetROM()

	for i, _ := range c.ROM {
		if c.Halted == true {
			break
		}
		c.Step()
		inst := cpu.Instructions[c.Curop]
		fmt.Printf("%2d\t%s\t%5d\t%5x\t%5d\t%5d\t%d\t%d\n", i, inst.Description, c.PC, c.IR, c.Regs[0], c.Regs[1], c.Regs[2], c.Regs[3])
	}
	fmt.Printf("ram[64] = %d\n", c.RAM[64])
	fmt.Printf("i\tOP\tPC\tIR\tReg0\tReg1\tReg2\tReg3\n")
}
