package main

import (
	"fmt"
	"sort"
)

func printInOrderOfMaxVals(m map[string][]int) {
	// Create a slice of key-value pairs
	type kv struct {
		Key   string
		Value []int
	}
	pairs := make([]kv, 0, len(m))

	for k, v := range m {
		pairs = append(pairs, kv{k, v})
	}

	// Sort slice by value
	sort.Slice(pairs, func(i, j int) bool {
		return len(pairs[i].Value) > len(pairs[j].Value) // Ascending order
	})

	// Print sorted key-value pairs
	for _, p := range pairs {
		fmt.Print(p.Key, " ")
		// fmt.Println(p.Key, ":", p.Value)
	}
}

func printInOrderOfMaxValsRemove(m map[string][]int) {
	// Create a slice of key-value pairs
	clonedMap := make(map[string][]int)
	for k, v := range m {
		clonedMap[k] = v
	}
	type kv struct {
		Key   string
		Value []int
	}
	pairs := make([]kv, 0, len(m))

	for len(clonedMap) > 0 {
		for k, v := range clonedMap {
			pairs = append(pairs, kv{k, v})
		}

		// Sort slice by value
		sort.Slice(pairs, func(i, j int) bool {
			return len(pairs[i].Value) > len(pairs[j].Value) // Ascending order
		})

		fmt.Print(pairs[0].Key, " ")
		clonedMap[pairs[0].Key] = clonedMap[pairs[0].Key][1:]
		if len(clonedMap[pairs[0].Key]) < 1 {
			clonedMap[pairs[0].Key] = nil
		}
		// Print sorted key-value pairs
		//	for _, p := range pairs {
		fmt.Print(pairs[0].Key, " ")
		// fmt.Println(p.Key, ":", p.Value)
		//  }
	}
}

/*
func findAllWords(wordToFind string, indexMap map[string][]int) {
	firstChar := string(wordToFind[0])
	indexes := indexMap[firstChar]
	output := [0]string
	for c := range wordToFind {
		for index := range indexes {
			//			joinIndexes()
		}
	}
}*/
// Break time....
// Let's turn this into a web service
// func joinIndexes(left: )
func main() {
	RunService("3006")

	/*fmt.Println("Hello, Go! Not so bad for my first go project!!!")
	fileName := "./input.txt"

	input_string := "Hello world"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Unable to open file, using default string")
	} else {
		input_string = string(data)
	}
	m := PackageInput(input_string)

	// fmt.Println(input_string)
	// fmt.Println(m)
	s := GenerateInOrderOfMaxValsRemove(m)
	//	printInOrderOfMaxVals(m)
	fmt.Println(s)*/
}
