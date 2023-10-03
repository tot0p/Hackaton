package main

import (
	"encoding/json"
	"github.com/tot0p/env"
	"hackaton/utils/db/mongodb"
	"io"
	"os"
)

// Create collections and insert the dataset in it
func main() {
	// Load .env file
	err := env.LoadPath("./.env")
	if err != nil {
		panic(err)
	}

	// Connect to MongoDB
	m, err := mongodb.NewMongoDB(env.Get("URI_MONGODB"))
	if err != nil {
		panic(err)
	}
	defer m.Disconnect()

	err = m.Ping()
	if err != nil {
		panic(err)
	}

	// Get all datasets names
	allNames, err := os.ReadDir("scraping/datasets")
	if err != nil {
		panic(err)
	}

	for _, name := range allNames {
		DataSetName := name.Name()[:len(name.Name())-5]
		// Create collections if not exists
		collection := m.GetCollection("DataSets", DataSetName)
		if collection == nil {
			m.NewCollection("DataSets", DataSetName)
		}

		// Insert dataset in collection
		file, err := os.Open("scraping/datasets/" + name.Name())
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		var data []interface{}

		err = json.Unmarshal(body, &data)

		_, err = m.InsertMany("DataSets", DataSetName, data)
		if err != nil {
			panic(err)
		}
		file.Close()
	}
}
