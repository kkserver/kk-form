package form

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FormCreateTaskResult struct {
	app.Result
	Form *Form `json:"form,omitempty"`
}

type FormCreateTask struct {
	app.Task
	Name    string `json:"name"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
	Result  FormCreateTaskResult
}

func (task *FormCreateTask) GetResult() interface{} {
	return &task.Result
}

func (task *FormCreateTask) GetInhertType() string {
	return "form"
}

func (task *FormCreateTask) GetClientName() string {
	return "Form.Create"
}
