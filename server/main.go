package main

import (
	// "log"
	// "fmt"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := setupDB()
	if err != nil {
		fmt.Println("could not open db")
		return
	}
	defer db.Close()

	populateDB(db)
	dbPrintBallots(db)

	router := gin.Default()
	router.GET("/ballots", getAllBallots)
	router.GET("/ballots/:id", getBallotByID)
	router.POST("/vote", addVote)
	// router.Run(":8080")
}
