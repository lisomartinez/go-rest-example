package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lisomartinez/go-rest-example/internal/database"
	"github.com/lisomartinez/go-rest-example/internal/logs"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

func main() {
	_ = logs.InitLogger()

	client := database.NewSqlClient("db:db@tcp(localhost:3306)/db")
	doMigrate(client, "db")
}

func doMigrate(client *database.MySqlClient, dbName string) {
	driver, _ := mysql.WithInstance(client.DB, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("current migrations version in %d", current)

	err = m.Migrate(migrationsScriptsVersion)

	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	}

}
