package main

import (
	"fmt"
	"os"
)

const SIZE int = 32768

func main() {

	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while opening file")
	}

	result, err := readData(data)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(result)
	}
}

func readData(data []byte) (result string, err error) {
	text := string(data)

	stack := make([]int, SIZE)

	var idx, ptr int = 0, 0
	var str string
	var resultData string

	for idx < len(text) {

		switch string(text[idx]) {
		case ".":
			resultData += fmt.Sprintf("%c", stack[ptr])
		case ",":
			fmt.Scanln(&str)
			for _, c := range str {
				stack[idx] = int(c)
			}

		case ">":
			ptr = (ptr + SIZE + 1) % SIZE
		case "<":
			ptr = (ptr + SIZE - 1) % SIZE
		case "+":
			stack[ptr] = (stack[ptr] + 255 + 1) % 255
		case "-":
			stack[ptr] = (stack[ptr] + 255 - 1) % 255
		case "[":
			if stack[ptr] == 0 {
				for string(text[idx]) != "]" {
					idx += 1
				}
			}
		case "]":
			if stack[ptr] != 0 {
				for string(text[idx]) != "[" {
					idx -= 1
				}
			}

		}
		idx++
	}
	return resultData, nil
}
