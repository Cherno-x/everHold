package utils

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	SEE_MASK_NOCLOSEPROCESS = 0x00000040
	SW_HIDE                 = 0
	SW_SHOW                 = 5
)

var (
	modShell32              = syscall.NewLazyDLL("shell32.dll")
	procShellExecuteExW     = modShell32.NewProc("ShellExecuteExW")
	procWaitForSingleObject = syscall.NewLazyDLL("kernel32.dll").NewProc("WaitForSingleObject")
	procGetExitCodeProcess  = syscall.NewLazyDLL("kernel32.dll").NewProc("GetExitCodeProcess")
)

const (
	INFINITE = 0xFFFFFFFF
)

type SHELLEXECUTEINFO struct {
	cbSize       uint32
	fMask        uint32
	hwnd         uintptr
	lpVerb       *uint16
	lpFile       *uint16
	lpParameters *uint16
	lpDirectory  *uint16
	nShow        int32
	hInstApp     uintptr
	lpIDList     uintptr
	lpClass      *uint16
	hkeyClass    uintptr
	dwHotKey     uint32
	hIcon        uintptr
	hProcess     uintptr
}

// create 函数用于创建进程，并可选地获取退出代码
func CreateProcess(payload string, params string, window bool, getExitCode bool) (int, bool) {
	var shinfo SHELLEXECUTEINFO
	shinfo.cbSize = uint32(unsafe.Sizeof(shinfo))
	shinfo.fMask = SEE_MASK_NOCLOSEPROCESS
	shinfo.lpFile, _ = syscall.UTF16PtrFromString(payload)
	if window {
		shinfo.nShow = SW_SHOW
	} else {
		shinfo.nShow = SW_HIDE
	}
	shinfo.lpParameters, _ = syscall.UTF16PtrFromString(params)

	ret, _, _ := procShellExecuteExW.Call(uintptr(unsafe.Pointer(&shinfo)))
	if ret == 0 {
		return 0, false
	}

	if getExitCode {
		procWaitForSingleObject.Call(shinfo.hProcess, uintptr(INFINITE))
		var exitCode uint32
		procGetExitCodeProcess.Call(shinfo.hProcess, uintptr(unsafe.Pointer(&exitCode)))
		return int(exitCode), true
	}

	return 0, true
}

func main() {
	exitCode, success := CreateProcess("mofcomp.exe", "params", false, true)
	if success {
		fmt.Printf("Exit code: %d\n", exitCode)
	} else {
		fmt.Println("Failed to create process")
	}
}
