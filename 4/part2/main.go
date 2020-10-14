package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s, e := readinput()
	c := 0

	for ; s <= e; s++ {
		bb := []byte(strconv.Itoa(s))

		ok := true

		// inc/dec check
		for i := range bb {
			if i+1 == len(bb) {
				continue
			}
			if bytes.Compare([]byte{bb[i]}, []byte{bb[i+1]}) > 0 {
				ok = false
			}
		}
		if !ok {
			continue
		}

		ok = false

		var counts []int
		for _, b := range bb {
			counts = append(counts, bytes.Count(bb, []byte{b}))
		}

		for _, i := range counts {
			if i == 2 {
				ok = true
			}
		}

		if !ok {
			continue
		}
		c++
	}

	fmt.Println(c)
}

func readinput() (int, int) {
	f, _ := os.Open("../input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	i := strings.Split(scanner.Text(), "-")

	v1, _ := strconv.Atoi(i[0])
	v2, _ := strconv.Atoi(i[1])

	return v1, v2
}
