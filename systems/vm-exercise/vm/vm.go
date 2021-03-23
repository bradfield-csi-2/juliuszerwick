package vm

import (
	"fmt"
	_ "time"
)

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
//
func compute(memory []byte) {

	// PC register is set to value 8 so that it starts at
	// the first byte representing the first instruction
	// in the memory slice.
	//	fmt.Println("Memory before:\n")
	for i := 0; i < 3; i++ {
		fmt.Printf("memory[%d]: %x\n", i, memory[i])
	}
	registers := [3]byte{8, 0, 0} // PC, R1 and R2
	pc := &registers[0]
	//r1 := &registers[1]
	//r2 := &registers[2]

	// Debug lines.
	fmt.Printf("Inside compute func\n")
	fmt.Printf("registers: %v\n\n", registers)

	// Keep looping, like a physical computer's clock
	// Stretch: If memory never contains a value 0xff, then we'll loop back around
	//          to the beginning of the memory array. We need to also stop execution
	//          if we reach the end of the memory array and possibly return an error.
	for {
		// The opcode represents the instruction to be executed.
		op := memory[*pc]
		fmt.Printf("op: %v\n\n", op)
		//time.Sleep(3 * time.Second)

		switch op {
		case Load:
			// Copy data value and store in register, overwriting previous register value.
			// Move PC forward and grab register to store data in.
			*pc += 1
			fmt.Printf("*pc = %v\n", *pc)
			reg := memory[*pc]
			fmt.Printf("reg: %v\n", reg)
			// Move PC forward and access data in memory location.
			*pc += 1
			fmt.Printf("*pc = %v\n", *pc)
			dataLocation := memory[*pc]
			data := memory[dataLocation]
			fmt.Printf("data to be loaded: %v\n", data)
			// Set the register value to data value.
			registers[reg] = data
			// Move PC forward to next instruction.
			*pc += 1
			fmt.Printf("*pc = %v\n", *pc)
			fmt.Printf("registers after Load: %v\n\n", registers)
		case Store:
			// Copy data value from register and store in memory location, overwriting
			// previous value at the location.
			//
			// The memory location should always be 0x00.
			// If it is not, then we should end execution, possibly return an error value.
			*pc += 1
			fmt.Printf("*pc = %v\n", *pc)
			reg := memory[*pc]
			fmt.Printf("reg: %v\n", reg)
			// Move PC forward and store data from register in memory location.
			*pc += 1
			fmt.Printf("*pc = %v\n", *pc)

			location := memory[*pc]
			fmt.Printf("location: %v\n", location)
			if location != 0x00 {
				fmt.Printf("Location is not the first byte in memory\n\n")
				return
			}

			memory[location] = registers[reg]
			fmt.Printf("first byte: %v\n", memory[location])

			*pc += 1
			fmt.Printf("*pc = %v\n", *pc)
			fmt.Printf("registers after Store: %v\n\n", registers)
		case Add:
			// Set r1 = r1 + r2
			*pc += 1
			a1 := memory[*pc]

			*pc += 1
			a2 := memory[*pc]

			registers[a1] = registers[a1] + registers[a2]

			*pc += 1
			fmt.Printf("registers after Add: %v\n\n", registers)
		case Sub:
			// Set r1 = r1 - r2
			*pc += 1
			a1 := memory[*pc]

			*pc += 1
			a2 := memory[*pc]

			registers[a1] = registers[a1] - registers[a2]

			*pc += 1
			fmt.Printf("registers after Sub: %v\n\n", registers)
		case Halt:
			// End execution.
			//fmt.Println("Memory after:\n")
			//for i := 0; i < 4; i++ {
			//	fmt.Printf("memory[%d]: %x\n", i, memory[i])
			//}
			//fmt.Println("\n")
			return
		}
	}
}

// Compute calls the private compute function.
// Only used to indirectly call compute from main function
// for creating a feedback loop during development.
func Compute(memory []byte) {
	compute(memory)
}
