package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var weightedTree = make(map[string]int)

func main() {
	buildTree(readinput(), []string{}, "COM")

	counter := 0
	for _, v := range weightedTree {
		counter += v
	}
	fmt.Println(counter)
}

func buildTree(base []string, trace []string, node string) {
	trace = append(trace, node)

	for _, s := range base {
		if ok, _ := regexp.MatchString(fmt.Sprintf(`%s\)`, node), s); ok {
			d := strings.Split(s, ")")[1]
			for _, n := range trace {
				weightedTree[n]++
			}
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
