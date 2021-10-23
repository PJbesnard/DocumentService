package main

import (
	"github.com/PJBesnard/DocumentService/internal/interface/handler"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()
	handler.DocumentHandler(r)
	err := r.Run(":8081")

	if err != nil {
		log.Fatal(err)
	}
}
