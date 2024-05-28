package tools_test

import (
	"everHold/src/tools"
	"fmt"
	"testing"
)

func TestUAC(t *testing.T) {
	result, error := tools.GetCurrentUACLevel()
	if error != nil {
		t.Fatal(error)
	}
	fmt.Println(result)
}

func TestAdmin(t *testing.T) {
	result, error := tools.IsUserAnAdmin()
	if error != nil {
		t.Fatal(error)
	}
	fmt.Println(result)
}
