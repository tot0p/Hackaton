package main

import (
	"encoding/json"
	"fmt"
	"github.com/tot0p/env"
	"hackaton/utils/db/mongodb"
	"io"
	"os"
	"sync"
)

// Create collections and insert the dataset in it
func main() {
	// Load .env file
	err := env.LoadPath("./.env")
	if err != nil {
		panic(err)
	}

	// Connect to MongoDB
	err = mongodb.NewMongoDB(env.Get("URI_MONGODB"))
	if err != nil {
		panic(err)
	}

	err = mongodb.DB.Ping()
	if err != nil {
		panic(err)
	}

	// Get all datasets names
	allNames, err := os.ReadDir("scraping/datasets")
	if err != nil {
		panic(err)
	}

	// Create a channel to limit concurrency to 4 goroutines

	workerLimit := 2
	workerCh := make(chan struct{}, workerLimit)

	var wg sync.WaitGroup

	for _, name := range allNames {
		wg.Add(1)
		go func(name os.DirEntry) {
			defer wg.Done()

			// Acquire a worker slot
			workerCh <- struct{}{}
			defer func() { <-workerCh }()

			DataSetName := name.Name()[:len(name.Name())-5]
			// Create collections if not exists
			collection := mongodb.DB.GetCollection("DataSets", DataSetName)
			if collection == nil {
				mongodb.DB.NewCollection("DataSets", DataSetName)
			}

			// Insert dataset in collection
			file, err := os.Open("scraping/datasets/" + name.Name())
			if err != nil {
				panic(err)
			}
			defer file.Close()

			body, err := io.ReadAll(file)
			if err != nil {
				panic(err)
			}
			var data []interface{}

			err = json.Unmarshal(body, &data)
			if err != nil {
				panic(err)
			}

			if len(data) > 0 {
				_, err = mongodb.DB.InsertMany("DataSets", DataSetName, data)
				if err != nil {
					panic(err)
				}
				fmt.Println("Inserted", DataSetName)
			} else {
				fmt.Println("Empty", DataSetName)
			}
		}(name)
	}

	wg.Wait()

	err = mongodb.DB.Disconnect()
	if err != nil {
		panic(err)
	}
}
