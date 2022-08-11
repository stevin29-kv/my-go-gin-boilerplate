package database

import (
	"database/sql"
	cfg "employee-app/config"
	"flag"
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./database/migrations", "directory with migration files")
)

func PrepareDatabase() (*gorm.DB, error) {
	config := cfg.GetConfig()

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUsername,
		config.PostgresDB,
		config.PostgresPassword,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.PostgresSchema + ".",
			SingularTable: true,
		},
	})
	if err != nil {
		log.Printf("error connecting database - %v", err)
		return nil, err
	}
	log.Println("db connection success")

	err = autoMigrate(dbURL)
	if err != nil {
		log.Printf("error running migration")
	}
	return db, nil
}

func autoMigrate(dbURL string) error {
	config := cfg.GetConfig()
	dbURL = fmt.Sprintf("%s search_path=%s", dbURL, config.PostgresSchema)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, *dir); err != nil {
		return err
	}

	log.Println("db migration complete")
	return nil
}
