package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sachinvarma/go-gin-postgresql-api/books"
	"github.com/sachinvarma/go-gin-postgresql-api/common/db"
)

func main() {

	port := ":9000"
	dbUrl := "postgresql://localhost/sachinvarma"

	router := gin.Default()
	dbHandler := db.Init(dbUrl)

	books.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	router.Run(port)

}
