package controllers

import (
	db2 "first-api-go/db"
	"first-api-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ShareBook( c *gin.Context){
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil{
		c.JSON(403,gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := db2.GetDataBase()

	var book models.Book
	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(403,gin.H{
			"error": "cannot find book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook (c *gin.Context){
	db := db2.GetDataBase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(403,gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&book).Error
	if err != nil {
		c.JSON(403,gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)

}

func ShowBooks(c *gin.Context)  {
	db := db2.GetDataBase()
	var p []models.Book
	err := db.Find(&p).Error
	if err != nil {
		c.JSON(403,gin.H{
			"error": "cannot list books: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,p)
}