package linuxOS

type CmdArg struct {
	Name string `yaml:"Name"`
	Desc string `yaml:"Desc"`
	Args string `yaml:"Args"`
}

type LinuxCMD struct {
	Name       string   `yaml:"Name"`
	Path       string   `yaml:"Path"`
	LogPath    string    `yaml:"LogPath"`
	Version    string   `yaml:"Version"`
	UseCase    []CmdArg `yaml:"UseCase"`
	InstallCMD string   `yaml:"InstallCMD"`
	Desc       string   `yaml:"Desc"`
}

func (c LinuxCMD) FindUseCase(useCase string) func(string) string {
	for _, v := range c.UseCase {
		return func(str string) string {
			return c.Path + " " + v.Args + " " + str
		}
	}

	return func(str string) string {
		return c.Path + " " + useCase + " " + str
	}
}
