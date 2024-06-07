package main

import (
	// "error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
  
type product struct{
	ID string `json:id`
	Name string `json:name`
	Price int `json:price`
}

var products = []product{
	{ID: "1", Name: "Fan", Price: 200},
	{ID: "2", Name: "Sneaker",  Price: 500},
	{ID: "3", Name: "Book",  Price: 600},
}

func generateID() string {
	// 使用當前產品數量加一作為新ID
	return strconv.Itoa(len(products) + 1)
}

func getProduct(c *gin.Context){
	c.JSON(http.StatusOK,products )
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

func main(){
	router := gin.Default()
	router.GET("/product",getProduct)
	router.POST("/product", createProduct)
	router.Run("localhost:8080")
}