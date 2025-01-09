package config

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/steyrin/mini-rest-api/internal/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	DatabaseURL string
}

func CreateTables(db *bun.DB) {
	ctx := context.Background()

	// Создание таблицы для пользователей
	if _, err := db.NewCreateTable().
		Model((*model.User)(nil)).
		IfNotExists().
		Exec(ctx); err != nil {
		logrus.Fatalf("Ошибка при создании таблицы User: %v", err)
	}

	// Создание таблицы для книг
	if _, err := db.NewCreateTable().
		Model((*model.Book)(nil)).
		IfNotExists().
		Exec(ctx); err != nil {
		logrus.Fatalf("Ошибка при создании таблицы Book: %v", err)
	}

	// Создание таблицы для связи UserBook
	if _, err := db.NewCreateTable().
		Model((*model.UserBook)(nil)).
		IfNotExists().
		Exec(ctx); err != nil {
		logrus.Fatalf("Ошибка при создании таблицы UserBook: %v", err)
	}

	// Создание таблицы для отзывов
	if _, err := db.NewCreateTable().
		Model((*model.Review)(nil)).
		IfNotExists().
		Exec(ctx); err != nil {
		logrus.Fatalf("Ошибка при создании таблицы Review: %v", err)
	}

	logrus.Info("Таблицы успешно созданы или уже существуют")
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
	CreateTables(db)

	return db
}
