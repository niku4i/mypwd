package main

import (
	"code.google.com/p/gopass"
	"fmt"
	"os"
)

type Scanner struct {
	Ask        string
	Validation bool
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (s *Scanner) Input() string {
	passwd, err := gopass.GetPass(s.Ask)
	handleErr(err)
	if s.Validation {
		passwd_c, err := gopass.GetPass("Confirm: ")
		handleErr(err)
		if passwd != passwd_c {
			os.Exit(1)
		}
	}
	return passwd
}
