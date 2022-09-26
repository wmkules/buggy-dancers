package main

import (
	// "log"
	// "fmt"

	"fmt"

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

	// dbPrintById(db, 39988)

	// v := vote{BallotID: "76296", PromptID: "2"}

	// if _, err := dbVote(db, v); err != nil {
	// 	fmt.Printf("%v", err)
	// }

	router := gin.Default()
	router.GET("/ballots", getAllBallots)
	router.GET("/ballots/:id", getBallotByID)
	router.POST("/vote", addVote)
	router.GET("/vs", voteSocket)
	router.Run(":8080")
}
