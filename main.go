package main

import (
	"everHold/src/cmd"
	"everHold/src/conf"
)

func init() {
	conf.Banner()
}

func main() {
	cmd.RootCmd.Execute()
}
