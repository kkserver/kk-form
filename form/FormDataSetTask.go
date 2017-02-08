package form

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FormDataSetTaskResult struct {
	app.Result
	Data *FormData `json:"data,omitempty"`
}

type FormDataSetTask struct {
	app.Task
	Id      int64       `json:"id"`
	Type    interface{} `json:"type"`
	Content interface{} `json:"content"`
	Tags    interface{} `json:"tags"`
	Result  FormDataSetTaskResult
}

func (task *FormDataSetTask) GetResult() interface{} {
	return &task.Result
}

func (task *FormDataSetTask) GetInhertType() string {
	return "form"
}

func (task *FormDataSetTask) GetClientName() string {
	return "Data.Set"
}
