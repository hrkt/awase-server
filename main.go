package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	// "github.com/swaggo/files"
	// "github.com/swaggo/gin-swagger"
	//_ "github.com/swaggo/gin-swagger/example/basic/docs"
	//_ "./docs"
)

// AppConfig is a struct for app-config.json
type AppConfig struct {
	// a command to execute
	Command string `json:"command"`
	// arguments for the command
	Arguments []string `json:"arguments"`
}

var (
	// Version number
	Version string
	// Revision number
	Revision string
	// application configuration
	appConfig AppConfig
)

func init() {
	file, err := ioutil.ReadFile("app-settings.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &appConfig)
	fmt.Printf("Command :%s\n", appConfig.Command)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Global middleware
	router.Use(gin.Logger())

	// Routing
	router.StaticFile("/", "./index.html")

	router.POST("/api/events", func(ctx *gin.Context) {
		event := NewEvent("test")
		SaveEvent(event)
		ctx.JSON(200, gin.H{
			"result": "created",
		})
	})

	router.GET("/api/events/:eventid", func(ctx *gin.Context) {
		eventID := ctx.Param("eventid")

		event, err := LoadEvent(eventID)

		if err != nil {
			ctx.JSON(404, gin.H{})
		}

		ctx.JSON(200, gin.H{
			"eventID": event.ID,
			"label":   event.Label,
		})
	})

	router.GET("/api/events", func(ctx *gin.Context) {
		res := LoadEventIds()

		ctx.JSON(200, gin.H{
			"eventIDs": res,
		})
	})

	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

// @title AWASE API
// @version 0.1
// @description AWASE server
// @termsOfService https://github.com/hrkt/awase-server

// @contact.name API Support
// @contact.url https://github.com/hrkt/awase-server
// @contact.email doesnotexist@example.com

// @license.name MIT
// @license.url https://github.com/hrkt/awase-server

// @host xxx.xxx.xxx
// @BasePath /v2
func main() {

	fmt.Println("AWASE Server : Version:" + Version + " Revision:" + Revision)

	endless.ListenAndServe(":8080", setupRouter())
}
