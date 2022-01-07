package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	respons_structures "github.com/drinosan/go_template_chainlink/response"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	// curl -X POST -H "content-type:application/json" "http://localhost:8080/" --data '{ "id": 0, "data": { "q": "New York" }}'
	// { "id": 0, "data": { "q": "New York" }}

	// This handler will match /user/john but will not match /user/ or /user
	router.POST("/", func(c *gin.Context) {
		api_url := "https://api.openweathermap.org/data/2.5/weather"
		u, _ := url.ParseRequestURI(api_url)

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		log.Println("NACH NEW REQUEST")

		var newSearch respons_structures.Params
		if err := c.BindJSON(&newSearch); err != nil {
			log.Fatal(err)
		}
		log.Println("NACH BIND")

		newS, _ := json.MarshalIndent(&newSearch, "", "    ")
		log.Println("Recieved to querry: ", string(newS))

		q := req.URL.Query()
		q.Add("q", newSearch.Data.Q)
		q.Add("appid", os.Getenv("appid"))
		//q.Add("units", "metric")
		req.URL.RawQuery = q.Encode()

		r, err := http.Get(req.URL.String())
		if err != nil {
			return
		}

		defer r.Body.Close()
		w := new(respons_structures.OpenWeatherResponse)

		json.NewDecoder(r.Body).Decode(w)

		if err != nil {
			log.Panic(err)
		}

		log.Println("TEMP: ", w.Main.Temp)
		c.IndentedJSON(200, gin.H{
			"jobRunID":   newSearch.ID,
			"data":       w,
			"result":     w.Main.Temp,
			"statusCode": http.StatusOK,
		})

	})

	// router.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// })

	// // However, this one will match /user/john/ and also /user/john/send
	// // If no other routers match /user/john, it will redirect to /user/john/
	// router.GET("/user/:name/*action", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	action := c.Param("action")
	// 	message := name + " is " + action
	// 	c.String(http.StatusOK, message)
	// })

	// // For each matched request Context will hold the route definition
	// router.POST("/user/:name/*action", func(c *gin.Context) {
	// 	b := c.FullPath() == "/user/:name/*action" // true
	// 	c.String(http.StatusOK, "%t", b)
	// })

	// // This handler will add a new router for /user/groups.
	// // Exact routes are resolved before param routes, regardless of the order they were defined.
	// // Routes starting with /user/groups are never interpreted as /user/:name/... routes
	// router.GET("/user/groups", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "The available groups are [...]")
	// })

	// // Query string parameters are parsed using the existing underlying request object.
	// // The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	// router.GET("/welcome", func(c *gin.Context) {
	// 	firstname := c.DefaultQuery("firstname", "Guest")
	// 	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	// 	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	// })

	// router.POST("/form_post", func(c *gin.Context) {
	// 	message := c.PostForm("message")
	// 	nick := c.DefaultPostForm("nick", "anonymous")

	// 	c.JSON(200, gin.H{
	// 		"status":  "posted",
	// 		"message": message,
	// 		"nick":    nick,
	// 	})
	// })

	router.Run(":8080")
}
