package tools

import (
	"syscall"

	"golang.org/x/sys/windows/registry"
)

// 0：永远不通知
// 1：在更改系统设置时通知
// 2：在安装应用时通知
// 3：始终通知
func GetCurrentUACLevel() (uint64, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`, registry.QUERY_VALUE)
	if err != nil {
		return 0, err
	}
	defer key.Close()

	val, _, err := key.GetIntegerValue("ConsentPromptBehaviorAdmin")
	if err != nil {
		return 0, err
	}

	return val, nil
}

func IsUserAnAdmin() (bool, error) {
	shell32, err := syscall.LoadLibrary("shell32.dll")
	if err != nil {
		return false, err
	}
	defer syscall.FreeLibrary(shell32)

	isUserAnAdminProc, err := syscall.GetProcAddress(shell32, "IsUserAnAdmin")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(uintptr(isUserAnAdminProc))
	return ret != 0, nil
}
