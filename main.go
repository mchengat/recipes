package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Run()
}

type Recipe struct {
	Name         string   `json:"name,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Ingredients  []string
	Instructions []string
	PublishedAt  time.Time
}
