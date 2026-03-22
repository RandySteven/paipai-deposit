// Package mysql_client provides a MySQL database client wrapper with connection
// pooling, health checks, and lifecycle management. If mysql.host is empty,
// it falls back to PostgreSQL using configs.postgres (lib/pq).
package mysql_client

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/RandySteven/paipai-deposit/configs"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type (
	// MySQL defines the interface for database operations including connection
	// management, health checks, and migrations.
	MySQL interface {
		// Close closes the database connection and releases resources.
		Close()
		// Ping verifies the database connection is still alive.
		Ping() error
		// Client returns the underlying *sql.DB instance for direct access.
		Client() *sql.DB
		// Migration runs database migrations within the given context.
		Migration(ctx context.Context) error
	}
	// mysqlClient is the internal implementation of the MySQL interface.
	mysqlClient struct {
		db *sql.DB
	}
)

// Client returns the underlying *sql.DB instance for direct database access.
func (m *mysqlClient) Client() *sql.DB {
	return m.db
}

// Close closes the database connection and releases all associated resources.
func (m *mysqlClient) Close() {
	m.db.Close()
}

// Ping verifies the database connection is still alive by sending a ping request.
func (m *mysqlClient) Ping() error {
	return m.db.Ping()
}

func dbDriverAndDSN(config *configs.Config) (driver string, dsn string) {
	m := config.Configs.Mysql
	if m.Host != "" {
		port := m.Port
		if port == "" {
			port = "3306"
		}
		return "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=UTC&charset=utf8mb4",
			m.DbUser, m.DbPass, m.Host, port, m.DbName)
	}
	p := config.Configs.Postgres
	port := p.Port
	if port == "" {
		port = "5432"
	}
	return "postgres", fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		p.DbUser, p.DbPass, p.Host, port, p.DbName)
}

// NewMYSQLClient creates a MySQL client when configs.mysql.host is set; otherwise
// a PostgreSQL client from configs.postgres. Connection pool: MaxIdleConns 10,
// MaxOpenConns 8, ConnMaxLifetime 10m, ConnMaxIdleTime 8m.
func NewMYSQLClient(config *configs.Config) (*mysqlClient, error) {
	driver, conn := dbDriverAndDSN(config)
	log.Println(driver+":", conn)
	db, err := sql.Open(driver, conn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(8)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(8 * time.Minute)
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return &mysqlClient{
		db: db,
	}, nil
}

var _ MySQL = &mysqlClient{}
