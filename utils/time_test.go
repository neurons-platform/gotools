package utils


import (
	"testing"
)


func TestParserStringToTime(t *testing.T) {
	tm := ParserStringToTime("20181015164300")
	LogPrintln(tm)
}
