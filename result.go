package bulk

import "database/sql"

type result struct {
	lastInsertId int64
	rowsAffected int64
}

func (r result) LastInsertId() (int64, error) {
	return r.lastInsertId, nil
}

func (r result) RowsAffected() (int64, error) {
	return r.rowsAffected, nil
}

func (r *result) Add(res sql.Result) {
	id, _ := res.LastInsertId()
	r.lastInsertId = id
	rows, _ := res.RowsAffected()
	r.rowsAffected += rows
}
