package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
)

type configDatabase struct {
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     string `env:"DB_PORT" env-required:"true"`
	Name     string `env:"DB_NAME" env-required:"true"`
	User     string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}

type argsCLI struct {
	MigrationType string
	NumberToApply int
}

func main() {
	config, err := readConfig()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("config = %+v\n", config)

	args, err := readArgs()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("args = %+v\n", args)

	err = runMigrations(config, args)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Migrations successfully applied")
}

func readConfig() (configDatabase, error) {
	const PathToEnvFile = "./"
	path := fmt.Sprintf("%v.env", PathToEnvFile)

	var config configDatabase
	err := cleanenv.ReadConfig(path, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func readArgs() (argsCLI, error) {
	args := os.Args

	if len(args) <= 1 {
		panic("got empty arguments list. Want command line arguments: up|down [N]\n")
	}

	migrationType := args[1]
	if !(migrationType == "up" || migrationType == "down") {
		panic(fmt.Sprintf("migrationType = %q; Allowed values = up|down\n", migrationType))
	}

	var migNum_to_apply int
	var err error
	if len(args) >= 3 {
		migNum_to_apply, err = strconv.Atoi(args[2])
		if err != nil {
			panic(fmt.Sprintf("Number of migrations to apply cannot be converted to int. Err = %v\n", err.Error()))
		}
		if migNum_to_apply <= 0 {
			panic("Number of migrations to apply should be possitive or omitted.\n")
		}
	}
	migNum_to_apply = max(migNum_to_apply, 1)

	resultArgs := argsCLI{MigrationType: migrationType, NumberToApply: migNum_to_apply}
	return resultArgs, nil
}

func runMigrations(configDB configDatabase, args argsCLI) error {
	db, err := sql.Open("postgres", buildDB_URL(configDB))
	if err != nil {
		return err
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		configDB.Name,
		driver,
	)
	if err != nil {
		return err
	}

	steps := args.NumberToApply
	if args.MigrationType == "down" {
		steps = -steps
	}

	err = m.Steps(steps)
	if err != nil {
		return err
	}

	return nil
}

func buildDB_URL(configDB configDatabase) string {
	dbURL := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		configDB.User,
		configDB.Password,
		configDB.Host,
		configDB.Port,
		configDB.Name,
	)
	return dbURL
}
