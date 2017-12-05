package main

import (
	"log"

	"github.com/wonderstream/twitch/aggregation"
	"github.com/wonderstream/twitch/aggregation/populate"
	"github.com/wonderstream/twitch/aggregation/service"
)

func main() {
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
		},
	}

	manager.Start(loader)
}
