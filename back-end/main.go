package main

import (
	"example.com/hello/handler"
	"example.com/hello/repository"
	"example.com/hello/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:XXVebw51@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	taskRespository := repository.NewTaskRepositoryDB(db)

	// tx, err := taskRespository.GetTaskAll()
	// if err != nil {
	// 	return
	// }
	// fmt.Println(tx)

	taskService := service.NewTaskService(taskRespository)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	router.GET("/tasks", taskHandler.GetTaskAllCos)
	router.GET("/tasks/:taskID", taskHandler.GetTaskByIdCos)
	router.POST("/tasks", taskHandler.CreateTaskCos)
	router.PUT("/tasks/:taskID", taskHandler.UpdateTaskCos)
	router.DELETE("/tasks/:taskID", taskHandler.DeleteTaskCos)

	router.Run(":8080")

}
