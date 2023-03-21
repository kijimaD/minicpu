package main

import (
	"fmt"

	"github.com/kijimaD/minicpu/pkg/cpu"
)

func main() {
	cpu := cpu.NewCPU()
	cpu.SetROM()
	fmt.Printf("%#v\n", cpu)
	cpu.Step()
	fmt.Printf("%#v\n", cpu)
}
