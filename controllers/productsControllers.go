package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"guisantosalves/apigolang/models"
)

type Message struct {
	Message string `json:"message"`
}

// ok
func GetProducts(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "Application/json")

	prod := []models.Products{}

	// getting all the results
	models.DB.Find(&prod)

	return c.JSON(http.StatusOK, prod)
}

// ok
func GetProductsById(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "Application/json")

	id := c.Param("id")

	var prod models.Products

	models.DB.First(&prod, id)

	return c.JSON(http.StatusOK, prod)
}

// ok
func PostProducts(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "Application/json")

	newProduct := new(models.Products)
	if err := c.Bind(newProduct); err != nil {
		panic(err)
	}

	//fmt.Println(newProduct)
	models.DB.Save(&newProduct)

	return c.JSON(http.StatusOK, newProduct)
}

// ok
func UpdateProducts(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "Application/json")

	var prod models.Products

	id := c.Param("id")

	// get the specific row
	models.DB.First(&prod, id)

	// overwriting the prod
	if err := c.Bind(&prod); err != nil {
		panic(err)
	}

	// saving the prod overwrited
	models.DB.Save(&prod)

	return c.JSON(http.StatusOK, prod)
}

// ok
func DeleteProducts(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "Application/json")

	prod := models.Products{}

	id := c.Param("id")

	models.DB.Delete(&prod, id)

	data := Message{}
	jsonbody := []byte(`{"message":"Excluído com sucesso!!"}`)
	if err := json.Unmarshal(jsonbody, &data); err != nil {
		panic(err)
	}
	// jsonbody -> data

	return c.JSON(http.StatusOK, data)
}
