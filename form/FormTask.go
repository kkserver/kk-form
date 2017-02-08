package form

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FormTaskResult struct {
	app.Result
	Form *Form `json:"form,omitempty"`
}

type FormTask struct {
	app.Task
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Result FormTaskResult
}

func (task *FormTask) GetResult() interface{} {
	return &task.Result
}

func (task *FormTask) GetInhertType() string {
	return "form"
}

func (task *FormTask) GetClientName() string {
	return "Form.Get"
}
