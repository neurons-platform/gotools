package utils

import (
	"bufio"
	"fmt"
	"time"
)

var W *bufio.Writer

func TimeCost(start time.Time) {
	dis := time.Now().Sub(start).Seconds() * 1000
	LogPrint("消耗时间:")
	LogPrint(dis)
	LogPrintln("ms")
}

func LogPrint(str interface{}) {
	_, err := fmt.Fprintf(W, "%v", str)
	W.Flush()
	fmt.Print(str)
	Throw(err)
}

func LogPrintf(format string, str ...interface{}) {
	_, err := fmt.Fprintf(W, format, str...)
	W.Flush()
	fmt.Printf(format, str...)
	Throw(err)
}

func LogPrintln(str ...interface{}) {
	fmt.Printf("%+v\n", str)
	// _, err := fmt.Fprintf(W, "%v\n", str...)
	// W.Flush()
	// fmt.Println(str...)
	// Throw(err)
}
