package main

import (
	"encoding/json"
	"log"
	"net/http"

	response_struct "github.com/drinosan/go_template_chainlink/response"
	"github.com/drinosan/go_template_chainlink/url_data"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// How to send the request
	// curl -X POST -H "content-type:application/json" "http://localhost:8080/" --data '{ "id": "0", "data": { "q": "New York", "units": "metric" }}'
	// { "id": 0, "data": { "q": "New York", "units": "metric" } }

	router.POST("/", func(c *gin.Context) {

		// Creating a struct variable of type response_struct.Params
		// which mirrors our querry parameters.
		// Our Querry Params are:
		// {
		// 	"id": "0",
		// 	"Data": {
		// 		"q": "New York"
		// 	}
		// }
		// curl -X POST -H "content-type:application/json" "http://localhost:8080/" --data '{ "id": 0, "data": { "q": "New York" }}'
		//
		//
		// type Params struct {
		// ID   string `json:"id"`
		// Data ParamQuery
		// }
		var searchParams response_struct.Params
		// We bind the values from the request body to the struct
		if err := c.BindJSON(&searchParams); err != nil {
			log.Fatal("Error while binding values", err)
		}

		// Only for logging
		newS, _ := json.MarshalIndent(&searchParams, "", "    ")
		log.Println("Recieved to querry: ", string(newS))

		// Creating querry parameters
		req := url_data.Query_params(&searchParams)

		// Finally making the call to the API
		r := url_data.Make_api_call(req)
		defer r.Body.Close()

		// We could read the body seperately and save it in a string
		// body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	log.Fatalln("Could not read Response Body: ", err)
		// }

		// OpenWeatherResponse is a struct which exactly mirrors the response
		// The struct itself was created with the help of https://mholt.github.io/json-to-go
		open_weather_respnse := new(response_struct.OpenWeatherResponse)

		// Writing the response in the struct
		json.NewDecoder(r.Body).Decode(open_weather_respnse)

		// Sending response to the Client, in our case the chainlink node
		// The chainlink node listens to this 4 fields
		log.Println("TEMP: ", open_weather_respnse.Main.Temp)
		c.IndentedJSON(200, gin.H{
			"jobRunID":   searchParams.ID,
			"data":       open_weather_respnse,
			"result":     open_weather_respnse.Main.Temp,
			"statusCode": http.StatusOK,
		})

	})

	router.Run(":8080")
}
