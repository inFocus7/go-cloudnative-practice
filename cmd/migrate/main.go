package main

import (
	"fabiangonz98/go-cloudnative-practice/adapter/db"
	"fabiangonz98/go-cloudnative-practice/config"
	"flag"
	"fmt"
	"github.com/pressly/goose"
	"log"
	"os"
)

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String("dir", "/go-cloudnative-practice/migrations", "directory with migration files")
)

func main() {
	flags.Usage = usage
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("parsing error: %v", err)
	}

	args := flags.Args()
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}

	command := args[0]
	switch command {
	case "create":
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	}

	appConf := config.AppConfig()
	appDb, err := db.New(appConf)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer appDb.Close()

	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Run(command, appDb, *dir, args[1:]...); err != nil {
		log.Fatalf("migrate run: %v", err)
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix = `Usage: migrate [OPTIONS] COMMAND
Examples:
  migrate status
Options:`
	usageCommands = `
Commands:
  up					Migrate the DB to the most recent version available
  up-by-one     		Migrate the DB up by 1
  up-to VERSION	    	Migrate the DB to a specific VERSION
  down					Roll back the version by 1
  down-to VERSION		Roll back the DB to a specific version
  redo					Re-run latest migration
  reset					Roll back all migrations
  status				Dump the migration status for the current DB
  version				Print current version of the DB
  create NAME [sql|go]	Creates new migration file with the current timestamp
  fix					Apply sequential ordering to migrations`
)
