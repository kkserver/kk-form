package form

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FormQueryCounter struct {
	PageIndex int `json:"p"`
	PageSize  int `json:"size"`
	PageCount int `json:"count"`
	RowCount  int `json:"rowCount"`
}

type FormQueryTaskResult struct {
	app.Result
	Counter *FormQueryCounter `json:"counter,omitempty"`
	Forms   []Form            `json:"forms,omitempty"`
}

type FormQueryTask struct {
	app.Task
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Keyword   string `json:"q"`
	OrderBy   string `json:"orderBy"` // desc, asc
	PageIndex int    `json:"p"`
	PageSize  int    `json:"size"`
	Counter   bool   `json:"counter"`
	Result    FormQueryTaskResult
}

func (task *FormQueryTask) GetResult() interface{} {
	return &task.Result
}

func (task *FormQueryTask) GetInhertType() string {
	return "form"
}

func (task *FormQueryTask) GetClientName() string {
	return "Form.Query"
}
