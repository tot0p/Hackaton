package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var Dataset []struct {
	ID string `json:"dataset_id"`
}

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
	// Open allDatasets.json file
	file, err := os.OpenFile("scraping/allDatasets.json", os.O_RDONLY, 0744)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)

	// Unmarshal json data
	err = json.Unmarshal(content, &Dataset)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create json file for each dataset
	/*
		for _, datasetID := range Dataset {
			// GET datas from dataset
			GetDataFromDataset(datasetID.ID)
			// Create map with theme as key and names as value
		}

	*/
	CreateJSONTheme()
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

type DatasetTheme struct {
	Name  string `json:"dataset_id"`
	Metas struct {
		Default struct {
			Theme []string `json:"theme"`
		} `json:"default"`
	} `json:"metas"`
}

// CreateJSONTheme create json file which link each dataset to a theme
func CreateJSONTheme() {
	var data []DatasetTheme

	body, err := os.ReadFile("scraping/allDatasets.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	// Clear json file
	err = os.Truncate("scraping/DatasetsTheme.json", 0)
	if err != nil {
		panic(err)
	}

	// Open json file
	file, err := os.OpenFile("scraping/DatasetsTheme.json", os.O_CREATE|os.O_WRONLY, 0744)
	if err != nil {
		panic(err)
	}

	// Write data to json file
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		panic(err)
	}

	file.Close()

	// Read the JSON file
	body, err = os.ReadFile("scraping/DatasetsTheme.json")

	// Unmarshal the JSON into an array of struct
	var dataStruct []DatasetTheme
	err = json.Unmarshal(body, &dataStruct)
	if err != nil {
		panic(err)
	}

	// Create a map to store themes and associated names
	themeMap := make(map[string][]string)

	// Iterate through the dataStruct and fill the map
	for _, item := range dataStruct {
		datasetID := item.Name
		themes := item.Metas.Default.Theme
		for _, theme := range themes {
			themeMap[theme] = append(themeMap[theme], datasetID)
		}
	}

	// Create the final result as a slice of maps
	var result []map[string]interface{}
	for theme, datasetIDs := range themeMap {
		resultItem := map[string]interface{}{
			"theme": theme,
			"name":  datasetIDs,
		}
		result = append(result, resultItem)
	}

	// Marshal the result back to JSON
	body, err = json.Marshal(result)
	if err != nil {
		panic(err)
	}

	// Write json file formatted
	err = os.WriteFile("scraping/DatasetThemeFormatted.json", body, 0744)
	if err != nil {
		panic(err)
	}
}
