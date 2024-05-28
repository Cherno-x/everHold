package elevate

import (
	"everHold/src/conf"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

func Runasadmin(newcmd *conf.RUNCMD) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	verb := "runas"

	// 使用 ShellExecute 提升权限
	operation, _ := syscall.UTF16PtrFromString(verb)
	executable, _ := syscall.UTF16PtrFromString("cmd.exe")
	params, _ := syscall.UTF16PtrFromString("/k \"" + exe + "\" " + newcmd.ModuleValue + " --method " + newcmd.MethodValue + " --payload " + newcmd.PayloadValue + " --add " + strconv.FormatBool(newcmd.AddValue) + " --name " + newcmd.NameValue)
	dir, _ := syscall.UTF16PtrFromString("")

	var showCmd int32 = 1 // SW_SHOWNORMAL

	ret, _, err := syscall.NewLazyDLL("shell32.dll").NewProc("ShellExecuteW").Call(
		uintptr(0),
		uintptr(unsafe.Pointer(operation)),
		uintptr(unsafe.Pointer(executable)),
		uintptr(unsafe.Pointer(params)),
		uintptr(unsafe.Pointer(dir)),
		uintptr(showCmd),
	)

	if ret <= 32 {
		return fmt.Errorf("ShellExecuteW failed with error code: %d", ret)
	}
	return nil
}
