package sql

import (
	"context"
	"database/sql"
	"testing"
)

func TestNullDriver(t *testing.T) {

	ctx := context.Background()

	db, err := OpenWithURI(ctx, "sql://null?dsn=null")

	if err != nil {
		t.Fatalf("Failed to open null database, %v", err)
	}

	_, err = db.ExecContext(ctx, "SELECT 1 FROM 1 WHERE 1=1")

	if err != nil {
		t.Fatalf("Null query failed, %v", err)
	}

	st, err := db.Prepare("SELECT col FROM table WHERE value = ?")

	if err != nil {
		t.Fatalf("Failed to prepare statement, %v", err)
	}

	row := st.QueryRowContext(ctx, 1)

	var col any

	err = row.Scan(&col)

	if err != nil && err != sql.ErrNoRows {
		t.Fatalf("Failed to execute query statement, %v", err)
	}

	err = st.Close()

	if err != nil {
		t.Fatalf("Failed to close statement, %v", err)
	}

	err = db.Close()

	if err != nil {
		t.Fatalf("Failed to close null database, %v", err)
	}
}
