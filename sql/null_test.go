package sql

import (
	"context"
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
}
