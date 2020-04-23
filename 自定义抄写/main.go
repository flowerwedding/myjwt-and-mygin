package main

import (
	"why-dream-believe/自定义抄写/rrf"
)

func main(){
	app :=rrf.Default()
	app.Use(rrf.Logging())

	app.GET("/book",QueryBook)
	app.POST("/cake",FormCake)

	app.Run(8080)
}

func QueryBook(c *rrf.Context){
	bid := c.Query("id")
	c.String("your book id is  "+bid)
}

func FormCake(c *rrf.Context){
	bid := c.PostForm("id")
	name:= c.PostForm("name")
	c.JSON(rrf.H{"count":bid,"belong":name})
}