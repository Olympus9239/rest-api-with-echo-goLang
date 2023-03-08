package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	v := validator.New()
	// procducts is a slice of map with integer key and string value
	products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptop"}}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Well! Hello there!!")
	})
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})
	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string
		// range over products slice and get each map like {1:"mobiles"}
		for _, p := range products {
			for k := range p {

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
	// e.GET("/products/:vendor", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, c.Param("vendor"))
	// })
	e.GET("/products/:vendor", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.QueryParam("olderThan"))
	})
	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := v.Struct(reqBody); err != nil {
			return err
		}
		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, product)
		return c.JSON(http.StatusOK, product)
	})
	e.Logger.Print("Server is running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
