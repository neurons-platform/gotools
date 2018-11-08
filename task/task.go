package task

import (
	"gopkg.in/robfig/cron.v2"
	U "github.com/jingminglang/gotools/utils"
)
var C *cron.Cron =  cron.New()
var TaskList = make(map[string]cron.EntryID)


func StartCrontab() {
	C.Start()
}


func BuildTaskFunc(task Task,callBack TaskCallBack) (f func ()) {
	f = func () {
		funcResult := task.Rule.Func.DoFunc()
		conditonResult :=  task.Rule.Condition.CheckCondition(funcResult)
		actionResult := task.Rule.Action.DoAction(conditonResult)
		callBack.CallBack(actionResult)
	}
	return f
}


func AddTask(task Task,callBack TaskCallBack) {
        f:= BuildTaskFunc(task,callBack)
	id ,err :=C.AddFunc(task.Spec,f)
	if(U.Throw(err)) {
		TaskList[task.Id]=id
	}
}

func RemoveTask(task Task) {
	id := TaskList[task.Id]
	delete(TaskList,task.Id)
	C.Remove(id)
}



type TaskCallBack interface {
	CallBack(ActionResult)
}


type Task struct {
	Id string
	Name string
	Spec string
	Rule Rule
}

type Rule struct {
	Func  Func
	Condition Condition
	Action Action
}

type FuncResult struct {
	Result string
}

type ConditionResult struct {
	Result string
}

type ActionResult struct {
	Result string
}

type Func interface {
	DoFunc() FuncResult
}

type Condition interface {
	CheckCondition(FuncResult) ConditionResult
}

type Action interface {
	DoAction(ConditionResult) ActionResult
}
