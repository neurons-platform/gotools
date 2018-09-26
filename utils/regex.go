package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func StrToTemplate(format string) string {
	re := regexp.MustCompile(`\$([a-zA-Z_0-9]+)`).ReplaceAllString(format+" ", "{{.$1}}")
	formatReg := strings.Trim(re, " ")
	return formatReg
}

func FormatToReg2(format string) *regexp.Regexp {
	re := regexp.MustCompile(`\\\$([a-zA-Z_0-9]+)((.)([^\\\$])?)`).ReplaceAllString(
		regexp.QuoteMeta(format+" "), "(?P<$1>[^$3$4]*)$2")
	format_reg := regexp.MustCompile(fmt.Sprintf("^%v", strings.Trim(re, " ")))
	return format_reg
}

func FormatToReg(format string) *regexp.Regexp {
	re := regexp.MustCompile(`\\\$([a-zA-Z_0-9]+)((.))`).ReplaceAllString(
		regexp.QuoteMeta(format+" "), "(?P<$1>[^$3]*)$2")
	format_reg := regexp.MustCompile(fmt.Sprintf("^%v", strings.Trim(re, " ")))
	return format_reg
}

func ParserStringToMap(str string, re *regexp.Regexp) (map[string]string, bool) {
	fields := make(map[string]string)
	f := re.FindStringSubmatch(str)

	if len(f) != len(re.SubexpNames()) {
		return fields, false
	}

	for i, name := range re.SubexpNames() {
		fields[name] = f[i]

	}
	return fields, true
}

func GetIPs(ipsStr string) []string {
	reg := regexp.MustCompile(`[0-9]{1,3}(\.[0-9]{1,3}){3}`)
	return reg.FindAllString(ipsStr, -1)
}
