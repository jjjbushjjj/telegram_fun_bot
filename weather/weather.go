package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

// Theese are default args
const (
	apikey       = "2d50974864511c9e4285d94c881095ff"
	base_url     = "https://api.openweathermap.org/data/2.5/weather"
	hourly_5_day = "https://api.openweathermap.org/data/2.5/forecast"
	daily_16_day = "https://api.openweathermap.org/data/2.5/forecast/daily"
	daily_days   = "16"
	town         = "Moscow"
	units        = "metric"
	lang         = "en"
)

type Weather struct {
	Town         string   `json:"name"`
	Weather_desc []W_desc `json:"weather"`
	Data         W_main   `json:"main"`
	Wind         W_wind   `json:"wind"`
	Rain         W_rain   `json:"rain"`
	Snow         W_snow   `json:"snow"`
	Clouds       W_cloud  `json:"clouds"`
	Dt           int64    `json:"dt"` // timestamp
}

type W_desc struct {
	Id   int    `json:"id"`
	Main string `json:"main"`
	Desc string `json:"description"`
	Icon string `json:"icon"`
}

type W_main struct {
	Humidity int     `json:"humidity"`
	Pressure float32 `json:"pressure"`
	Temp     float32 `json:"temp"`
	Temp_max float32 `json:"temp_max"`
	Temp_min float32 `json:"temp_min"`
}

type W_wind struct {
	Speed float32 `json:"speed"`
	Deg   int     `json:"deg"`
}

type W_rain struct {
	Rain_3 int `json:"3h"` // in mm
	Rain_1 int `json:"1h"`
}

type W_snow struct {
	Snow_3 int `json:"3h"` // in mm
	Snow_1 int `json:1h"`
}

type W_cloud struct {
	All int `json:"all"` // in %
}

type City struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type List struct {
	Dt           int64    `json:"dt"`   // Time in Unix format
	Data         W_main   `json:"main"` // Weather data
	Wind         W_wind   `json:"wind"`
	Clouds       W_cloud  `json:"clouds"`
	Weather_desc []W_desc `json:"weather"`
}

type List_16 struct {
	Dt           int64    `json:"dt"` // Time in Unix format
	Temp         Temp     `json:"temp"`
	Weather_desc []W_desc `json:"weather"`
}

type Temp struct {
	Day   float32 `json:"day"`
	Min   float32 `json:"min"`
	Max   float32 `json:"max"`
	Night float32 `json:"night"`
	Eve   float32 `json:"eve"`
	Morn  float32 `json:"morn"`
}

type Forecast struct {
	City City   `json:"city"`
	List []List `json:"list"`
}

type Forecast_16 struct {
	City City      `json:"city"`
	List []List_16 `json:"list"`
}

func formatAsDate(t int64) time.Time {
	// Convert from inix timestamp to human readable format
	return time.Unix(t, 0)
}

func send(w_api_req string) []byte {

	// fmt.Println(w_api_req)
	response, err := http.Get(w_api_req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	return responseData

}

func Get_weather() string {
	// Get current weather

	// Init Weather struct Default params
	w_api_req := fmt.Sprintf("%s?q=%s&units=%s&lang=%s&APPID=%s",
		base_url, town, units, lang, apikey)
	resp := send(w_api_req)
	var weather Weather
	json.Unmarshal(resp, &weather)
	// Format using template
	t, err := template.ParseFiles("weather.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(os.Stdout, weather)
	return ""
}

func Get_forecast_5() string {
	// Get 5 day / 3 hour forecast data

	// Init Weather struct Default params
	w_api_req := fmt.Sprintf("%s?q=%s&units=%s&lang=%s&APPID=%s",
		hourly_5_day, town, units, lang, apikey)
	resp := send(w_api_req)
	var forecast Forecast
	json.Unmarshal(resp, &forecast)

	// Format using template
	fmap := template.FuncMap{
		"formatAsDate": formatAsDate,
	}
	t, err := template.New("forecast_5.tmpl").Funcs(fmap).ParseFiles("forecast_5.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(os.Stdout, forecast)
	return ""
}

// This REQUIRED Paid subscription!!! So it is not tested (I don't have paid sub)
func Get_forecast_16(days string) string {
	// Get 16 day / daily forecast data default or
	// 1-16 day /daily forecast data if days param is specified
	if days == "" {
		days = daily_days
	}

	// Init Weather struct Default params
	w_api_req := fmt.Sprintf("%s?q=%s&units=%s&lang=%s&APPID=%s&cnt=%s",
		daily_16_day, town, units, lang, apikey, days)
	resp := send(w_api_req)
	var forecast Forecast_16
	json.Unmarshal(resp, &forecast)
	// Format using template
	fmap := template.FuncMap{
		"formatAsDate": formatAsDate,
	}
	t, err := template.New("forecast_16.tmpl").Funcs(fmap).ParseFiles("forecast_16.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(os.Stdout, forecast)
	return ""
	return ""
}
