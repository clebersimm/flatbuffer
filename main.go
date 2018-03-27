package main

import (
	player "teste"

	"github.com/gin-gonic/gin"
	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	r := gin.Default()
	r.GET("/Player", func(c *gin.Context) {
		builder := flatbuffers.NewBuilder(0)
		player.PlayerStart(builder)
		playerName := builder.CreateString("test")
		playerTitle := builder.CreateString("the destroyer")
		player.PlayerAddName(builder, playerName)
		player.PlayerAddTitle(builder, playerTitle)
		currentPlayer := player.PlayerEnd(builder)
		builder.Finish(currentPlayer)
		c.Writer.Header().Set("Content-type", "application/octet-stream")
		c.Writer.WriteHeader(200)
		c.Writer.Write(builder.Bytes)
		//TODO add response with flatbuffer
	})
	r.Run(":8040")
}
