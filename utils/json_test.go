package utils

import (
	"testing"
)

func TestJsonStrToJsonSJ(t *testing.T) {
	data := `
{
  "requestPara": {
    "startTime": "20181015164200",
    "scopeValues": "chat.message.deliver",
    "dataType": "TP99",
    "dagaCycle": "oneMinute",
    "scope": "Key",
    "monitorType": "Method",
    "endTime": "20181015171200"
  },
  "responseInfo": {
    "message": "",
    "state": "success"
  },
  "result": {
    "chat.message.deliver": [
      {},
      {
        "TP99": 115,
        "dataTime": "20181015171200"
      },
      {
        "TP99": 86,
        "dataTime": "20181015171100"
      },
      {
        "TP99": 100,
        "dataTime": "20181015171000"
      },
      {
        "TP99": 105,
        "dataTime": "20181015170900"
      },
      {
        "TP99": 91,
        "dataTime": "20181015170800"
      },
      {
        "TP99": 96,
        "dataTime": "20181015170700"
      },
      {
        "TP99": 106,
        "dataTime": "20181015170600"
      },
      {
        "TP99": 108,
        "dataTime": "20181015170500"
      },
      {
        "TP99": 106,
        "dataTime": "20181015170400"
      },
      {
        "TP99": 106,
        "dataTime": "20181015170300"
      },
      {
        "TP99": 115,
        "dataTime": "20181015170200"
      },
      {
        "TP99": 110,
        "dataTime": "20181015170100"
      },
      {
        "TP99": 109,
        "dataTime": "20181015170000"
      },
      {
        "TP99": 102,
        "dataTime": "20181015165900"
      },
      {
        "TP99": 102,
        "dataTime": "20181015165800"
      },
      {
        "TP99": 104,
        "dataTime": "20181015165700"
      },
      {
        "TP99": 111,
        "dataTime": "20181015165600"
      },
      {
        "TP99": 107,
        "dataTime": "20181015165500"
      },
      {
        "TP99": 104,
        "dataTime": "20181015165400"
      },
      {
        "TP99": 112,
        "dataTime": "20181015165300"
      },
      {
        "TP99": 107,
        "dataTime": "20181015165200"
      },
      {
        "TP99": 104,
        "dataTime": "20181015165100"
      },
      {
        "TP99": 111,
        "dataTime": "20181015165000"
      },
      {
        "TP99": 106,
        "dataTime": "20181015164900"
      },
      {
        "TP99": 110,
        "dataTime": "20181015164800"
      },
      {
        "TP99": 109,
        "dataTime": "20181015164700"
      },
      {
        "TP99": 94,
        "dataTime": "20181015164600"
      },
      {
        "TP99": 103,
        "dataTime": "20181015164500"
      },
      {
        "TP99": 105,
        "dataTime": "20181015164400"
      },
      {
        "TP99": 114,
        "dataTime": "20181015164300"
      }
    ]
  }
}
`
	got, err := JsonStrToJsonSJ(data)
	Throw(err)
	key,_:=got.Get("requestPara").Get("scopeValues").String()
	dataType,_:=got.Get("requestPara").Get("dataType").String()
	LogPrintln(dataType)
	r,_ := got.Get("result").Get(key).Array()
	for _,v := range r {
		m := v.(map[string]interface{})
		if val, ok := m[dataType]; ok {
			LogPrintln(val)
		}
		if val, ok := m["dataTime"]; ok {
			LogPrintln(val)
		}

	}
}
