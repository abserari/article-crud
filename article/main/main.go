package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	article "github.com/yhyddr/article-crud/article/controller/"
)

func main() {
	router := gin.Default()

	dbConn, err := sql.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}

	articleCon := article.New(dbConn, "article")

	articleCon.RegisterRouter(router.Group("/api/v1/article"))

	router.Run(":8000")
}
