package sql

import (
	"database/sql"
	"database/sql/driver"
	"io"
)

func init() {
	sql.Register("null", &nullDriver{})
}

// nullDriver implements driver.Driver.  It always succeeds and returns a
// connection that does nothing.
type nullDriver struct{}

func (d *nullDriver) Open(name string) (driver.Conn, error) {
	return &nullConn{}, nil
}

type nullConn struct{}

func (c *nullConn) Prepare(query string) (driver.Stmt, error) {
	return &nullStmt{}, nil
}

func (c *nullConn) Close() error {
	return nil
}

func (c *nullConn) Begin() (driver.Tx, error) {
	return &nullTx{}, nil
}

type nullStmt struct{}

func (s *nullStmt) Close() error {
	return nil
}

func (s *nullStmt) NumInput() int {
	return 0
}

func (s *nullStmt) Exec(args []driver.Value) (driver.Result, error) {
	return &nullResult{}, nil
}

func (s *nullStmt) Query(args []driver.Value) (driver.Rows, error) {
	return &nullRows{}, nil
}

type nullTx struct{}

func (t *nullTx) Commit() error {
	return nil
}

func (t *nullTx) Rollback() error {
	return nil
}

type nullRows struct{}

func (r *nullRows) Columns() []string {
	return []string{}
}

func (r *nullRows) Close() error {
	return nil
}

func (r *nullRows) Next(dest []driver.Value) error {
	// No rows – signal the end of the result set.
	return io.EOF
}

type nullResult struct{}

func (r *nullResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (r *nullResult) RowsAffected() (int64, error) {
	return 0, nil
}
