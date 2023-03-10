package lionics

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var e = echo.New()
var v = validator.New()

//Start the application

func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.PUT("/products/:id", updateProduct)
	e.POST("/products", createProduct)

	e.Logger.Print("Server is running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))

}
