package weather

import (
	"bufio"
	"bytes"
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
	lang         = "ru"
)

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
	fmap := template.FuncMap{
		"formatAsDate": formatAsDate,
	}
	t, err := template.New("weather.tmpl").Funcs(fmap).ParseFiles("weather.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
	wr := bufio.NewWriter(&b)

	t.Execute(wr, weather)
	wr.Flush()
	return b.String()
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

	var b bytes.Buffer
	wr := bufio.NewWriter(&b)

	t.Execute(wr, forecast)
	wr.Flush()
	return b.String()
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
