package utils_test

import (
	"everHold/src/utils"
	"fmt"
	"testing"
)

func TestUAC(t *testing.T) {
	result, error := utils.GetCurrentUACLevel()
	if error != nil {
		t.Fatal(error)
	}
	fmt.Println(result)
}

func TestAdmin(t *testing.T) {
	result, error := utils.IsUserAnAdmin()
	if error != nil {
		t.Fatal(error)
	}
	fmt.Println(result)
}
