package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type State struct {
	Name string
	Size int
}

var states []State

var average int

func (s State) DisplayState() string {

	temp := "less than"
	if s.Size > average {
		temp = "greater than"
	}

	result := fmt.Sprintf("%s is %s average (%d).\n",
		s.Name,
		temp,
		s.Size)

	return result
}
func main() {

	data, _ := ioutil.ReadFile("states.txt")
	entries := strings.Split(string(data), "\n")
	total := 0
	for _, value := range entries {
		index := strings.LastIndex(value, " ")
		state := value[0:index]
		size := value[index+1:]
		fmt.Println(state, size)
		isize, _ := strconv.Atoi(strings.TrimSpace(
			size))
		states = append(states, State{state, isize})
		total = total + isize
	}

	average = total / len(entries)
	
	t := template.New("")

	t.Parse("{{range .}} {{.DisplayState}} {{end}}")

	t.Execute(os.Stdout, states)
}
