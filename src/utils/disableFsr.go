package utils

import (
	"syscall"
	"unsafe"
)

var (
	modkernel32                        = syscall.NewLazyDLL("kernel32.dll")
	procWow64DisableWow64FsRedirection = modkernel32.NewProc("Wow64DisableWow64FsRedirection")
	procWow64RevertWow64FsRedirection  = modkernel32.NewProc("Wow64RevertWow64FsRedirection")
)

// disableFSR 函数用于禁用 Windows 文件系统重定向
func DisableFSR() (uintptr, error) {
	var oldValue uintptr
	ret, _, err := procWow64DisableWow64FsRedirection.Call(uintptr(unsafe.Pointer(&oldValue)))
	if ret == 0 {
		return 0, err
	}
	return oldValue, nil
}

// revertFSR 函数用于恢复 Windows 文件系统重定向
func RevertFSR(oldValue uintptr) error {
	ret, _, err := procWow64RevertWow64FsRedirection.Call(oldValue)
	if ret == 0 {
		return err
	}
	return nil
}
