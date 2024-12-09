package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		panic("not enough command line args")
	}

	bytes, _ := ioutil.ReadFile(os.Args[1])
	str := string(bytes)
	words := strings.Fields(str)
	wordcount := make(map[string]int)
	for _, word := range words {
		wordcount[word] = wordcount[word] + 1
	}
	result := ""
	for key, value := range wordcount {
		result = result + key + " " + strconv.Itoa(value) + "\n"
	}

	ioutil.WriteFile("result.txt", []byte(result), 644)
}
