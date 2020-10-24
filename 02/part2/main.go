package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readinput()
	var breakloop bool
	mem := make([]int, len(input))
	var idx int

	for i := 1; i < 100; i++ {
		for j := 1; j < 100; j++ {
			copy(mem, input)
			mem[1] = i
			mem[2] = j
			for {
				if breakloop {
					breakloop = false
					break
				}
				switch mem[idx] {
				case 1:
					mem[mem[idx+3]] = mem[mem[idx+1]] + mem[mem[idx+2]]
					idx += 4
				case 2:
					mem[mem[idx+3]] = mem[mem[idx+1]] * mem[mem[idx+2]]
					idx += 4
				case 99:
					if mem[0] == 19690720 {
						fmt.Println(100*i + j)
						os.Exit(1)
					} else {
						idx = 0
						breakloop = true
						break
					}
				default:
					fmt.Printf("unknown code: %d", mem[idx])
				}
			}
		}
	}

}

func readinput() []int {
	f, _ := os.Open("../input")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	input := []int{}

	scanner.Scan()
	inputstr := scanner.Text()

	for _, s := range strings.Split(inputstr, ",") {
		l, _ := strconv.Atoi(s)
		input = append(input, l)
	}
	return input
}
