package sql

import (
	"context"
	db_sql "database/sql"
	"fmt"
	"log/slog"
)

// LoadDuckDBExtensions will issue 'INSTALL' and 'LOAD' statements for 'extensions' using 'db'.
func LoadDuckDBExtensions(ctx context.Context, db *db_sql.DB, extensions ...string) error {

	for _, ext := range extensions {

		commands := []string{
			fmt.Sprintf("INSTALL %s", ext),
			fmt.Sprintf("LOAD %s", ext),
		}

		for _, cmd := range commands {

			_, err := db.ExecContext(ctx, cmd)

			if err != nil {
				return fmt.Errorf("Failed to issue command for extension '%s', %w", cmd, err)
			}
		}
	}

	return nil
}

func ConfigureDuckDBDatabase(ctx context.Context, db *db_sql.DB, opts *ConfigureDatabaseOptions) error {

	if opts.CreateTablesIfNecessary {

		for _, t := range opts.Tables {

			logger := slog.Default()
			logger = logger.With("table", t.Name())

			exists, err := HasDuckDBTable(ctx, db, t.Name())

			if err != nil {
				return fmt.Errorf("Failed to determine if table %s exists, %w", t.Name(), err)
			}

			if exists {
				logger.Debug("Table already exists")
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

			logger.Debug("Created table")
		}
	}

	return nil
}

func HasDuckDBTable(ctx context.Context, db *db_sql.DB, table_name string) (bool, error) {

	logger := slog.Default()
	logger = logger.With("table", table_name)

	q := "SELECT EXISTS(SELECT * FROM pg_tables WHERE schemaname='public' AND tablename=$1)"

	row := db.QueryRowContext(ctx, q, table_name)

	var exists bool

	err := row.Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("Failed to query table, %w", err)
	}

	logger.Debug("Does table exist", "exists", exists)
	return exists, nil
}
