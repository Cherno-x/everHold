package utils

import (
	"os"
	"path/filepath"
)

func CheckExe(payload []string) (bool, string) {
	if len(payload) == 0 {
		return false, ""
	}

	exePath := payload[0]
	if fileInfo, err := os.Stat(exePath); err == nil && !fileInfo.IsDir() && filepath.Ext(exePath) == ".exe" {
		commandline := ""
		for index, obj := range payload {
			if index == len(payload)-1 {
				commandline += obj
			} else {
				commandline += obj + " "
			}
		}
		return true, commandline
	}
	return false, ""
}

