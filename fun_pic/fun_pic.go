package fun_pic

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var m = make(map[string]string)

// Get random key from map
func randIntMapKey() string {
	i := rand.Intn(len(m))
	for k := range m {
		if i == 0 {
			return k
		}
		i--
	}
	panic("never")
}

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
	// See if the kref attribute exists on the element
	s := []string{}
	href, exists := element.Attr("href")
	if exists {
		s = strings.Split(href, "/")
		if s[1] == "i" {
			m[s[2]] = s[1]
		}
		// fmt.Println(m)
	}
}

func GetFunPic() string {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	fun_pic_link := fmt.Sprintf("https://imgflip.com/?page=%d", int(rand.Intn(100)))
	// Create and modify HTTP request before sending
	// fmt.Println(fun_pic_link)
	request, err := http.NewRequest("GET", fun_pic_link, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Go simple web scrapper 0.0.0.1 Alfa")

	// Make request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// // Copy data from the response to standard output
	// _, err = io.Copy(os.Stdout, response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(processElement)
	fun_pic_link = fmt.Sprintf("Enjoy: <a href=\"https://i.imgflip.com/%s.jpg\">&#8205;</a>", randIntMapKey())
	return fun_pic_link
}
