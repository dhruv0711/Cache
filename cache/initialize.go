package cache

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dhruv0711/Cache-Nokia/db"
	"github.com/dhruv0711/Cache-Nokia/types"
	"github.com/streadway/amqp"
)

func InitializeCache() {

	err := db.FetchFromDB(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgs, close, err := subscribe("cachedata")
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	go func(c <-chan amqp.Delivery) {
		list := []types.Todo{}
		var workerChan <-chan amqp.Delivery
		workerChan = c
		for v := range workerChan {
			err := json.Unmarshal(v.Body, &list)
			if err != nil {
				panic(err)
			}
			updateCache(list)
		}

		fmt.Println("cache updated successfully")
		fmt.Println("current cache status:", Cache)
	}(msgs)
}
