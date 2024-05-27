package cmd

import (
	"everHold/src/method"
	"everHold/src/tools"

	"github.com/spf13/cobra"
)

var methodValue int
var addValue bool
var nameValue string

func init() {
	RootCmd.AddCommand(StartupCmd)
	PersistCmd.PersistentFlags().intVar(&methodValue, "method", int, "choose method to add StartUp")
	StartupCmd.PersistentFlags().BoolVar(&addValue, "add", true, "add payload or delete payload option, default true")
	StartupCmd.PersistentFlags().StringVar(&nameValue, "name", "", "name add to StartUp menu")
}

var RootCmd = &cobra.Command{
	Use:   "everHold",
	Short: "Windows Persistence Toolset",
}

var StartupCmd = &cobra.Command{
	Use:   "StartUp",
	Short: "add payload to StartUp menu",
	Run: func(cmd *cobra.Command, args []string) {
		payload := "C:\\Windows\\System32\\calc.exe"
		add := true
		name := "demo"
		state := method.CallpersistStartup(payload, add, name)
		if state {
			tools.PrintSuccess(payload + "add to StartUp Success")
		}
	},
}

var PersistCmd = &cobra.Command{
	Use:   "Persist",
	Short: "Windows Persistence via StartUp",
	Run: func(cmd *cobra.Command, args []string) {
		payload := "C:\\Windows\\System32\\calc.exe"
		add := true
		name := "demo"
		state := method.CallpersistStartup(payload, add, name)
		if state {
			tools.PrintSuccess(payload + "add to StartUp Success")
		}
	},
}
