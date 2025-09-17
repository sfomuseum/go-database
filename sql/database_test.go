package sql

import (
	_ "context"
	"net/url"
	"testing"
)

func TestOpenWithURIEngine(t *testing.T) {

	tests := map[string]string{
		"sql://mysql/?dsn={...}@tcp(sfomuseum-debug.hdhdhhdjs99.us-east-1.rds.amazonaws.com)/sfomueum_debug&tls=true": "mysql",
		"sql://null?dsn=null": "null",
	}

	for db_uri, expected := range tests {

		u, err := url.Parse(db_uri)

		if err != nil {
			t.Fatalf("Failed to parse db_uri, %v", err)
		}

		q := u.Query()

		engine := u.Host
		dsn := q.Get("dsn")

		if engine == "" {
			t.Fatal("Missing engine")
		}

		if dsn == "" {
			t.Fatal("Missing dsn")
		}

		if engine != expected {
			t.Fatalf("Invalid engine. Got '%s' but expected '%s'", engine, expected)
		}

	}
}
