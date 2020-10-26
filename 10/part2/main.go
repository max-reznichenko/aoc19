package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	x int
	y int
}

func main() {
	Map := readinput()

	counter, cx, cy := 0, 0, 0
	for y := 0; y < len(Map); y++ {
		for x := 0; x < len(Map[0]); x++ {
			if string(Map[y][x]) == "#" {
				c := uniqAsteroids(Map, x, y)
				if counter == 0 || c > counter {
					counter = c
					cx = x
					cy = y
				}
			}
		}
	}

	dMap := buildDMap(Map, cx, cy)

	baseAngle := 0.0
	res := 0
	for i := 0; i < 200; i++ {
		baseAngle, res = vaporizeNext(&dMap, cx, cy, baseAngle)
	}

	fmt.Println(res)
}

func vaporizeNext(dMap *[][]float64, inX, inY int, inAngle float64) (float64, int) {
	// re-run this function from 0 angle if asteroid has not been detected at the end
	detected := false

	// smallest angle
	sAngle := 0.0

	// shortest distance
	sDistance := 0.0

	// coordinates of the detected asteroid
	resX, resY := 0, 0

	// to make sure we are not hitting the same angle next round
	angleShift := 0.0001

	for y, dMapRow := range *dMap {
		for x, dv := range dMapRow {
			if math.IsNaN(dv) { // nothing here
				continue
			}

			di := distance(inX, inY, x, y)

			if dv >= inAngle { // one of matching asteroids
				if !detected { // set base
					sDistance = di
					sAngle = dv
					resX, resY = x, y
					detected = true
					continue
				}

				// found asteroid at a better angle OR the same angle but closer
				if dv < sAngle || dv == sAngle && di < sDistance {
					sDistance = di
					sAngle = dv
					resX, resY = x, y
				}
			}
		}
	}

	if !detected { // nothing found. reset angle
		return vaporizeNext(dMap, inX, inY, 0.0)
	}

	(*dMap)[resY][resX] = math.NaN()
	return sAngle + angleShift, resX*100 + resY

}

// distance between base and detected asteroid
func distance(startX, startY, endX, endY int) float64 {
	lengthX := math.Abs(float64(endX - startX))
	lengthY := math.Abs(float64(endY - startY))

	return math.Sqrt(lengthX*lengthX + lengthY*lengthY)
}

func buildDMap(Map [][]byte, inX, inY int) [][]float64 {
	dMap := make([][]float64, len(Map))
	for y := 0; y < len(Map); y++ {
		dMapRow := make([]float64, len(Map[0]))
		for x := 0; x < len(Map[0]); x++ {
			if string(Map[y][x]) != "#" || x == inX && y == inY {
				dMapRow[x] = math.NaN()
				continue
			}
			dMapRow[x] = degrees(x-inX, y-inY)
		}
		dMap[y] = dMapRow
	}
	return dMap
}

func uniqAsteroids(Map [][]byte, inX, inY int) int {
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
	return math.Mod(math.Atan2(float64(y), float64(x))*(180/math.Pi)+360+90, 360)
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
