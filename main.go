package main

// Interpreter is from https://github.com/samuel-pratt/brainfuck-go

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
)

//go:embed program.bf
var Program string

func main() {
	var arr [30000]byte
	var curr int = 0

	file := Program

	for i := 0; i < len(file); i++ {
		switch file[i] {
		// +
		case 43:
			arr[curr] = arr[curr] + 1
		// -
		case 45:
			arr[curr] = arr[curr] - 1
		// ,
		case 44:
			reader := bufio.NewReader(os.Stdin)
			char, err := reader.ReadByte()
			if err != nil {
				fmt.Print(err)
				return
			}
			arr[curr] = char

		// .
		case 46:
			fmt.Print(string(arr[curr]))

		// [
		case 91:
			if arr[curr] == 0 {
				var skips int = 1
				for skips != 0 {
					i = i + 1
					if file[i] == 91 {
						skips = skips + 1
					}
					if file[i] == 93 {
						skips = skips - 1
					}
				}
			}

		// ]
		case 93:
			if arr[curr] != 0 {
				var skips int = 1
				for skips != 0 {
					i = i - 1
					if file[i] == 93 {
						skips = skips + 1
					}
					if file[i] == 91 {
						skips = skips - 1
					}
				}
			}

		// <
		case 60:
			curr = curr - 1

		// >
		case 62:
			curr = curr + 1

		}
	}
}
