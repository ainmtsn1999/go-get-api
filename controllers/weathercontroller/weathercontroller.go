package weathercontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getWeather(loc string, API_KEY string) []interface{} {

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + loc + "&appid=" + API_KEY)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//fmt.Println(string(respData))

	var data map[string]interface{}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	datas := data["weather"].([]interface{})

	//fmt.Println(datas)

	return datas
}

func Index(c *gin.Context) {

	loc := c.Query("loc")
	API_KEY := "f7e80336d482f800cfbdba35d08e9539"

	//get weather
	data := getWeather(loc, API_KEY)

	c.JSON(http.StatusOK, gin.H{"weather": data})
}
