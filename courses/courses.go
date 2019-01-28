package courses

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetCourses() string {
	var result string
	current_date := time.Now().Local()
	current := current_date.Format("2006/01/02")
	// currency codes 978 - EUR, 840 - USD
	curr_codes := make(map[string]string)
	curr_codes["978"] = "EUR"
	curr_codes["840"] = "USD"

	for key, val := range curr_codes {
		url := "http://cbrates.rbc.ru/tsv/" + key + "/" + current + ".tsv"
		resp, _ := http.Get(url)
		bytes, _ := ioutil.ReadAll(resp.Body)

		data := string(bytes)
		r := csv.NewReader(strings.NewReader(data))
		r.Comma = '\t'
		r.Comment = '#'

		records, err := r.Read()
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(val, records[1])
		result += val
		result += ":"
		result += records[1]
		result += "\n"

		resp.Body.Close()
	}
	log.Println(result)
	return result
}
