package elevate_test

import (
	"everHold/src/conf"
	"everHold/src/func/elevate"

	"testing"
)

func TestRunadmin(t *testing.T) {
	newcmd := conf.NewRunCMD("", "", true, "", "")
	err := elevate.Runasadmin(newcmd)
	if err != nil {
		t.Fatal(err)
	}
}
