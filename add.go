package main

import (
//"fmt"
)

var cmdNew = &Command{
	Run:       runAdd,
	UsageLine: "add [-u user] [-p passwd] [-r] [service]",
	Short:     "add add service",
	Long:      "add add service",
}

var addU = cmdNew.Flag.String("u", "", "username")

func runAdd(cmd *Command, args []string) {
}
