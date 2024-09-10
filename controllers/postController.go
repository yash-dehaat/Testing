package controllers

import (
	"fmt"
	"goblogart/inits"
	"goblogart/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {

	var body struct {
		Name  string
		Age   int
		Email string
	}

	ctx.BindJSON(&body)

	post := models.Post{Name: body.Name, Age: body.Age, Email: body.Email}

	result := inits.DB.Create(&post)
	fmt.Println(post)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": post})

}

func GetPosts(ctx *gin.Context) {

	var posts []models.Post

	result := inits.DB.Find(&posts)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": posts})

}

func UpdatePost(ctx *gin.Context) {

	var body struct {
		Name  string
		Age   int
		Email string
	}

	ctx.BindJSON(&body)

	var post models.Post

	result := inits.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	inits.DB.Model(&post).Updates(models.Post{Name: body.Name, Age: body.Age, Email: body.Email})

	ctx.JSON(200, gin.H{"data": post})

}

func DeletePost(ctx *gin.Context) {

	id := ctx.Param("id")

	inits.DB.Delete(&models.Post{}, id)

	ctx.JSON(200, gin.H{"data": "post has been deleted successfully"})

}
