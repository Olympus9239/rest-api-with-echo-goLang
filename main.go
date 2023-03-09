package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ProdcutValidator  echo Validator for Product
type ProdcutValidator struct {
	validator *validator.Validate
}

// Validate validates product request body
func (p *ProdcutValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}
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
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		// range over products slice and get each map like {1:"mobiles"}
		for _, p := range products {
			for k := range p {

				if pID == k {
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
	// e.GET("/products/:vendor", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, c.QueryParam("olderThan"))
	// })
	// DELETE
	e.DELETE("/products/:id", func(c echo.Context) error {
		var product map[int]string
		var index int
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		// range over products slice and get each map like {1:"mobiles"}
		for i, p := range products {
			for k := range p {
				if pID == k {
					product = p
					index = i
				}
			}
		}
		if product == nil {
			// status not found is 404
			return c.JSON(http.StatusNotFound, "product not found")
		}
		splice := func(s []map[int]string, index int) []map[int]string {
			return append(s[:index], s[index+1:]...)
		}
		products = splice(products, index)
		return c.JSON(http.StatusOK, product)
	})
	e.PUT("/products/:id", func(c echo.Context) error {
		var product map[int]string
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for _, p := range products {
			for k := range p {
				if pID == k {
					product = p
				}
			}

		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found ")
		}
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}
		var reqBody body
		e.Validator = &ProdcutValidator{validator: v}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := c.Validate(reqBody); err != nil {
			return err
		}
		product[pID] = reqBody.Name
		return c.JSON(http.StatusOK, product)
	})

	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"` //validate added
		}
		var reqBody body
		//adding echo validator
		e.Validator = &ProdcutValidator{validator: v}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		//adding echoValidator
		if err := c.Validate(reqBody); err != nil {
			return err
		}
		// //adding validation
		// if err := v.Struct(reqBody); err != nil {
		// 	return err
		// }
		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, product)
		return c.JSON(http.StatusOK, product)
	})
	e.Logger.Print("Server is running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
