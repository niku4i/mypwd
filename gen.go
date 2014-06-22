package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var lettersLower = "abcdefghijklmnopqrstuvwxyz"
var lettersUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lettersNumber = "0123456789"
var lettersCode = "-/:;()@.,?!'[]{}%^*+=_~<>"
var letters = lettersLower + lettersLower + lettersUpper + lettersUpper + lettersNumber + lettersCode

var cmdGen = &Command{
	Run:       runGen,
	UsageLine: "gen",
	Short:     "generate random password string",
	Long:      "generate random password string",
}

func runGen(cmd *Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randString())
}

func randString() string {
	l := 16
	a := make([]string, l)
	for i := 0; i < l; i++ {
		a[i] = fmt.Sprintf("%c", letters[rand.Intn(len(letters))])
	}
	return strings.Join(a, "")
}
