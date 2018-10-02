package linuxOS

import (
	"testing"

	U "github.com/jingminglang/gotools/utils"
)

func TestYamlStrToLinuxOS(t *testing.T) {
	f := U.ReadAllFile("./yaml/centos.yaml")
	got := YamlStrToLinuxOS(f)
	U.LogPrintln(got)
}

func TestLinuxOS_FindCMD(t *testing.T) {
	f := U.ReadAllFile("./yaml/centos.yaml")
	linuxOS := YamlStrToLinuxOS(f)
	// cmd,_ := got.FindCMD("date")
	cmd := linuxOS.FindCMD("ntpdate").FindUseCase("sync datetime")("")
	U.LogPrintln(cmd)
	cmd = linuxOS.FindCMD("date").FindUseCase("show time zone")("")
	U.LogPrintln(cmd)
	cmd = linuxOS.FindCMD("jstack").FindUseCase("11")("")
	U.LogPrintln(cmd)
	cmd = linuxOS.FindCMD("nginx").FindUseCase("nginx reload")("")
	U.LogPrintln(cmd)
	cmd = linuxOS.FindCMD("yum").FindUseCase("install")("ntpdate")
	U.LogPrintln(cmd)
}

func TestLinuxOS_FindScript(t *testing.T) {
	f := U.ReadAllFile("./yaml/centos.yaml")
	linuxOS := YamlStrToLinuxOS(f)
	// sh := linuxOS.FindScript("./tpl/jstack_top_thread.tpl",`"server1" 10`)
	// U.LogPrintln(sh)
	// sh := linuxOS.FindScript("./tpl/jstat.tpl",`"server1"`)
	// U.LogPrintln(sh)
	sh := linuxOS.FindScript("./tpl/nginx_top_ip.tpl",`"server1"`)
	U.LogPrintln(sh)
}
