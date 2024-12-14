package sql

import (
	"context"
	"database/sql"
	"fmt"
)

func ConfigurePostgresDatabase(ctx context.Context, db *sql.DB, opts *ConfigureDatabaseOptions) error {

	if opts.CreateTablesIfNecessary {

		for _, t := range opts.Tables {

			exists, err := HasPostgresTable(ctx, db, t.Name())

			if err != nil {
				return fmt.Errorf("Failed to determine if table %s exists, %w", t.Name(), err)
			}

			if exists {
				continue
			}

			schema, err := t.Schema(db)

			if err != nil {
				return fmt.Errorf("Failed to derive schema for table %s, %w", t.Name(), err)
			}

			_, err = db.ExecContext(ctx, schema)

			if err != nil {
				return fmt.Errorf("Failed to create %s table, %w", t.Name(), err)
			}
		}
	}

	return nil
}

// https://stackoverflow.com/questions/20582500/how-to-check-if-a-table-exists-in-a-given-schema

func HasPostgresTable(ctx context.Context, db *sql.DB, table_name string) (bool, error) {

	sql := "SELECT EXISTS(SELECT * FROM pg_tables WHERE schemaname='public' AND tablename  = '?')"

	row := db.QueryRowContext(ctx, sql, table_name)

	var exists bool

	err := row.Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("Failed to query table, %w", err)
	}

	return exists, nil
}
