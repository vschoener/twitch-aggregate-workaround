package main

import (
	"flag"
	"log"
	"os"

	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/migrate"
)

// Options of this binary
type Options struct {
	Install bool
	Drop    bool
	Update  bool
}

func checkArgs() *Options {
	options := &Options{}

	flag.BoolVar(&options.Install, "install", false, "Start install. It won't try anything if the table exists for example")
	flag.BoolVar(&options.Update, "update", false, "Update structure if possible")
	flag.BoolVar(&options.Drop, "drop", false, "Drop the table before creating it")
	flag.Parse()

	if options.Install == false && options.Update == false {
		flag.Usage()
		return nil
	}

	return options
}

func main() {
	options := checkArgs()
	if options == nil {
		os.Exit(0)
	}

	c := credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	if err := c.LoadSetting(); err != nil {
		log.Fatal(err)
	}
	l := logger.NewLogger()
	l.Connect(c.LoggerSettings)
	dm := storage.GetDM()
	dm.ConnectNewDatabase(storage.DBAggregation, c.GetDB(storage.DBAggregation), l.Share())
	defer dm.Get(storage.DBAggregation).Close()

	migrate := migrate.Migrate{
		DB: dm.Get(storage.DBAggregation),
		Options: migrate.Options{
			DropIfInstall: options.Drop,
			IsInstall:     options.Install,
			IsUpdate:      options.Update,
		},
	}

	migrate.Install()
}
