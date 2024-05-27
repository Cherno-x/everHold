package tools

import (
	"fmt"
)

func PrintError(message string) {
	fmt.Println("[ERROR]", message)
}

func PrintSuccess(message string) {
	fmt.Println("[SUCCESS]", message)
}

func PrintInfo(message string) {
	fmt.Println("[INFO]", message)
}
