package detect

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"everHold/src/utils"

	"github.com/StackExchange/wmi"
)

type OSVERSIONINFOEX struct {
	dwOSVersionInfoSize uint32
	dwMajorVersion      uint32
	dwMinorVersion      uint32
	dwBuildNumber       uint32
	dwPlatformId        uint32
	szCSDVersion        [128]uint16
	wServicePackMajor   uint16
	wServicePackMinor   uint16
	wSuiteMask          uint16
	wProductType        byte
	wReserved           byte
}

type WINDOWSVERSION struct {
	dwMajorVersion uint32
	dwMinorVersion uint32
	dwBuildNumber  uint32
}

type AntivirusProcess struct {
	ProcessName string
	DisplayName string
}

type Win32_Processor struct {
	Name                          string
	VirtualizationFirmwareEnabled bool
}

var (
	modNtdll          = syscall.NewLazyDLL("ntdll.dll")
	procRtlGetVersion = modNtdll.NewProc("RtlGetVersion")
)

func GetWindowVersion() (WINDOWSVERSION, error) {
	plantform := runtime.GOOS
	if plantform != "windows" {
		utils.PrintError("非windows平台，功能不支持")
		return WINDOWSVERSION{}, fmt.Errorf("非Windows平台，功能不支持")
	}

	var version OSVERSIONINFOEX
	version.dwOSVersionInfoSize = uint32(unsafe.Sizeof(version))

	ret, _, err := procRtlGetVersion.Call(uintptr(unsafe.Pointer(&version)))
	if ret != 0 {
		utils.PrintError("无法获取 Windows 版本信息")
		return WINDOWSVERSION{}, fmt.Errorf("无法获取 Windows 版本信息: %v", err)
	}

	fmt.Sprintf("Windows 版本: %d.%d (Build %d)",
		version.dwMajorVersion, version.dwMinorVersion, version.dwBuildNumber)

	var windowsVersion WINDOWSVERSION
	windowsVersion.dwMajorVersion = version.dwMajorVersion
	windowsVersion.dwMinorVersion = version.dwMinorVersion
	windowsVersion.dwBuildNumber = version.dwBuildNumber

	return windowsVersion, nil

}

// 检查系统中是否存在常见的杀毒软件进程
func CheckAntivirusProcesses() ([]AntivirusProcess, error) {
	// 常见的杀毒软件进程名称列表,完整列表在av.txt
	antivirusProcesses := []AntivirusProcess{
		{"360tray.exe", "360安全卫士-实时保护"},
		{"360safe.exe", "360安全卫士-主程序"},
		{"ZhuDongFangYu.exe", "360安全卫士-主动防御"},
		{"360skylarsvc", "360终端安全管理系统"},
		{"360WebSafe.exe", "360主机卫士Web"},
		{"QHSrv.exe", "360主机卫士Web"},
		{"QHWebshellGuard.exe", "360主机卫士Web"},
		{"LiveUpdate360.exe", "360杀毒"},
		{"360sd.exe", "360杀毒"},
		{"MsMpEng.exe", "Microsoft Security Essentials"},
		{"NisSrv.exe", "Microsoft Security Essentials"},
		{"MsSense.exe", "Microsoft Security Essentials"},
		{"msseces.exe", "Microsoft Security Essentials"},
		{"MpCmdRun.exe", "Windows Defender Antivirus"},
		{"avp.exe", "Kaspersky"},
		{"KvMonXP.exe", "江民杀毒"},
		{"RavMonD.exe", "瑞星杀毒"},
		{"Mcshield.exe", "Mcafee"},
		{"Tbmon.exe", "Mcafee"},
		{"Frameworkservice.exe", "Mcafee"},
		{"kxetray.exe", "金山毒霸"},
		{"kxescore.exe", "金山毒霸"},
		{"kupdata.exe", "金山毒霸"},
		{"kwsprotect64.exe", "金山毒霸"},
		{"avcenter.exe", "Avira(小红伞)"},
		{"avguard.exe", "Avira(小红伞)"},
		{"avgnt.exe", "Avira(小红伞)"},
		{"sched.exe", "Avira(小红伞)"},
		{"ashDisp.exe", "Avast网络安全"},
		{"rtvscan.exe", "诺顿杀毒"},
		{"ccapp.exe", "Symantec Norton"},
		{"NPFMntor.exe", "Norton杀毒软件相关进程"},
		{"ccSetMgr.exe", "赛门铁克"},
		{"ccRegVfy.exe", "Norton杀毒软件自身完整性检查程序"},
		{"vptray.exe", "Norton病毒防火墙-盾牌图标程序"},
		{"ksafe.exe", "金山卫士"},
		{"QQPCRTP.exe", "QQ电脑管家"},
		{"mssecess.exe", "微软杀毒"},
		{"beikesan.exe", "贝壳云安全"},
		{"KSWebShield.exe", "金山网盾"},
		{"SafeDogGuardCenter.exe", "安全狗"},
		{"safedogupdatecenter.exe", "安全狗"},
		{"safedogguardcenter.exe", "安全狗"},
		{"SafeDogSiteIIS.exe", "安全狗"},
		{"SafeDogTray.exe", "安全狗"},
		{"SafeDogServerUI.exe", "安全狗"},
		{"D_Safe_Manage.exe", "D盾"},
		{"d_manage.exe", "D盾"},
		{"hipstray.exe", "火绒"},
		{"wsctrl.exe", "火绒"},
		{"usysdiag.exe", "火绒"},
		{"HipsDaemon.exe", "火绒"},
		{"HipsLog.exe", "火绒"},
		{"HipsMain.exe", "火绒"},
		{"Notifier.exe", "亚信安全服务器深度安全防护系统"},
	}

	// 获取所有运行中的进程
	output, err := exec.Command("tasklist").Output()
	if err != nil {
		return nil, fmt.Errorf("无法获取进程列表: %v", err)
	}

	// 将输出转换为字符串并检查是否包含任何杀毒软件进程
	processList := strings.ToLower(string(output))
	var detectedProcesses []AntivirusProcess
	for _, process := range antivirusProcesses {
		if strings.Contains(processList, strings.ToLower(process.ProcessName)) {
			detectedProcesses = append(detectedProcesses, process)
		}
	}
	return detectedProcesses, nil
}

func Check360hejing() bool {
	var processors []Win32_Processor
	query := "SELECT Name, VirtualizationFirmwareEnabled FROM Win32_Processor"
	err := wmi.Query(query, &processors)
	if err != nil {
		utils.PrintError("wmi执行失败")
	}

	for _, processor := range processors {
		fmt.Printf("Processor: %s\n", processor.Name)
		fmt.Printf("Virtualization Enabled: %t\n", processor.VirtualizationFirmwareEnabled)
		if processor.VirtualizationFirmwareEnabled == true {
			return true
		}
	}
	return false
}
