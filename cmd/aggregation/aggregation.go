package main

import (
	"flag"
	"log"
	"time"

	"github.com/wonderstream/twitch/aggregation"
	"github.com/wonderstream/twitch/aggregation/populate"
	"github.com/wonderstream/twitch/aggregation/precompute"
	"github.com/wonderstream/twitch/aggregation/service"
	"github.com/wonderstream/twitch/storage"
)

func checkArgs() storage.DateFilter {
	date := storage.DateFilter{}

	var dateStart string
	var dateEnd string
	flag.StringVar(&dateStart, "date_start", "", "Date range to start (format YYYY-MM-DD HH:mm:ss)")
	flag.StringVar(&dateEnd, "date_end", "", "Date range to end (format YYYY-MM-DD HH:mm:ss)")
	flag.Parse()

	var err error
	if len(dateStart) > 0 {
		date.DateStart, err = time.Parse(storage.SIMPLEFORMATSQL, dateStart)
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(dateEnd) > 0 {
		date.DateEnd, err = time.Parse(storage.SIMPLEFORMATSQL, dateEnd)
		if err != nil {
			log.Fatal(err)
		}
	}

	return date
}

func main() {
	date := checkArgs()
	log.Print("Initializing aggregation")
	loader := &service.Loader{}
	loader.Initialize()
	defer loader.Close()

	aggregationLogger := loader.Logger.Share()
	aggregationLogger.SetPrefix("AGGREGATION")
	manager := aggregation.AggregatorManager{
		Aggregators: []aggregation.Aggregator{
			&populate.Channel{},
			&populate.User{},
			// &populate.Stream{},
			&precompute.Channel{
				QueryFilter: storage.QueryFilter{
					DateFilter: date,
				},
			},
		},
	}

	manager.Start(loader)
}
