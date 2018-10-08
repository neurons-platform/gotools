package linuxOS

type CmdArg struct {
	Name string `yaml:"Name"`
	Desc string `yaml:"Desc"`
	Args string `yaml:"Args"`
}

type LinuxCMD struct {
	// 命令名
	Name       string   `yaml:"Name"`
	// 命令所在位置 fullpath
	Path       string   `yaml:"Path"`
	// 执行命令的程序
	ExeCmd string `yaml:"ExeCmd"`
	// 日志位置
	LogPath    string    `yaml:"LogPath"`
	// 命令版本
	Version    string   `yaml:"Version"`
	// 使用方法
	UseCase    []CmdArg `yaml:"UseCase"`
	// 安装方法
	InstallCMD string   `yaml:"InstallCMD"`
	// 描述
	Desc       string   `yaml:"Desc"`
}

func (c LinuxCMD) FindUseCase(useCase string) func(string) string {
	for _, v := range c.UseCase {
		return func(str string) string {
			if len(c.ExeCmd) >0 {
				return c.ExeCmd + " " +  c.Path + " " + v.Args + " " + str
			} else {
				return c.Path + " " + v.Args + " " + str
			}
		}
	}

	return func(str string) string {
		if len(c.ExeCmd) >0 {
			return c.ExeCmd + " " +  useCase + " " + str
		} else {
			return c.Path + " " + useCase + " " + str
		}
	}
}
