package main

import (
	"P2/RESTAPI/adapters/database"
	"P2/RESTAPI/api/controller"
	"P2/RESTAPI/api/middleware"
	"P2/RESTAPI/core/ports"
	"P2/RESTAPI/core/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)


func main() {
	postGreSql := NewConnect()
	postGreSqlCollect := database.NewPostGresql(postGreSql)
	HumansRepository := ports.PortService(postGreSqlCollect)
	HumansService:= service.NewHumansService(HumansRepository)
	HumansServer := controller.NewController(HumansService)
	r:= gin.Default()
	r.POST("/login",HumansServer.Login)
	api := r.Group("/api",middleware.Auth())
	{
		api.GET("/entities",HumansServer.FindAll)
	}
	r.POST("/gentoken",HumansServer.RegisterUser)
	r.POST("/humans",HumansServer.CreateHumans)
	r.GET("/all",HumansServer.FindAll)
	r.GET("/id/:id",HumansServer.FindById)
	r.PUT("/update/:id",HumansServer.UpdateById)
	r.DELETE("/delete/:id",HumansServer.DeleteById)
	r.GET("/search/",HumansServer.FindByForm)
	r.Run()
}

func NewConnect()*gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=Humans port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect fail")
	} else {
		fmt.Print("Connect successfully")
	}
	return db
}