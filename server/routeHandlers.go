package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024 ,
	WriteBufferSize: 1024 ,
	 // Resolve cross-domain problems 
	CheckOrigin: func(r *http.Request) bool {
		 return  true
	},
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		fmt.Println(t)
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
		fmt.Println(msg)
	}
}

func voteSocket(c *gin.Context) {
	wshandler(c.Writer, c.Request)
	return
	// c.String(http.StatusOK, "Hi, WebSocket!")
	// return
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
