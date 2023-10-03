package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Scraping get all datasets from opendata.paris.fr
func main() {
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

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Clear file
		err = os.Truncate("scraping/allDatasets.json", 0)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Write data to json file
		err = os.WriteFile("scraping/allDatasets.json", body, 0744)
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

	// Create json file for each dataset
	for num, datasetID := range data {
		// GET datas from dataset
		if num == 0 {
			continue
		} else {
			GetDataFromDataset(datasetID.DatasetID)
			fmt.Println("Num : ", num, " DatasetID : ", datasetID.DatasetID)
		}
	}
}

// GetDataFromDataset get data from dataset api request
func GetDataFromDataset(name string) {
	// Create json file
	file, err := os.OpenFile("scraping/datasets/"+name+".json", os.O_CREATE|os.O_WRONLY, 0744)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	// API URL to get dataset
	apiURL := "https://opendata.paris.fr/api/explore/v2.1/catalog/datasets/" + name + "/exports/json"
	// GET Request
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Check status code
	if response.StatusCode == http.StatusOK {
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
}
