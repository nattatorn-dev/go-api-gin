package main

import (
	"flag"
	"fmt"
	"os"
	// "net/http"

	"api/config"
	"api/services/aws"
	// "github.com/rs/cors"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db *dynamo.DB
var awsdb *dynamodb.DynamoDB


func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	
	config.Init(*enviroment)
	config := config.GetConfig()

	awsSession := session.Must(session.NewSession())
	dynamoConfig := service.CreateConfig(config.GetString("dynamodb.id"), config.GetString("dynamodb.secret"), "")
	db = dynamo.New(awsSession, dynamoConfig)
	awsdb = dynamodb.New(awsSession, dynamoConfig)



	// initDBTable(db)

	if *enviroment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := setupRouter()

	r.Run(":" + config.GetString("port"))
}