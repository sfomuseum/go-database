package database

type ConfigureSQLDatabaseOptions struct {
	CreateTablesIfNecessary bool
	Tables                  []Table
	Pragma                  []string
}

func DefaultConfigureSQLDatabaseOptions() *ConfigureSQLDatabaseOptions {
	opts := &ConfigureSQLDatabaseOptions{}
	return opts
}
