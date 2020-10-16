package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// var weightedTree = make(map[string]int)
var youTrace []string
var sanTrace []string

var sanTraceD bool

func main() {
	buildTree(readinput(), []string{}, "COM")

	i := len(sanTrace)
	if len(youTrace) > len(sanTrace) {
		i = len(youTrace)
	}

	c := 0
	for j := 0; j < i; j++ {
		if j < len(youTrace) && j < len(sanTrace) {
			if youTrace[j] == sanTrace[j] && youTrace[j] != "YOU" && sanTrace[j] != "SAN" {
				c++
			} else {
				break
			}
		}
	}

	fmt.Println(len(youTrace) - c - 1 + len(sanTrace) - c - 1)
}

func buildTree(base []string, trace []string, node string) {
	trace = append(trace, node)

	switch node {
	case "SAN":
		sanTrace = make([]string, len(trace))
		copy(sanTrace, trace)
	case "YOU":
		youTrace = make([]string, len(trace))
		copy(youTrace, trace)
	}

	for _, s := range base {
		if ok, _ := regexp.MatchString(fmt.Sprintf(`%s\)`, node), s); ok {
			d := strings.Split(s, ")")[1]
			buildTree(base, trace, d)
		}

	}
}

func readinput() []string {
	f, _ := os.Open("../input")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var i []string
	for scanner.Scan() {
		i = append(i, scanner.Text())
	}
	return i
}
