package form

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FormRemoveTaskResult struct {
	app.Result
	Form *Form `json:"form,omitempty"`
}

type FormRemoveTask struct {
	app.Task
	Id     int64 `json:"id"`
	Result FormRemoveTaskResult
}

func (task *FormRemoveTask) GetResult() interface{} {
	return &task.Result
}

func (task *FormRemoveTask) GetInhertType() string {
	return "form"
}

func (task *FormRemoveTask) GetClientName() string {
	return "Form.Remove"
}
