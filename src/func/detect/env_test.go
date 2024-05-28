package detect

import (
	"everHold/src/utils"
	"fmt"
	"testing"
)

func TestWindows(t *testing.T) {
	result, error := GetWindowVersion()
	if error != nil {
		t.Fatal(error)
	}
	fmt.Printf("Windows 版本: %d.%d (Build %d)\n",
		result.dwMajorVersion, result.dwMinorVersion, result.dwBuildNumber)
}

func TestAV(t *testing.T) {
	// 测试Add函数
	results, error := CheckAntivirusProcesses()
	if error != nil {
		t.Fatal(error)
	}
	for _, result := range results {
		fmt.Printf("process:%s , AV:%s\n", result.ProcessName, result.DisplayName)
	}
}

func TestHejing(t *testing.T) {
	// 测试Add函数
	hejingStatus := Check360hejing()
	if hejingStatus == true {
		utils.PrintSuccess("检测到核晶开启")
	}
}
