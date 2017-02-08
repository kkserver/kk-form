package form

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FormDataCreateTaskResult struct {
	app.Result
	Data *FormData `json:"data,omitempty"`
}

type FormDataCreateTask struct {
	app.Task
	FormId  int64  `json:"formId"`
	Uid     int64  `json:"uid"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
	Result  FormDataCreateTaskResult
}

func (task *FormDataCreateTask) GetResult() interface{} {
	return &task.Result
}

func (task *FormDataCreateTask) GetInhertType() string {
	return "form"
}

func (task *FormDataCreateTask) GetClientName() string {
	return "Data.Create"
}
