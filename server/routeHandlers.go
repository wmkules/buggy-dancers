package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBallotByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range ballots {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func getAllBallots(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ballots)
}

// postAlbums adds an album from JSON received in the request body.
func addVote(c *gin.Context) {
	var newVote vote
	// fmt.Println(c)

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newVote); err != nil {
		return
	}

	bal, err := dbVote(db, newVote)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Could not vote")
		return
	}
	c.IndentedJSON(http.StatusOK, bal)
	// c.IndentedJSON(http.StatusOK, bal)
}
