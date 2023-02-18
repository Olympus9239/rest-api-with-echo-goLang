package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

//1 :-->Echo Starting Video 9,10
// func main() {
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		//	return c.String(200, "Well! Hello there!!")
// 		return c.String(http.StatusOK, "Well! Hello there!!")

//		})
//		e.Logger.Print("Server is running on port 8080")
//		e.Logger.Fatal(e.Start(":8080"))
//	}
func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	// procducts is a slice of map with integer key and string value
	products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptop"}}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Well! Hello there!!")
	})
	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string
		// range over products slice and get each map like {1:"mobiles"}
		for _, p := range products {
			for k := range p {

				// product id is string so we need to convert it to integer, pID is fetcjhe from path parameter,here id is path parameter
				pId, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pId == k {
					product = p
				}
			}
		}
		if product == nil {
			// status not found is 404
			return c.JSON(http.StatusNotFound, "product not found")
		}
		return c.JSON(http.StatusOK, product)
	})
	e.Logger.Print("Server is running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
