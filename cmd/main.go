package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/steyrin/mini-rest-api/config"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	router := gin.Default()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Добро пожаловать в REST API!",
		})
	})

	logrus.Info("Сервер запущен на порту 8080")
	if err := router.Run(":8080"); err != nil {
		logrus.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
