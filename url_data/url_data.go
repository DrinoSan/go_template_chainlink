package url_data

import (
	"log"
	"net/http"
	"net/url"
	"os"

	respons_structures "github.com/drinosan/go_template_chainlink/response"
	"github.com/joho/godotenv"
)

var Basic_url string = "https://api.openweathermap.org/data/2.5/"
var Endpoint_url string = "weather"

func Get_url() *url.URL {

	url, err := url.Parse(Basic_url)
	if err != nil {
		log.Fatal(err)
	}
	rel, err := url.Parse(Endpoint_url)
	if err != nil {
		log.Fatal(err)
	}

	return rel
}

func Query_params(query_paramter *respons_structures.Params) *http.Request {

	// loading .env file.
	// content of .env:
	// appid=YOUR_API_KEY
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api_url := Get_url()

	req, err := http.NewRequest("GET", api_url.String(), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("q", query_paramter.Data.Q)
	q.Add("appid", os.Getenv("appid"))
	//q.Add("units", "metric")
	req.URL.RawQuery = q.Encode()

	return req
}

func Make_api_call(req *http.Request) *http.Response {

	r, err := http.Get(req.URL.String())
	if err != nil {
		log.Fatal("API CALL FAILD: ", err)
	}

	return r
}
