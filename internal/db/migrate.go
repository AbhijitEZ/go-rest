package db

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (db *Database) MigrateDB() error {
	fmt.Println("migrating the database ↻")
	driver, err := postgres.WithInstance(db.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create PG driver: %w", err)
	}

	appPath, _ := filepath.Abs("")
	appPath = appPath + "/migrations"

	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+appPath,
		"postgres",
		driver,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("could not run up migrations %w", err)
	}

	fmt.Println("migration ran successfully ✅")
	return nil
}
