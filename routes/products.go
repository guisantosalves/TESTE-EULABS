package routes

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"guisantosalves/apigolang/controllers"
)

var E = echo.New()

type response struct {
	Message string
}

func Routing() {

	jsonBody := []byte(`{"message":"Bem vindo a API de produtos do teste da EULABS"}`)
	data := response{}

	if err := json.Unmarshal(jsonBody, &data); err != nil {
		panic(err)
	}

	E.GET("/", func(c echo.Context) error {
		c.Response().Header().Set("content-type", "application/json")
		return c.JSON(http.StatusOK, data)
	})

	// methods
	E.GET("/products", controllers.GetProducts)
	E.GET("/products/:id", controllers.GetProductsById)
	E.POST("/products", controllers.PostProducts)
	E.PUT("/products/:id", controllers.UpdateProducts)
	E.DELETE("/products/:id", controllers.DeleteProducts)
}
