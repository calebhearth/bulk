package bulk

import (
	"database/sql/driver"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestInsert_Exec(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	result := sqlmock.NewResult(2, 2)
	prep := mock.ExpectPrepare(`INSERT INTO people \(id, name\) VALUES \(\$1::bigint, \$2\), \(\$3::bigint, \$4\)`)
	prep.ExpectExec().WithArgs(1, "Caleb", 2, "Jinzhu").WillReturnResult(result)

	rows := [][]driver.Value{
		{1, "Caleb"},
		{2, "Jinzhu"},
	}

	result, err = NewInsert(
		db,
		"INSERT INTO people (id, name) VALUES <values>",
		[]string{"bigint", ""},
	).Exec(rows)

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}

	if id, _ := result.LastInsertId(); id != 2 {
		t.Fatal("Expected LastInsertId to be 2")
	}

	if id, _ := result.RowsAffected(); id != 2 {
		t.Fatal("Expected RowsAffected to be 2")
	}
}
