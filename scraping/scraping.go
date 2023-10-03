package scraping

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Scraping get all datasets from opendata.paris.fr
func Scraping() {
	// API URL to get all datasets
	apiURL := "https://opendata.paris.fr/api/explore/v2.1/catalog/exports/json"

	// GET request
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Check status code
	if response.StatusCode == http.StatusOK {
		file, err := os.OpenFile("scraping/allDatasets.json", os.O_CREATE|os.O_WRONLY, 0744)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		defer file.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Write data to json file
		_, err = file.Write(body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

	} else {
		fmt.Printf("Request failed with status code: %d\n", response.StatusCode)
	}

	CreateAllDatasets()
}

// CreateAllDatasets create all datasets json files
func CreateAllDatasets() {
	var data []struct {
		DatasetID string `json:"dataset_id"`
	}

	// Open allDatasets.json file
	file, err := os.OpenFile("scraping/allDatasets.json", os.O_RDONLY, 0744)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)

	// Unmarshal json data
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(data[0].DatasetID)
}
