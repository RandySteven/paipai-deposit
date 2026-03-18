package mysql_client

import (
	"context"

	"github.com/RandySteven/go-kopi/queries"
)

func registerMigration() []queries.TableMigration {
	return []queries.TableMigration{
		queries.MigrateUser,
	}
}

// Migration runs database migrations within the given context.
// Currently returns nil as migrations are not implemented.
func (m *mysqlClient) Migration(ctx context.Context) error {
	for _, migration := range registerMigration() {
		_, err := m.db.ExecContext(ctx, migration.ToString())
		if err != nil {
			return err
		}
	}
	return nil
}
