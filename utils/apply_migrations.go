package utils

import (
	"database/sql"

	"github.com/achintya-7/dating-app/config"
	"github.com/achintya-7/dating-app/logger"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func ApplyMigrations() {
	db, err := sql.Open("mysql", config.Values.MySqlUrl)
	if err != nil {
		logger.Fatal(nil, "Error opening database: ", err.Error())
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logger.Fatal(nil, "Error creating driver: ", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/sql/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		logger.Fatal(nil, "Error creating migration instance: ", err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Fatal(nil, "Error applying migrations: ", err.Error())
	}

	logger.Info(nil, "Migrations applied successfully")
}
