package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	Map := readinput()

	counter := 0
	for y := 0; y < len(Map); y++ {
		for x := 0; x < len(Map[0]); x++ {
			if string(Map[y][x]) == "#" {
				c := visibleAsteroids(Map, x, y)
				if counter == 0 || c > counter {
					counter = c
				}
			}
		}
	}

	fmt.Printf("max = %d\n", counter)
}

func visibleAsteroids(Map [][]byte, inX, inY int) int {
	dMap := []float64{}
	for y := 0; y < len(Map); y++ {
		for x := 0; x < len(Map[0]); x++ {
			if string(Map[y][x]) != "#" || x == inX && y == inY {
				continue
			}
			dMap = append(dMap, degrees(x-inX, y-inY))
		}
	}
	return len(unique(dMap))
}

func degrees(x, y int) float64 {
	return math.Atan2(float64(y), float64(x)) * (180 / math.Pi)
}

func readinput() [][]byte {
	f, _ := os.Open("../input")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	input := [][]byte{}

	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
	}
	return input
}

// copied from https://www.golangprograms.com/remove-duplicate-values-from-slice.html
func unique(intSlice []float64) []float64 {
	keys := make(map[float64]bool)
	list := []float64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
