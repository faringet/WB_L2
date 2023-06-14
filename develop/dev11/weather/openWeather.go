package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Location struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	State   string `json:"state"`
	Country string `json:"country"`
	Coord   struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
}

func ReturnMap() map[int]string {

	// Read the JSON file
	data, err := ioutil.ReadFile("city.list.json")
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal the JSON data
	var locations []Location
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println(err)
	}

	// Create a map to store the values
	locationMap := make(map[int]string)
	for _, location := range locations {
		locationMap[location.ID] = location.Name
	}

	return locationMap
}

func GetCityID(m map[int]string, val string) int {
	for k, v := range m {
		if v == val {
			return k
		}
	}
	// Return an empty string if the value is not present
	return 0
}

func GetTemperature(c int) float64 {
	apiKey := os.Getenv("OW_KEY")

	cityID := strconv.Itoa(c)

	// Make the API request
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?id=" + cityID + "&units=metric&appid=" + apiKey)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal the JSON response into a map
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
	// Type switch
	i := data["main"].(map[string]interface{})["temp"]
	f, ok := i.(float64)
	if ok {
	} else {
		fmt.Println("i is not a float64")
	}

	return f
}
