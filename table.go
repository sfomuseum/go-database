package database

import (
	"context"
	"database/sql"
	"regexp"
)

var re_mem *regexp.Regexp
var re_vfs *regexp.Regexp
var re_file *regexp.Regexp

func init() {
	re_mem = regexp.MustCompile(`^(file\:)?\:memory\:.*`)
	re_vfs = regexp.MustCompile(`^vfs:\.*`)
	re_file = regexp.MustCompile(`^file\:([^\?]+)(?:\?.*)?$`)
}

type Table interface {
	Name() string
	Schema() string
	InitializeTable(context.Context, *sql.DB) error
	IndexRecord(context.Context, *sql.DB, interface{}) error
}

func CreateTableIfNecessary(ctx context.Context, db *sql.DB, t Table) error {

	create := false

	has_table, err := HasTable(ctx, db, t.Name())

	if err != nil {
		return err
	}

	if !has_table {
		create = true
	}

	if create {

		sql := t.Schema()

		if err != nil {
			return err
		}

		_, err = db.ExecContext(ctx, sql)

		if err != nil {
			return err
		}

	}

	return nil
}
