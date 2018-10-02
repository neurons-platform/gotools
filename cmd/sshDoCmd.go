package cmd

import (
	// U "neurons_master/utils"
	U "github.com/jingminglang/gotools/utils"
)

func DoCmd(cmd CMD) (CMD, bool) {
	if cmd.CMDType == SSH {
		out := U.ExecuteCmd(cmd.IP, cmd.Command, cmd.User, cmd.Password, U.Int2Str(cmd.Port))
		if val, ok := out[cmd.IP]; ok {
			cmd.Out = val
			return cmd, true
		}
	}
	return cmd, false
}
