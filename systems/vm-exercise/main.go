package main

import (
	"fmt"

	"juliuszerwick/systems/vm-exercise/vm"
)

func main() {
	memory := make([]byte, 256)
	memory[1] = 0x01
	memory[1] = 0x03
	memory[8] = 0x01
	memory[9] = 0x01
	memory[10] = 0x01
	memory[11] = 0x01
	memory[12] = 0x02
	memory[13] = 0x02
	memory[14] = 0x02
	memory[15] = 0x01
	memory[16] = 0x00
	memory[17] = 0x03
	memory[18] = 0x01
	memory[19] = 0x02
	memory[20] = 0x04
	memory[21] = 0x01
	memory[22] = 0x02
	memory[23] = 0xff

	for i := 0; i <= 20; i++ {
		fmt.Printf("memory[%d]: %x\n", i, memory[i])
	}

	vm.Compute(memory)

	for i := 0; i <= 20; i++ {
		fmt.Printf("memory[%d]: %x\n", i, memory[i])
	}
}
