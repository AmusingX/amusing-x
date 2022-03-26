package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"database/sql"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/jmoiron/sqlx"
	"strings"
)

func InsetCategory(ctx context.Context, category *charon.Category) (*charon.Category, aerror.Error) {
	insertSql := `INSERT INTO category (parent_id, name, description) VALUES (:parent_id, :name, :description)`

	result, err := charon2.CharonDB.NamedExecContext(ctx, insertSql, category)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "get last insert id failed")
	}

	category.ID = id

	return category, nil
}

func CategoryQuery(ctx context.Context, req *charonservice.CategoryListRequest) ([]*charon.Category, int64, aerror.Error) {
	wherePlaceholder := `{{whereCondition}}`
	sqlStr := fmt.Sprintf(`SELECT id, name, description FROM category %s limit ?, ?`, wherePlaceholder)
	countStr := fmt.Sprintf(`SELECT count(*) FROM category  %s`, wherePlaceholder)

	var whereConditions []string
	var params []interface{}
	if req.Filter.Id > 0 {
		whereConditions = append(whereConditions, `id = ?`)
		params = append(params, req.Filter.Id)
	}

	if len(req.Filter.Name) != 0 {
		whereConditions = append(whereConditions, "name = ?")
		params = append(params, req.Filter.Name)
	}

	if len(req.Filter.Desc) != 0 {
		whereConditions = append(whereConditions, "description = ?")
		params = append(params, req.Filter.Desc)
	}

	params = append(params, req.Offset, req.Limit)

	if len(whereConditions) != 0 {
		whereSQl := "WHERE " + strings.Join(whereConditions, " AND ")
		sqlStr = strings.Replace(sqlStr, wherePlaceholder, whereSQl, -1)
		countStr = strings.Replace(countStr, wherePlaceholder, whereSQl, -1)
	} else {
		sqlStr = strings.Replace(sqlStr, wherePlaceholder, "", -1)
		countStr = strings.Replace(countStr, wherePlaceholder, "", -1)
	}

	tx, err := charon2.CharonDB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "begin transaction failed")
	}

	var count int64
	err = tx.QueryRowx(countStr, params[:len(params)-2]...).Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		count = 0
	case err != nil:
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "query count error")
	}

	var categories []*charon.Category
	err = tx.SelectContext(ctx, &categories, sqlStr, params...)
	if err != nil {
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select total category list failed")
	}

	return categories, count, nil
}

func Delete(ctx context.Context, ids []int64) aerror.Error {
	delSql := `DELETE FROM category WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = charon2.CharonDB.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del category failed")
	}

	return nil
}

func QueryCategoryByID(ctx context.Context, id int64) (*charon.Category, aerror.Error) {
	sqlStr := `SELECT id, name, description FROM category WHERE id = ?`

	var category charon.Category
	err := charon2.CharonDB.GetContext(ctx, &category, sqlStr, id)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select failed")
	}

	return &category, nil
}

func UpdateCategory(ctx context.Context, category *charon.Category) aerror.Error {
	sqlStr := `UPDATE category SET name = :name, description = :description WHERE id = :id`

	_, err := charon2.CharonDB.NamedExecContext(ctx, sqlStr, category)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func QueryCategoryByName(ctx context.Context, name string) (*charon.Category, aerror.Error) {
	sqlStr := `SELECT id, name, description FROM category WHERE name = ?`

	var category charon.Category
	err := charon2.CharonDB.GetContext(ctx, &category, sqlStr, name)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select failed")
	}

	return &category, nil
}
