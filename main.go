package main

import (
	"fmt"
	"github.com/VladislavZernov/blogging-platform-api/routes"
	"log"
	//

	"github.com/VladislavZernov/blogging-platform-api/config" // Подключаем базу данных
	//"github.com/gin-gonic/gin"
)

func main() {
	// Подключаем базу данных
	config.ConnectDatabase()

	// Настраиваем маршруты
	router := routes.SetupRouter()

	// Запускаем сервер
	port := "8080"
	fmt.Println("Server is running on port " + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
