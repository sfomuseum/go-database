package sql

import (
	"context"
	"testing"
)

func TestDriver(t *testing.T) {

	ctx := context.Background()

	db, err := OpenWithURI(ctx, "sql://null?dsn=null")

	if err != nil {
		t.Fatalf("Failed to open null database, %v", err)
	}

	dr := Driver(db)

	if dr != NULL_DRIVER {
		t.Fatalf("Unexpected driver, %s", dr)
	}

}
