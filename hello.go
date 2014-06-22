package main

import (
	"fmt"
)

var cmdHello = &Command{
	Run:       runHello,
	UsageLine: "hello [name]",
	Short:     "print hello",
	Long: `
print hello to you!
`,
}

func runHello(cmd *Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}

	fmt.Println("Hello")
}
