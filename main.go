package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"social-todo-list/common"
	"social-todo-list/middleware"
	ginitem "social-todo-list/modules/item/transport/gin"
)

func main() {
	dsn := os.Getenv("DB_CONNECT_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.Use(middleware.Recovery())

	//Các cách dùng middleware
	//c1:
	// r.Use(middleware.Recovery())
	//C2
	//v1 := r.Group("/v1", middleware.Recovery())
	//C3:
	//items.POST("", middleware.Recovery(),ginitem.CreateItem(db))

	//CRUD

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0])
		}()
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
}
