package main

import (
	// "log"
	// "fmt"

	"fmt"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/gorilla/websocket"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	db, err := setupDB()
	if err != nil {
		fmt.Println("could not open db")
		return
	}
	defer db.Close()

	// populateDB(db)
	// populateDB(db)
	// dbPrintBallots(db)

	// var tempbal ballot

	// if tempbal, err = dbGetCurrentBallot(db); err != nil {
	// 	fmt.Printf("%v", err)
	// }

	// fmt.Println("Pringint ballot")
	// fmt.Println(tempbal)

	// dbPrintById(db, 39988)

	// v := vote{BallotID: "76296", PromptID: "2"}

	// if _, err := dbVote(db, v); err != nil {
	// 	fmt.Printf("%v", err)
	// }

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/mysecretkey/ballots", getAllBallots)
	router.GET("/mysecretkey/ballots/setCurrent/:id", setCurrentBallot)
	// router.POST("/mysecretkey/setCurrentBallot", setCurrentBallot)
	// router.GET("/ballots/:id", getBallotByID)
	router.GET("/ballots/current", getCurrentBallot)
	router.POST("/vote", addVote)
	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
