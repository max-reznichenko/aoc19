package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type computer struct {
	id           int
	iset         []int
	pointer      int
	relativeBase int
	inChan       chan int
	wg           *sync.WaitGroup
}

func newComputer(id int, iset []int, wg *sync.WaitGroup) *computer {
	c := computer{
		id:     id,
		inChan: make(chan int, 1),
		wg:     wg,
	}
	c.iset = make([]int, len(iset)+100)
	copy(c.iset, iset)

	return &c
}

func (c *computer) compute() {
	for {
		op, mode1, mode2, mode3 := c.nextInstruction()
		switch op {
		case 1: // addition
			*c.getMemoryValue(mode3, 3) = *c.getMemoryValue(mode1, 1) + *c.getMemoryValue(mode2, 2)
			c.pointer += 4
		case 2: // multiplication
			*c.getMemoryValue(mode3, 3) = *c.getMemoryValue(mode1, 1) * *c.getMemoryValue(mode2, 2)
			c.pointer += 4
		case 3: // input
			select {
			case msg := <-c.inChan:
				*c.getMemoryValue(mode1, 1) = msg
				c.pointer += 2
			}
		case 4: // output
			fmt.Println(*c.getMemoryValue(mode1, 1))
			c.pointer += 2
		case 5: // jump if non zero
			if *c.getMemoryValue(mode1, 1) != 0 {
				c.pointer = *c.getMemoryValue(mode2, 2)
			} else {
				c.pointer += 3
			}
		case 6: // jump if zero
			if *c.getMemoryValue(mode1, 1) == 0 {
				c.pointer = *c.getMemoryValue(mode2, 2)
			} else {
				c.pointer += 3
			}
		case 7: // less than
			if *c.getMemoryValue(mode1, 1) < *c.getMemoryValue(mode2, 2) {
				*c.getMemoryValue(mode3, 3) = 1
			} else {
				*c.getMemoryValue(mode3, 3) = 0
			}
			c.pointer += 4
		case 8: // equals
			if *c.getMemoryValue(mode1, 1) == *c.getMemoryValue(mode2, 2) {
				*c.getMemoryValue(mode3, 3) = 1
			} else {
				*c.getMemoryValue(mode3, 3) = 0
			}
			c.pointer += 4
		case 9:
			c.relativeBase += *c.getMemoryValue(mode1, 1)
			c.pointer += 2
		case 99:
			c.wg.Done()
			return
		default:
			fmt.Printf("got invalid code: %d\n", op)
			os.Exit(1)
		}
	}
}

// parses instruction to op + modes
func (c *computer) nextInstruction() (int, int, int, int) {
	ins := fmt.Sprintf("%05d", c.iset[c.pointer])

	op, _ := strconv.Atoi(string(ins[3:]))
	mode1, _ := strconv.Atoi(string(ins[2]))
	mode2, _ := strconv.Atoi(string(ins[1]))
	mode3, _ := strconv.Atoi(string(ins[0]))

	return op, mode1, mode2, mode3
}

// returns a pointer to the iset value based on mode
func (c *computer) getMemoryValue(mode, offset int) *int {
	switch mode {
	case 0:
		return &c.iset[c.iset[c.pointer+offset]]
	case 1:
		return &c.iset[c.pointer+offset]
	case 2:
		return &c.iset[c.iset[c.pointer+offset]+c.relativeBase]
	default:
		panic(fmt.Sprintf("invalid mode %d\n", mode))
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	c := newComputer(1, readinput(), &wg)
	c.inChan <- 1
	c.compute()
	wg.Wait()
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
