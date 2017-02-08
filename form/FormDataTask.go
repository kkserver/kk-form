package form

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FormDataTaskResult struct {
	app.Result
	Data *FormData `json:"data,omitempty"`
}

type FormDataTask struct {
	app.Task
	Id     int64 `json:"id"`
	Result FormDataTaskResult
}

func (task *FormDataTask) GetResult() interface{} {
	return &task.Result
}

func (task *FormDataTask) GetInhertType() string {
	return "form"
}

func (task *FormDataTask) GetClientName() string {
	return "Data.Get"
}
