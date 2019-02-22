package weather

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
