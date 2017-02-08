package form

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FormSetTaskResult struct {
	app.Result
	Form *Form `json:"form,omitempty"`
}

type FormSetTask struct {
	app.Task
	Id      int64       `json:"id"`
	Title   interface{} `json:"title"`
	Summary interface{} `json:"summary"`
	Type    interface{} `json:"type"`
	Content interface{} `json:"content"`
	Tags    interface{} `json:"tags"`
	Result  FormSetTaskResult
}

func (task *FormSetTask) GetResult() interface{} {
	return &task.Result
}

func (task *FormSetTask) GetInhertType() string {
	return "form"
}

func (task *FormSetTask) GetClientName() string {
	return "Form.Set"
}
