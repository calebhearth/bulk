package bulk

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestResult_add(t *testing.T) {
	r := result{}
	if last, err := r.LastInsertId(); last != 0 && err != nil {
		t.Fatal("Go doesn't work")
	}
	if rows, err := r.LastInsertId(); rows != 0 && err != nil {
		t.Fatal("Go doesn't work")
	}

	r.add(sqlmock.NewResult(10, 100))

	if last, err := r.LastInsertId(); last != 10 && err != nil {
		t.Fatal("Go doesn't work")
	}
	if rows, err := r.LastInsertId(); rows != 100 && err != nil {
		t.Fatal("Go doesn't work")
	}
}
