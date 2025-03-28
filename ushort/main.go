package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ushort/config"
	"ushort/services"
)

type RequestBody struct {
	Password string `json:"password,omitempty"`
}

func main() {
	r := gin.Default()

	r.POST("/", handlePostRequest)
	r.GET("/:id", handleGetRequest)

	r.Run(":8080")
}

func handlePostRequest(c *gin.Context) {
	cfg := config.LoadConfig(c)

	link := c.Query("link")
	var body RequestBody

	password := ""
	if err := c.ShouldBindJSON(&body); err == nil {
		password = body.Password
	}

	lnk, err := services.SaveLink(c, *cfg, link, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"link": lnk})
}

func handleGetRequest(c *gin.Context) {
	cfg := config.LoadConfig(c)

	link := c.Param("id")
	var body RequestBody

	password := ""
	if err := c.ShouldBindJSON(&body); err == nil {
		password = body.Password
	}

	lnk, err := services.GetLink(c, *cfg, link, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(302, lnk)
}
