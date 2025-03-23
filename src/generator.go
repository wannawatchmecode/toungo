package main

import (
	"sort"
	"strings"
)

func GenerateInOrderOfMaxVals(m map[string][]int) string {
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

	var sb strings.Builder
	// Print sorted key-value pairs
	for _, p := range pairs {
		sb.WriteString(p.Key)
		sb.WriteString(" ")
		//		fmt.Print(p.Key, " ")
		// fmt.Println(p.Key, ":", p.Value)
	}
	return sb.String()
}

func GenerateInOrderOfMaxValsRemove(m map[string][]int) string {
	// Create a slice of key-value pairs
	clonedMap := make(map[string][]int)
	for k, v := range m {
		clonedMap[k] = v
	}
	type kv struct {
		Key   string
		Value []int
	}
	var sb strings.Builder

	for len(clonedMap) > 0 {

		pairs := make([]kv, 0)
		for k, v := range clonedMap {
			if len(v) < 1 {
				continue
			}
			pairs = append(pairs, kv{k, v})
		}

		if len(pairs) == 0 {
			break
		}

		// Sort slice by value
		sort.Slice(pairs, func(i, j int) bool {
			return len(pairs[i].Value) > len(pairs[j].Value) // Ascending order
		})

		//		clonedMap[pairs[0].Key] = clonedMap[pairs[0].Key][1:]
		// fmt.Print(pairs[0].Key, " ")

		clonedMap[pairs[0].Key] = clonedMap[pairs[0].Key][1:]
		if len(clonedMap[pairs[0].Key]) < 1 {
			//			clonedMap[pairs[0].Key] = nil
		} else {
			//			clonedMap[pairs[0].Key] = clonedMap[pairs[0].Key][1:]
		}
		// Print sorted key-value pairs
		//	for _, p := range pairs {
		//		fmt.Print(pairs[0].Key, " ")
		sb.WriteString(pairs[0].Key)
		sb.WriteString(" ")
		// fmt.Println(p.Key, ":", p.Value)
		//  }
	}
	return sb.String()
}
