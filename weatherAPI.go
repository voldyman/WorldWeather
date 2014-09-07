package worldweather

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"appengine"
	"appengine/urlfetch"
)

type weatherAPIResponse struct {
	Name string `json: name`
	Main struct {
		Temp     float64 `json: "temp"`
		Pressure float64 `json: "pressure"`
		Humidity float64 `json: "humidity"`
		Temp_Min float64 `json: "temp_min"`
		Temp_Max float64 `json: "temp_max"`
	} `json: "main"`
	Weather []struct {
		Main        string `json: main`
		Description string
		Icon        string
	} `json: weather`
}

type WeatherResult struct {
	CityName string `json: 'cityName'`
	Values   struct {
		Temperature string `json: "Temperature"`
		Humidity    string `json: "Humidity"`
		Condition   string `json: "Condition"`
		TempMin     string `json: "MinTemperature"`
		TempMax     string `json: "MaxTemperature"`
	}
}

//Fetches data from openweathermap and returns the result as a json byte array
//requres a appengine.Context to use urlfetch
func FetchWeatherData(ctx appengine.Context, name string) (ret []byte, err error) {
	client := urlfetch.Client(ctx)
	resp, err := client.Get("http://api.openweathermap.org/data/2.5/weather?q=" + name)
	if err != nil {
		return
	}

	var data []byte
	data, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	var apiResponse weatherAPIResponse
	json.Unmarshal(data, &apiResponse)

	webResponse := &WeatherResult{
		CityName: apiResponse.Name,
	}

	webResponse.Values.Temperature = strconv.FormatFloat(apiResponse.Main.Temp-273, 'f', 1, 64)
	webResponse.Values.TempMin = strconv.FormatFloat(apiResponse.Main.Temp_Min-273, 'f', 1, 64)
	webResponse.Values.TempMax = strconv.FormatFloat(apiResponse.Main.Temp_Max-273, 'f', 1, 64)
	webResponse.Values.Humidity = strconv.FormatFloat(apiResponse.Main.Humidity, 'f', 1, 64) + "%"
	if len(apiResponse.Weather) > 0 {
		webResponse.Values.Condition = apiResponse.Weather[0].Main
	} else {
		webResponse.Values.Condition = "<null>"
	}

	ctx.Infof("%v", apiResponse)
	ret, err = json.Marshal(webResponse)
	if err != nil {
		return
	}

	return
}
