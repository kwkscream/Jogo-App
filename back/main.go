// i dont rly know what ive done here
// 🦆🦆 🌱🌱 🐿️🐿️ 👀👀

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"log"
  "os"
)

type QueryParams struct {
	Q string `form:"q" binding:"required"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Result struct {
	BusinessStatus   string       `json:"business_status"`
	FormattedAddress string       `json:"formatted_address"`
	Geometry         Geometry     `json:"geometry"`
	Icon             string       `json:"icon"`
	Name             string       `json:"name"`
	OpeningHours     OpeningHours `json:"opening_hours"`
	Photos           []Photo      `json:"photos"`
	PlaceID          string       `json:"place_id"`
	PlusCode         PlusCode     `json:"plus_code"`
	Rating           float64      `json:"rating"`
	Reference        string       `json:"reference"`
	Types            []string     `json:"types"`
	UserRatingsTotal int          `json:"user_ratings_total"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type OpeningHours struct {
	OpenNow bool `json:"open_now"`
}

type Photo struct {
	Height           int      `json:"height"`
	HtmlAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int      `json:"width"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

type ApiResponse struct {
	HtmlAttributions []interface{} `json:"html_attributions"`
	Results          []Result      `json:"results"`
	Status           string        `json:"status"`
}

func getPlaces(q string) (*ApiResponse, error) {
	API_KEY := os.Getenv("API_KEY")
	client := resty.New()
	res, err := client.R().
		SetHeader("Accept", "application/json").
		SetResult(&ApiResponse{}).
		Get("https://maps.googleapis.com/maps/api/place/textsearch/json?query=" + q + "&key=" + API_KEY)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	responseData := res.Result().(*ApiResponse)
	return responseData, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	router.GET("/search", func(c *gin.Context) {
		var query QueryParams
		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		places, placesErr := getPlaces(query.Q)

		if placesErr != nil {
			c.JSON(500, gin.H{
				"error": "fuckinbg hell",
			})
		}

		c.JSON(200, places.Results)

	})

	router.Run()
}
