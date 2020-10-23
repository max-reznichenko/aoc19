package main

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readinput()

	inX, inY := 25, 6

	layers := [][]int{}
	layer := []int{}
	for i, d := range input {
		if i%(inX*inY) == 0 && i != 0 {
			layers = append(layers, layer)
			layer = []int{}
		}

		layer = append(layer, d)

		if i+1 == len(input) {
			layers = append(layers, layer)
		}
	}

	mergedLayer := make([]int, inX*inY)
	for i := 0; i < inX*inY; i++ {
		mergedLayer[i] = 2
		for _, layer := range layers {
			if layer[i] != 2 && mergedLayer[i] == 2 {
				mergedLayer[i] = layer[i]
			}
		}
	}

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{inX, inY}})

	for y := 0; y < inY; y++ {
		for x := 0; x < inX; x++ {
			switch mergedLayer[y*inX+x] {
			case 1:
				img.Set(x, y, color.White)
			case 0:
				img.Set(x, y, color.Black)
			}
		}
	}

	f, _ := os.Create("pass.png")
	png.Encode(f, img)
}

func readinput() []int {
	f, _ := os.Open("../input")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	input := []int{}

	scanner.Scan()
	inputstr := scanner.Text()

	for _, s := range strings.Split(inputstr, "") {
		l, _ := strconv.Atoi(s)
		input = append(input, l)
	}
	return input
}
