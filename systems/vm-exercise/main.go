package main

import (
	"fmt"

	"juliuszerwick/systems/vm-exercise/vm"
)

func main() {
	memory := make([]byte, 256)
	memory[1] = 0x01
	memory[8] = 0x01
	memory[9] = 0x01
	memory[10] = 0x01
	memory[11] = 0xff
	//	memory[12] = 0x02
	//	memory[13] = 0x01
	//	memory[14] = 0x00
	//	memory[15] = 0xff

	for i := 0; i <= 11; i++ {
		fmt.Printf("memory[%d]: %x\n", i, memory[i])
	}

	fmt.Println("\n")

	vm.Compute(memory)

	for i := 0; i <= 11; i++ {
		fmt.Printf("memory[%d]: %x\n", i, memory[i])
	}
}
