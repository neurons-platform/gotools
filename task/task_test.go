package task

import (
	"testing"
	U "github.com/jingminglang/gotools/utils"
)

type TestFunc struct {

}

type TestCondition struct{
}

type TestAction struct {
}

func (t TestFunc) DoFunc() FuncResult {
	r := FuncResult{}
	r.Result = "do test func"
	U.LogPrintln(r.Result)
	return r
}

func (c TestCondition) CheckCondition(r FuncResult) ConditionResult {
	cr := ConditionResult{}
	cr.Result = "check test condition"
	U.LogPrintln(cr.Result)
	return cr
}

func (a  TestAction) DoAction(c ConditionResult) ActionResult {
	ar := ActionResult{}
	ar.Result = "do action"
	U.LogPrintln(ar.Result)
	return ar
}


type TestCallBack struct {
}

func (c TestCallBack) CallBack(ar ActionResult) {
	U.LogPrintln("call back" + ar.Result)
}


func TestStartCrontab(t *testing.T) {
	task := Task{}
	task.Id = "test"
	task.Name = "test"
	// task.Spec = "* * * * * *"
	task.Spec = "0 * * * * *"
	task.Rule = Rule{}
	task.Rule.Func = TestFunc{}
	task.Rule.Condition = TestCondition{}
	task.Rule.Action = TestAction{}
	U.LogPrintln(task)

	StartCrontab()
	AddTask(task, TestCallBack{})

	select {}
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)
}
