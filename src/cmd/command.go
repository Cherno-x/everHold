package cmd

import (
	"everHold/src/method"
	"everHold/src/tools"

	"github.com/spf13/cobra"
)

var methodValue string
var addValue bool
var nameValue string
var payloadValue string

func init() {
	RootCmd.AddCommand(PersistCmd)
	PersistCmd.PersistentFlags().StringVar(&methodValue, "method","1", "choose method to add StartUp")
	PersistCmd.PersistentFlags().StringVar(&nameValue, "name", "everHold", "name add to StartUp menu")
	PersistCmd.PersistentFlags().BoolVar(&addValue, "add", true, "add payload or delete payload option, default true")
	PersistCmd.PersistentFlags().StringVar(&payloadValue, "payload", "", "payload add to StartUp menu")
}

var RootCmd = &cobra.Command{
	Use:   "everHold",
	Short: "Windows Persistence Toolset",
}

var PersistCmd = &cobra.Command{
	Use:   "persist",
	Short: "Windows Persistence via StartUp",
	Run: func(cmd *cobra.Command, args []string) {
		if(methodValue=="1"){
			state := method.CallpersistStartup(payloadValue, addValue, nameValue)
			if state {
				tools.PrintSuccess(payloadValue + " add to StartUp Success")
			}
		}
		
	},
}
