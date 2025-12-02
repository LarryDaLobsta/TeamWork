package dal

import (
	"context"
	"database/sql"
    "fmt"
    "log"
    "os"
    "teamplayer/ent"
    _ "github.com/lib/pq"
)



type DBConnection struct {
	SQL		*sql.DB
	Ent		*ent.Client
}

// small helper for env vars with defaults
func getenv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}

func NewDbConnection(ctx context.Context) (*DBConnection, error) {
    // Eventually pull these from env
    host := getenv("PG_HOST", "localhost")
    port := getenv("PG_PORT", "5432")
    user := getenv("PG_USER", "postgres")
    pass := getenv("PG_PASSWORD", "postgres")
    name := getenv("PG_DBNAME", "todos")

	// build the connection string for postgres
	postgresConnStr := fmt.Sprintf(
        "postgresql://%s:%s@%s:%s/%s?sslmode=disable",
        user, pass, host, port, name,
    )

	// build the connection string for the ent framework
	entFrameConnStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, pass, name,
    )

	// connect to the database with postgres
	sqlDB, err := sql.Open("postgres", postgresConnStr)

	if err != nil {
		log.Fatalf("Failed opening connection to postgres: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
        return nil, fmt.Errorf("sql ping: %w", err)
    }

	// ent client connection
	entClient, err := ent.Open("postgres", entFrameConnStr)

	if err != nil {
        return nil, fmt.Errorf("ent open: %w", err)
    }

    // Auto-migrate
    if err := entClient.Schema.Create(ctx); err != nil {
        return nil, fmt.Errorf("ent schema: %w", err)
    }

	log.Println("DB connections established")

	return &DBConnection{
		SQL: sqlDB,
		Ent: entClient,
	}, nil
}

func (dbConn *DBConnection) DbHealthCheck(ctx context.Context) error {
	if err := dbConn.SQL.Ping(); err != nil {
		return fmt.Errorf("sql ping: %w", err)
	}

	// make sure that ent is doing its job
	_, err := dbConn.Ent.User.Query().Limit(1).All(ctx)
	if err != nil && !ent.IsNotFound(err){

		// 0 rows is fine to be returned - table exists and query succeeded
		return fmt.Errorf("ent query: %w", err)
	}

	return nil 
}