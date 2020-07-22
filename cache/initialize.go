package cache

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dhruv0711/Cache/db"
	"github.com/dhruv0711/Cache/types"
	"github.com/streadway/amqp"
)

//InitializeCache function to initialize cache on server startup
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
