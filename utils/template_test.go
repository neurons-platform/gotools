package utils

import (
	"testing"
)

func TestParserJsonStrToTemplateStr(t *testing.T) {
	js := `
{
  "result": {
    "model": "蔡装",
    "venderId": 7388,
    "state": 1,
    "supplierCode": "dm",
    "pid": "42710",
    "ebrandName": "ZISS",
    "brandId": 273,
    "brandName": "蔡司",
    "url": "http://item.jd.com/480.html",
    "size": "",
    "category": "1315,1346,9789",
    "class1": 1315,
    "color": "1.60非球面防蓝光膜",
    "class2": 1346,
    "shopId": 1078,
    "class3": 9789,
    "name": "蔡装",
    "imgurl": "http://df.jpg"
  },
  "code": 1
}
`
	tpl := `= clean {{.result.pid}} {{.result.venderId | tostring}} {{.result.supplierCode}}`
	r, ok := ParserJsonStrToTemplateStr(js, tpl)
	if ok {
		LogPrintln(r)
	}

}

func TestParserMapToString(t *testing.T) {
	mp := map[string]string{"pid":"12"}
	tpl := `
	f() {
             Pid=""
             while read pid
             do
                if $(cat /proc/$pid/environ |strings  |grep $1 2>&1 >/dev/null)
                then
                     Pid=$pid
                fi
             done <  <(ps -ef |grep java |grep -v grep  |awk '{print $2}')
             if [ ! -z "$Pid" ]
             then
                  /export/servers/jdk1.6.0_25/bin/jstat  -gcutil -h10 $Pid  1s 5
             else
      	          echo "not find process"
             fi
        }
        f {{.pid}}
`
	got := ParserMapToString(mp, tpl)
	LogPrintln(got)
}
