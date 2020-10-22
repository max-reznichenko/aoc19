package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type amp struct {
	id         int
	iset       []int
	pointer    int
	inChan     chan int
	inBaseChan chan int
	outChan    chan int
	lastOut    int
	wg         *sync.WaitGroup
}

func newAmp(id int, iset []int, wg *sync.WaitGroup) *amp {
	a := amp{
		id:         id,
		inChan:     make(chan int, 1),
		inBaseChan: make(chan int, 1),
		outChan:    make(chan int, 1),
		wg:         wg,
	}
	a.iset = make([]int, len(iset))
	copy(a.iset, iset)

	return &a
}

func (a *amp) nextInstruction() (op, v1, v2, v3 int) {
	ins := fmt.Sprintf("%05d", a.iset[a.pointer])

	op, _ = strconv.Atoi(string(ins[3:]))
	pmode1, _ := strconv.Atoi(string(ins[2]))
	pmode2, _ := strconv.Atoi(string(ins[1]))
	_, _ = strconv.Atoi(string(ins[0])) // not used

	switch op {
	case 1, 2, 5, 6, 7, 8:
		if pmode1 == 1 {
			v1 = a.iset[a.pointer+1]
		} else {
			v1 = a.iset[a.iset[a.pointer+1]]
		}

		if pmode2 == 1 {
			v2 = a.iset[a.pointer+2]
		} else {
			v2 = a.iset[a.iset[a.pointer+2]]
		}
		return op, v1, v2, a.iset[a.pointer+3]
	case 3:
		return op, a.iset[a.pointer+1], 0, 0
	case 4:
		if pmode1 == 1 {
			v1 = a.iset[a.pointer+1]
		} else {
			v1 = a.iset[a.iset[a.pointer+1]]
		}
		return op, v1, 0, 0
	case 99:
		return op, 0, 0, 0
	default:
		return 0, 0, 0, 0
	}
}

func (a *amp) compute() {
	for {
		op, v1, v2, v3 := a.nextInstruction()
		switch op {
		case 1:
			a.iset[v3] = v1 + v2
			a.pointer += 4
		case 2:
			a.iset[v3] = v1 * v2
			a.pointer += 4
		case 3:
			select {
			case msg := <-a.inBaseChan:
				a.iset[v1] = msg
				a.pointer += 2
				a.wg.Done()
			case msg := <-a.inChan:
				a.iset[v1] = msg
				a.pointer += 2
			}
		case 4:
			a.outChan <- v1
			a.lastOut = v1
			a.pointer += 2
		case 5:
			if v1 > 0 {
				a.pointer = v2
			} else {
				a.pointer += 3
			}
		case 6:
			if v1 == 0 {
				a.pointer = v2
			} else {
				a.pointer += 3
			}
		case 7:
			if v1 < v2 {
				a.iset[v3] = 1
			} else {
				a.iset[v3] = 0
			}
			a.pointer += 4
		case 8:
			if v1 == v2 {
				a.iset[v3] = 1
			} else {
				a.iset[v3] = 0
			}
			a.pointer += 4
		case 99:
			a.wg.Done()
			return
		default:
			fmt.Printf("got invalid code: %d\n", op)
			os.Exit(1)
		}
	}
}

func main() {
	code := readinput()

	thrust := 0
	for _, perm := range permutation([]int{5, 6, 7, 8, 9}) {
		var wg sync.WaitGroup
		wg.Add(5)

		// init amplifiers and set base value
		var amplifiers [5]*amp
		for ampID, phaseSetting := range perm {
			amplifiers[ampID] = newAmp(ampID, code, &wg)
			amplifiers[ampID].inBaseChan <- phaseSetting
			go amplifiers[ampID].compute()
		}

		// wire amplifiers
		for id := range amplifiers {
			if id+1 == len(amplifiers) {
				amplifiers[0].inChan = amplifiers[id].outChan
			} else {
				amplifiers[id+1].inChan = amplifiers[id].outChan
			}
		}
		wg.Wait()

		wg.Add(5)
		for _, amp := range amplifiers {
			amp.wg = &wg
		}

		amplifiers[0].inChan <- 0
		wg.Wait()

		if amplifiers[4].lastOut > thrust {
			thrust = amplifiers[4].lastOut
		}

	}
	fmt.Printf("max thrust: %d\n", thrust)
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

// copied from https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}
