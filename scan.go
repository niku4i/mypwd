package main

import (
	"bufio"
	"os"
)

type Scanner struct {
	Ask        string
	Validation bool
}

func (s *Scanner) Input() string {
	print(s.Ask)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if s.Validation {
		print("Confirm: ")
		input_c, _ := reader.ReadString('\n')
		if input != input_c {
			os.Exit(2)
		}
	}
	return input
}
