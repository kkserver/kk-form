package form

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FormDataQueryCounter struct {
	PageIndex int `json:"p"`
	PageSize  int `json:"size"`
	PageCount int `json:"count"`
	RowCount  int `json:"rowCount"`
}

type FormDataQueryTaskResult struct {
	app.Result
	Counter *FormDataQueryCounter `json:"counter,omitempty"`
	Datas   []FormData            `json:"datas,omitempty"`
}

type FormDataQueryTask struct {
	app.Task
	Id        int64  `json:"id"`
	Uid       int64  `json:"uid"`
	FormId    int64  `json:"formId"`
	Keyword   string `json:"q"`
	OrderBy   string `json:"orderBy"` // desc, asc
	PageIndex int    `json:"p"`
	PageSize  int    `json:"size"`
	Counter   bool   `json:"counter"`
	Result    FormDataQueryTaskResult
}

func (task *FormDataQueryTask) GetResult() interface{} {
	return &task.Result
}

func (task *FormDataQueryTask) GetInhertType() string {
	return "form"
}

func (task *FormDataQueryTask) GetClientName() string {
	return "Data.Query"
}
