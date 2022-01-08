package main

import (
	"encoding/json"
	"log"
	"net/http"

	res_struct "github.com/drinosan/go_template_chainlink/response"
	"github.com/drinosan/go_template_chainlink/url_data"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// curl -X POST -H "content-type:application/json" "http://localhost:8080/" --data '{ "id": 0, "data": { "q": "New York" }}'
	// { "id": 0, "data": { "q": "New York" }}

	router.POST("/", func(c *gin.Context) {

		var newSearch res_struct.Params
		if err := c.BindJSON(&newSearch); err != nil {
			log.Fatal(err)
		}

		newS, _ := json.MarshalIndent(&newSearch, "", "    ")
		log.Println("Recieved to querry: ", string(newS))

		req := url_data.Query_params(&newSearch)

		r := url_data.Make_api_call(req)
		defer r.Body.Close()

		w := new(res_struct.OpenWeatherResponse)
		json.NewDecoder(r.Body).Decode(w)

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
