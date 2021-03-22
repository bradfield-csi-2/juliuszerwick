package vm

import "fmt"

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
		//		fmt.Printf("Load: %v\n\n", Load)

		switch op {
		case Load:
			// Copy data value and store in register, overwriting previous register value.

			// Move PC forward and grab register to store data in.
			*pc += 1
			reg := memory[*pc]

			// Move PC forward and access data in memory location.
			*pc += 1
			data := memory[*pc]

			// Set the register value to data value.
			registers[reg] = data

			// Move PC forward to next instruction.
			*pc += 1

			fmt.Printf("registers after Load: %v\n\n", registers)
		case Store:
			// Copy data value from register and store in memory location, overwriting
			// previous value at the location.
			//
			// The memory location should always be 0x00.
			// If it is not, then we should end execution, possibly return an error value.
		case Add:
			// Set r1 = r1 + r2
		case Sub:
			// Set r1 = r1 - r2
		case Halt:
			// End execution.
			return
		}
	}

	return
}

// Compute calls the private compute function.
// Only used to indirectly call compute from main function
// for creating a feedback loop during development.
func Compute(memory []byte) {
	compute(memory)
}
