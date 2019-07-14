package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

var W *bufio.Writer

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func InitLog(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
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
