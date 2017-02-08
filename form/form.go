package form

import (
	"database/sql"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
)

type Form struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Content string `json:"content,omitempty"`
	Tags    string `json:"tags"`
	Ctime   int64  `json:"ctime"`
}

type FormData struct {
	Id      int64  `json:"id"`
	FormId  int64  `json:"formId"`
	Uid     int64  `json:"uid"`
	Type    string `json:"type"`
	Content string `json:"content,omitempty"`
	Tags    string `json:"tags"`
	Ctime   int64  `json:"ctime"`
}

type IFormApp interface {
	app.IApp
	GetDB() (*sql.DB, error)
	GetPrefix() string
	GetFormTable() *kk.DBTable
	GetFormDataTable() *kk.DBTable
}

type FormApp struct {
	app.App
	DB *app.DBConfig

	Remote *remote.Service

	Form *FormService
	Data *FormDataService

	FormTable     kk.DBTable
	FormDataTable kk.DBTable
}

func (C *FormApp) GetDB() (*sql.DB, error) {
	return C.DB.Get(C)
}

func (C *FormApp) GetPrefix() string {
	return C.DB.Prefix
}

func (C *FormApp) GetFormTable() *kk.DBTable {
	return &C.FormTable
}

func (C *FormApp) GetFormDataTable() *kk.DBTable {
	return &C.FormDataTable
}
