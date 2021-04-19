package migrations

type pgClass struct {
	tableName struct{} `pg:"pg_catalog.pg_class"`
	Name      string   `pg:"relname,type:name"`
	Kind      rune     `pg:"relkind,type:char"`
}

// MigrateResponse is a standardized type representing what a service will
// return after a migration
type MigrateResponse struct {
	Messages []string `json:"messages"`
}
