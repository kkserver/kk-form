package form

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FormDataRemoveTaskResult struct {
	app.Result
	Data *FormData `json:"data,omitempty"`
}

type FormDataRemoveTask struct {
	app.Task
	Id     int64 `json:"id"`
	Result FormDataRemoveTaskResult
}

func (task *FormDataRemoveTask) GetResult() interface{} {
	return &task.Result
}

func (task *FormDataRemoveTask) GetInhertType() string {
	return "form"
}

func (task *FormDataRemoveTask) GetClientName() string {
	return "Data.Remove"
}
