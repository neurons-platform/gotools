package linuxOS

import (
	"bytes"
	F "github.com/neurons-platform/gotools/file"
	U "github.com/neurons-platform/gotools/utils"
	"text/template"
)

type LinuxType int

const (
	Centos LinuxType = iota
	Ubuntu
)

type LinuxOS struct {
	LinuxType   LinuxType
	Version     string     `yaml:"Version"`
	DefaultUser string     `yaml:"DefaultUser"`
	LinuxCMDs   []LinuxCMD `yaml:"LinuxCMDs"`
}

// func (c LinuxOS) FindCMD(cmdName string) (LinuxCMD,bool) {
// 	for _,v := range c.LinuxCMDs {
// 		if cmdName == v.Name {
// 			return v,true
// 		}
// 	}
// 	return LinuxCMD{},false
// }

func (c LinuxOS) FindCMD(cmdName string) LinuxCMD {
	for _, v := range c.LinuxCMDs {
		if cmdName == v.Name {
			return v
		}
	}
	return LinuxCMD{}
}

func (c LinuxOS) GetJstackPath() string {
	return c.FindCMD("jstack").Path
}

func (c LinuxOS) GetJstatPath() string {
	return c.FindCMD("jstat").Path
}

func (c LinuxOS) GetNginxLogPath() string {
	return c.FindCMD("nginx").LogPath
}

func (c LinuxOS) FindScript(filePath string, args string) string {
	var buf bytes.Buffer
	tpl := F.ReadAllFile(filePath)
	tmpl, err := template.New("").Parse(tpl)
	if !U.Throw(err) {
		return ""
	}
	err = tmpl.Execute(&buf, c)
	if !U.Throw(err) {
		return ""
	}

	tmpl, err = template.New("").Parse(`
f {{.}}
`)
	if !U.Throw(err) {
		return ""
	}
	err = tmpl.Execute(&buf, args)
	if !U.Throw(err) {
		return ""
	}

	return buf.String()

}

func YamlStrToLinuxOS(str string) LinuxOS {
	linuxOS := LinuxOS{}
	U.YamlStrToStruct(str, &linuxOS)
	return linuxOS
}
