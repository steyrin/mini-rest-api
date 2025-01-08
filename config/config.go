package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	DatabaseURL string
}

func InitDB() *bun.DB {
	dsn := "postgres://admin:admin@localhost:5432/rest_api_db?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	//db := bun.NewDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)), nil)

	if err := db.Ping(); err != nil {
		logrus.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	logrus.Info("Подключение к базе данных успешно")

	return db
}
