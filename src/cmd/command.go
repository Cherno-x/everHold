package cmd

import (
	"everHold/src/conf"
	"everHold/src/func/persist"
	"everHold/src/utils"

	"github.com/spf13/cobra"
)

var newcmd conf.RUNCMD

func init() {

	RootCmd.AddCommand(PersistCmd)
	RootCmd.AddCommand(InjectCmd)
	PersistCmd.PersistentFlags().StringVar(&newcmd.MethodValue, "method", "1", "choose method")
	PersistCmd.PersistentFlags().StringVar(&newcmd.NameValue, "name", "everHold", "name add to StartUp menu")
	PersistCmd.PersistentFlags().BoolVar(&newcmd.AddValue, "add", true, "add payload or delete payload option, default true")
	PersistCmd.PersistentFlags().StringVar(&newcmd.PayloadValue, "payload", "", "payload add to StartUp menu")
}

var RootCmd = &cobra.Command{
	Use:   "everHold",
	Short: "Windows Persistence Toolset",
}

var PersistCmd = &cobra.Command{
	Use:   "persist",
	Short: "Windows Persistence",
	Run: func(cmd *cobra.Command, args []string) {
		newcmd.ModuleValue = "persist"
		if newcmd.MethodValue == "1" {
			state := persist.Callmethod1(&newcmd)
			if state {
				utils.PrintSuccess(newcmd.PayloadValue + " add to StartUp Success")
			}

		} else if newcmd.MethodValue == "2" {
			persist.Callmethod2(&newcmd)
		}

	},
}

var InjectCmd = &cobra.Command{
	Use:   "inject",
	Short: "Windows Process Injection",
	Run: func(cmd *cobra.Command, args []string) {
		newcmd.ModuleValue = "inject"
		if newcmd.MethodValue == "1" {
			state := persist.Callmethod1(&newcmd)
			if state {
				utils.PrintSuccess(newcmd.PayloadValue + " add to StartUp Success")
			}
		}

	},
}
