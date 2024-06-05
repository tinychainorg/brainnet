package brainnet

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func isValidBF(code string) bool {
	language := "><+-.,[]"
	for _, char := range code {
		if !strings.ContainsRune(language, char) {
			return false
		}
	}
	return true
}

func createJumpTable(chars string) map[int]int {
	jumpTable := make(map[int]int)
	leftPositions := []int{}
	for pos, char := range chars {
		if char == '[' {
			leftPositions = append(leftPositions, pos)
		} else if char == ']' {
			left := leftPositions[len(leftPositions)-1]
			leftPositions = leftPositions[:len(leftPositions)-1]
			jumpTable[left] = pos
			jumpTable[pos] = left
		}
	}
	return jumpTable
}

func dumpMem(mem map[int]int) {
	for key, value := range mem {
		fmt.Printf("%04d: %c\n", key, value)
	}
}

type OutOfGasException struct {
	message string
}

func (e *OutOfGasException) Error() string {
	return e.message
}

func brainfuckRun(code string, input string, mem map[int]int, gasLimit int, debugger bool) (string, string, map[int]int, int, error) {
	dp, ip, inputIdx, gasUsed := 0, 0, 0, 0
	output := ""
	jmpTable := createJumpTable(code)
	debugLog := ""

	for ip < len(code) {
		if gasUsed >= gasLimit {
			return "", "", mem, gasUsed, &OutOfGasException{fmt.Sprintf("gas_limit = %d, gas_used = %d", gasLimit, gasUsed)}
		}

		opcode := code[ip]
		debugLog += fmt.Sprintf("%05d %c\n", ip, opcode)

		switch opcode {
		case '>':
			dp++
		case '<':
			dp--
		case '+':
			mem[dp]++
			gasUsed += 3
		case '-':
			mem[dp]--
			gasUsed += 3
		case '.':
			output += string(mem[dp] % 256)
		case ',':
			if inputIdx < len(input) {
				mem[dp] = int(input[inputIdx])
				inputIdx++
			} else {
				mem[dp] = 0
			}
		case '[':
			if mem[dp] == 0 {
				ip = jmpTable[ip]
			}
		case ']':
			if mem[dp] != 0 {
				ip = jmpTable[ip]
			}
		case ';':
			for ip < len(code) && code[ip] != ';' {
				ip++
			}
		}

		ip++
		gasUsed++
	}

	return output, debugLog, mem, gasUsed, nil
}

type BrainfuckVM struct {
	memory map[int]int
}

func NewBrainfuckVM() *BrainfuckVM {
	return &BrainfuckVM{memory: make(map[int]int)}
}

func (vm *BrainfuckVM) Eval(code string, gasLimit int) (string, int, error) {
	memCopy := make(map[int]int)
	for k, v := range vm.memory {
		memCopy[k] = v
	}
	output, _, mem, gasUsed, err := brainfuckRun(code, "", memCopy, gasLimit, false)
	if err != nil {
		return "", 0, err
	}
	return output, gasUsed, nil
}

func (vm *BrainfuckVM) Apply(code string, gasLimit int) (string, int, error) {
	memCopy := make(map[int]int)
	for k, v := range vm.memory {
		memCopy[k] = v
	}
	output, _, mem, gasUsed, err := brainfuckRun(code, "", memCopy, gasLimit, false)
	if err != nil {
		return "", 0, err
	}
	vm.memory = memCopy
	return output, gasUsed, nil
}

func (vm *BrainfuckVM) DumpMemory() string {
	var result strings.Builder
	for key, value := range vm.memory {
		result.WriteString(fmt.Sprintf("%04d: %c\n", key, value))
	}
	return result.String()
}


