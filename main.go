package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID       string `json:id`
	Name     string `json:name`
	Price    int    `json:price`
	Quantity int    `json:quantity`
}

var products = []product{
	{ID: "1", Name: "Fan", Price: 200, Quantity: 2},
	{ID: "2", Name: "Sneaker", Price: 500, Quantity: 4},
	{ID: "3", Name: "Book", Price: 600, Quantity: 5},
}

func generateID() string {
	return strconv.Itoa(len(products) + 1)
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func ProductById(c *gin.Context) {
	id := c.Param("id")

	fmt.Print(id)
	for i, p := range products {
		if p.ID == id {
			fmt.Print(p)
			c.JSON(http.StatusOK, &products[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Product not found."})
}

func getProductById(id string) (*product, error) {
	for i, b := range products {
		if b.ID == id {
			return &products[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func createProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct.ID = generateID()

	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func checkoutProduct(c *gin.Context) {
	id, StatusOK := c.GetQuery("id")

	if !StatusOK {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	product, err := getProductById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found."})
		return
	}

	if product.Price <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not available"})
		return
	}

	product.Quantity -= 1
	c.IndentedJSON(http.StatusOK, product)

}

func main() {
	router := gin.Default()
	router.GET("/product", getProduct)
	router.GET("/product/:id", ProductById)
	router.POST("/product", createProduct)
	router.PATCH("/checkout", checkoutProduct)
	router.Run("localhost:8080")
}
