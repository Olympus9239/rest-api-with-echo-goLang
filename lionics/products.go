package lionics

import (
	"net/http"
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

var products = []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptop"}}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
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
}

func deleteProduct(c echo.Context) error {
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
		// [1,2,3,4,5,6]
		// suppose u deleted 3 and want to show the rest in database
		// [1,2] + [4,5,6] = [1,,2,4,5,6] (idiomatic in go,inefficient)
	}
	products = splice(products, index)
	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
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
}

func createProduct(c echo.Context) error {
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
}
