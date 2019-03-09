package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Post struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Post  string `json:"post"`
}

func main() {
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function
	db, err = gorm.Open("sqlite3", "./hello-gin.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&Post{})
	r := gin.Default()
	r.GET("/post/", GetPosts)
	r.GET("/post/:id", GetPost)
	r.POST("/post", CreatePost)
	r.PUT("/post/:id", UpdatePost)
	r.DELETE("/post/:id", DeletePost)
	r.Run(":8080")
}
func DeletePost(c *gin.Context) {
	id := c.Params.ByName("id")
	var post Post
	d := db.Where("id = ?", id).Delete(&post)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
func UpdatePost(c *gin.Context) {
	var post Post
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&post)
	db.Save(&post)
	c.JSON(200, post)
}
func CreatePost(c *gin.Context) {
	var post Post
	c.BindJSON(&post)
	db.Create(&post)
	c.JSON(200, post)
}
func GetPost(c *gin.Context) {
	id := c.Params.ByName("id")
	var post Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, post)
	}
}

func GetPosts(c *gin.Context) {
	var people []Post
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}
}
