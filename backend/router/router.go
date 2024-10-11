package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/parvejmia9/healthythakun/database"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 3600, // Max age for preflight requests
		AllowCredentials: true,      // Allow credentials (e.g., cookies, authorization headers)
	}

	// Use CORS middleware
	router.Use(cors.New(corsConfig))

	return router
}

func HandleReq(router *gin.Engine) error {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "pong",
		})
	})
	router.GET("/tasks", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"tasks": database.GetTasks(),
		})
	})
	router.GET("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		tasks := database.GetTasks()
		for _, task := range tasks {
			if task.ID == id {
				c.JSON(200, gin.H{
					"task": task,
				})
				return
			}
		}
	})
	err := router.Run(":8181")
	if err != nil {
		return err
	}
	return nil
}
