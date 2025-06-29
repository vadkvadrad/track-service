package db

import (
	"context"
    "fmt"

    "github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

type DbConfig struct {
	Dsn string
}

func NewDbConfig(dsn string) *DbConfig {
	return &DbConfig{
		Dsn: dsn,
	}
}


var dbPool *pgxpool.Pool

// InitDB инициализирует пул соединений с PostgreSQL
func InitDB(conf *DbConfig) *pgxpool.Pool {
    if dbPool != nil {
        return dbPool
    }

    config, err := pgxpool.ParseConfig(conf.Dsn)
    if err != nil {
        panic(fmt.Sprintf("Failed to parse DB config: %v", err))
    }

    pool, err := pgxpool.ConnectConfig(context.Background(), config)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to DB: %v", err))
    }

    // Проверка подключения
    if err := pool.Ping(context.Background()); err != nil {
        panic(fmt.Sprintf("Failed to ping DB: %v", err))
    }

    fmt.Println("Connected to database successfully")

    dbPool = pool
    return dbPool
}