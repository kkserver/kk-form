package form

import (
	"bytes"
	"fmt"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/dynamic"
	"time"
)

type FormDataService struct {
	app.Service

	Create *FormDataCreateTask
	Set    *FormDataSetTask
	Remove *FormDataRemoveTask
	Get    *FormDataTask
	Query  *FormDataQueryTask
}

func (S *FormDataService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *FormDataService) HandleFormDataCreateTask(a IFormApp, task *FormDataCreateTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := FormData{}

	v.FormId = task.FormId
	v.Uid = task.Uid
	v.Type = task.Type
	v.Content = task.Content
	v.Tags = task.Tags
	v.Ctime = time.Now().Unix()

	_, err = kk.DBInsert(db, a.GetFormDataTable(), a.GetPrefix(), &v)

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Data = &v

	return nil
}

func (S *FormDataService) HandleFormDataSetTask(a IFormApp, task *FormDataSetTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := FormData{}

	rows, err := kk.DBQuery(db, a.GetFormDataTable(), a.GetPrefix(), " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {
		scanner := kk.NewDBScaner(&v)
		err = scanner.Scan(rows)
		if err != nil {
			task.Result.Errno = ERROR_FORM
			task.Result.Errmsg = err.Error()
			return nil
		}
	} else {
		task.Result.Errno = ERROR_FORM_NOT_FOUND
		task.Result.Errmsg = "Not Found form data"
		return nil
	}

	keys := map[string]bool{}

	if task.Type != nil {
		v.Type = dynamic.StringValue(task.Type, v.Type)
		keys["type"] = true
	}

	if task.Tags != nil {
		v.Tags = dynamic.StringValue(task.Tags, v.Tags)
		keys["tags"] = true
	}

	if task.Content != nil {
		v.Content = dynamic.StringValue(task.Content, v.Content)
		keys["content"] = true
	}

	_, err = kk.DBUpdateWithKeys(db, a.GetFormDataTable(), a.GetPrefix(), &v, keys)

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Data = &v

	return nil
}

func (S *FormDataService) HandleFormDataTask(a IFormApp, task *FormDataTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := FormData{}

	rows, err := kk.DBQuery(db, a.GetFormDataTable(), a.GetPrefix(), " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {
		scanner := kk.NewDBScaner(&v)
		err = scanner.Scan(rows)
		if err != nil {
			task.Result.Errno = ERROR_FORM
			task.Result.Errmsg = err.Error()
			return nil
		}
	} else {
		task.Result.Errno = ERROR_FORM_NOT_FOUND
		task.Result.Errmsg = "Not Found form data"
		return nil
	}

	task.Result.Data = &v

	return nil
}

func (S *FormDataService) HandleFormDataRemoveTask(a IFormApp, task *FormDataRemoveTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := FormData{}

	rows, err := kk.DBQuery(db, a.GetFormDataTable(), a.GetPrefix(), " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {

		scanner := kk.NewDBScaner(&v)
		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_FORM
			task.Result.Errmsg = err.Error()
			return nil
		}

		_, err = kk.DBDelete(db, a.GetFormDataTable(), a.GetPrefix(), " WHERE id=?", task.Id)

		if err != nil {
			task.Result.Errno = ERROR_FORM
			task.Result.Errmsg = err.Error()
			return nil
		}

	} else {
		task.Result.Errno = ERROR_FORM_NOT_FOUND
		task.Result.Errmsg = "Not Found form data"
		return nil
	}

	task.Result.Data = &v

	return nil
}

func (S *FormDataService) HandleFormDataQueryTask(a IFormApp, task *FormDataQueryTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	var datas = []FormData{}

	var args = []interface{}{}

	var sql = bytes.NewBuffer(nil)

	sql.WriteString(" WHERE 1")

	if task.Id != 0 {
		sql.WriteString(" AND id=?")
		args = append(args, task.Id)
	}

	if task.Uid != 0 {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.FormId != 0 {
		sql.WriteString(" AND formid=?")
		args = append(args, task.FormId)
	}

	if task.Keyword != "" {
		q := "%" + task.Keyword + "%"
		sql.WriteString(" AND (tags LIKE ?)")
		args = append(args, q)
	}

	if task.OrderBy == "asc" {
		sql.WriteString(" ORDER BY id ASC")
	} else {
		sql.WriteString(" ORDER BY id DESC")
	}

	var pageIndex = task.PageIndex
	var pageSize = task.PageSize

	if pageIndex < 1 {
		pageIndex = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	if task.Counter {
		var counter = FormDataQueryCounter{}
		counter.PageIndex = pageIndex
		counter.PageSize = pageSize
		counter.RowCount, err = kk.DBQueryCount(db, a.GetFormDataTable(), a.GetPrefix(), sql.String(), args...)
		if err != nil {
			task.Result.Errno = ERROR_FORM
			task.Result.Errmsg = err.Error()
			return nil
		}
		if counter.RowCount%pageSize == 0 {
			counter.PageCount = counter.RowCount / pageSize
		} else {
			counter.PageCount = counter.RowCount/pageSize + 1
		}
		task.Result.Counter = &counter
	}

	sql.WriteString(fmt.Sprintf(" LIMIT %d,%d", (pageIndex-1)*pageSize, pageSize))

	var v = FormData{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, a.GetFormDataTable(), a.GetPrefix(), sql.String(), args...)

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	for rows.Next() {

		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_FORM
			task.Result.Errmsg = err.Error()
			return nil
		}

		datas = append(datas, v)
	}

	task.Result.Datas = datas

	return nil
}
