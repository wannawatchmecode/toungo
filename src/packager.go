package main

func PackageInput(input string) map[string][]int {
	m := make(map[string][]int)
	for i := 1; i < len(input); i++ {
		// insertPrefixes(m, input, i)
		insertWordsWithSize(m, input, i)
	}
	/*  	for index, value := range input {
		next := string(value)
		m[next] = append(m[next], index)
	}
	*/
	return m
}

func insertPrefixes(m map[string][]int, input string, wordSize int) {
	for i := 0; i < len(input)-wordSize; i++ {
		next := input[i : i+wordSize]
		m[next] = append(m[next], i)
	}
}

func insertWordsWithSize(m map[string][]int, input string, wordSize int) {
	buf := make([]byte, wordSize)
	bufIndex := 0
	i := 0
	for i < len(input) {
		if input[i] == byte(' ') {
			if bufIndex == wordSize {
				bufString := string(buf)
				m[bufString] = append(m[bufString], i-wordSize)
				bufIndex = 0
				i++
				continue
			} else {
				bufIndex = 0
			}
		} else if bufIndex >= wordSize {
			if input[i] == ' ' {
				bufIndex = 0
				i++
				continue
			}
			bufIndex++
		} else {
			buf[bufIndex] = input[i]
			bufIndex++
		}

		i++
	}
	if bufIndex == wordSize {
		bufString := string(buf)
		m[bufString] = append(m[bufString], i-wordSize)
	}
}
