package brainnet

import "testing"

func TestInterpreter() {
	vm := NewBrainfuckVM()
	code, err := ioutil.ReadFile("brainfuck/solong.bf")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	output, gasCost, err := vm.Apply(string(code), 250000000000)
	if err != nil {
		fmt.Println("Error executing Brainfuck code:", err)
		return
	}

	fmt.Printf("Output: %s\nGas Cost: %d\nMemory Dump:\n%s", output, gasCost, vm.DumpMemory())
}