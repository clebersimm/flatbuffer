package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/Player", func(c *gin.Context) {
		//TODO add response with flatbuffer
	})
	r.Run("4040")
}
