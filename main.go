package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData(path string) [][]int {
	dat, err := ioutil.ReadFile(path)
	check(err)

	var matrix [][]int
	json.Unmarshal([]byte(dat), &matrix)
	return matrix
}

func main() {

	for _, row := range readData("data.json") {
		for _, val := range row {
			fmt.Printf("%d, ", val)
		}
		fmt.Printf("\n")
	}

	println("it's over 9000!")
}
