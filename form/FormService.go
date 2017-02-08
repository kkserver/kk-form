package form

import (
	"bytes"
	"fmt"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/dynamic"
	"time"
)

type FormService struct {
	app.Service

	Create *FormCreateTask
	Set    *FormSetTask
	Remove *FormRemoveTask
	Get    *FormTask
	Query  *FormQueryTask
}

func (S *FormService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *FormService) HandleFormCreateTask(a IFormApp, task *FormCreateTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	if task.Name != "" {

		count, err := kk.DBQueryCount(db, a.GetFormTable(), a.GetPrefix(), " WHERE name=?", task.Name)

		if err != nil {
			task.Result.Errno = ERROR_FORM
			task.Result.Errmsg = err.Error()
			return nil
		}

		if count > 0 {
			task.Result.Errno = ERROR_FORM_NAME
			task.Result.Errmsg = "Article name already exists"
			return nil
		}
	}

	v := Form{}

	v.Name = task.Name
	v.Title = task.Title
	v.Summary = task.Summary
	v.Type = task.Type
	v.Content = task.Content
	v.Tags = task.Tags
	v.Ctime = time.Now().Unix()

	_, err = kk.DBInsert(db, a.GetFormTable(), a.GetPrefix(), &v)

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Form = &v

	return nil
}

func (S *FormService) HandleFormSetTask(a IFormApp, task *FormSetTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Form{}

	rows, err := kk.DBQuery(db, a.GetFormTable(), a.GetPrefix(), " WHERE id=?", task.Id)

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
		task.Result.Errmsg = "Not Found form"
		return nil
	}

	keys := map[string]bool{}

	if task.Title != nil {
		v.Title = dynamic.StringValue(task.Title, v.Title)
		keys["title"] = true
	}

	if task.Summary != nil {
		v.Summary = dynamic.StringValue(task.Summary, v.Summary)
		keys["summary"] = true
	}

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

	_, err = kk.DBUpdateWithKeys(db, a.GetFormTable(), a.GetPrefix(), &v, keys)

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Form = &v

	return nil
}

func (S *FormService) HandleFormTask(a IFormApp, task *FormTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Form{}

	if task.Id != 0 {

		rows, err := kk.DBQuery(db, a.GetFormTable(), a.GetPrefix(), " WHERE id=?", task.Id)

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
			task.Result.Errmsg = "Not Found form"
			return nil
		}

	} else if task.Name != "" {

		rows, err := kk.DBQuery(db, a.GetFormTable(), a.GetPrefix(), " WHERE name=?", task.Name)

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
			task.Result.Errmsg = "Not Found form"
			return nil
		}

	} else {
		task.Result.Errno = ERROR_FORM_NOT_FOUND
		task.Result.Errmsg = "Not Found form"
		return nil
	}

	task.Result.Form = &v

	return nil
}

func (S *FormService) HandleFormRemoveTask(a IFormApp, task *FormRemoveTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Form{}

	rows, err := kk.DBQuery(db, a.GetFormTable(), a.GetPrefix(), " WHERE id=?", task.Id)

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

		_, err = kk.DBDelete(db, a.GetFormTable(), a.GetPrefix(), " WHERE id=?", task.Id)

		if err != nil {
			task.Result.Errno = ERROR_FORM
			task.Result.Errmsg = err.Error()
			return nil
		}

	} else {
		task.Result.Errno = ERROR_FORM_NOT_FOUND
		task.Result.Errmsg = "Not Found form"
		return nil
	}

	task.Result.Form = &v

	return nil
}

func (S *FormService) HandleFormQueryTask(a IFormApp, task *FormQueryTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FORM
		task.Result.Errmsg = err.Error()
		return nil
	}

	var forms = []Form{}

	var args = []interface{}{}

	var sql = bytes.NewBuffer(nil)

	sql.WriteString(" WHERE 1")

	if task.Id != 0 {
		sql.WriteString(" AND id=?")
		args = append(args, task.Id)
	}

	if task.Name != "" {
		sql.WriteString(" AND name=?")
		args = append(args, task.Name)
	}

	if task.Keyword != "" {
		q := "%" + task.Keyword + "%"
		sql.WriteString(" AND (title LIKE ? OR summary LIKE ? OR tags LIKE ?)")
		args = append(args, q, q, q)
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
		var counter = FormQueryCounter{}
		counter.PageIndex = pageIndex
		counter.PageSize = pageSize
		counter.RowCount, err = kk.DBQueryCount(db, a.GetFormTable(), a.GetPrefix(), sql.String(), args...)
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

	var v = Form{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := db.Query(fmt.Sprintf("SELECT id,name,title,type,summary,ctime FROM %s%s %s", a.GetPrefix(), a.GetFormTable().Name, sql.String()), args...)

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

		forms = append(forms, v)
	}

	task.Result.Forms = forms

	return nil
}
