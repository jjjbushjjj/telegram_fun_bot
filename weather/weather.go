package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	base_url = "https://api.openweathermap.org/data/2.5/weather"
	apikey   = "2d50974864511c9e4285d94c881095ff"
	town     = "Moscow"
	units    = "metric"
	lang     = "en"
)

type Weather struct {
	Town string `json:"name"`
	Data W_main `json:"main"`
}

type W_main struct {
	Humidity int     `json:"humidity"`
	Pressure int     `json:"pressure"`
	Temp     float32 `json:"temp"`
	Temp_max float32 `json:"temp_max"`
	Temp_min float32 `json:"temp_min"`
}

func Get_weather() string {

	// Init Weather struct Default params
	w_api_req := fmt.Sprintf("%s?q=%s&units=%s&lang=%s&APPID=%s",
		base_url, town, units, lang, apikey)

	fmt.Println(w_api_req)
	response, err := http.Get(w_api_req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	var weather Weather
	json.Unmarshal(responseData, &weather)

	// fmt.Println(weather.Town)
	// fmt.Println(weather.Data.Temp)

	// // Copy data from the response to standard output
	// _, err = io.Copy(os.Stdout, response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return fmt.Sprintf("Place: %s\n Temperature: %.1f C\n Pressure: %d hPa\n Humidity: %d%%",
		weather.Town, weather.Data.Temp, weather.Data.Pressure, weather.Data.Humidity)

}
