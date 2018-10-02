package cmd


import (
	"time"
	"fmt"
)

func ParallelDo(cmds []CMD,timeOut int) []CMD {
	result := make(chan CMD, 200)
	timeout := time.After(time.Duration(timeOut) * time.Second)
	for _, cmd := range cmds {
		go func(cmd CMD) {
			r,_  := DoCmd(cmd)
			result <- r
		}(cmd)
	}
	var rs []CMD

	// for _, cmd := range cmds {
	// 	timeout := time.After(time.Duration(cmd.TimeOut) * time.Second)
	// 	select {
	// 	case res := <-result:
	// 		rs = append(rs,res)
	// 	case <-timeout:
	// 		rs = append(rs,cmd)
	// 		fmt.Println("Timed out!")
	// 	}
	// }

	for i := 0; i < len(cmds); i++ {
		select {
		case res := <-result:
			rs = append(rs,res)
		case <-timeout:
			fmt.Println("Timed out!")
		}
	}
	return rs

}
