package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yhyddr/article-crud/article/controller"
)

func main() {
	router := gin.Default()

	dbConn, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/?parseTime=true")
	if err != nil {
		panic(err)
	}

	articleCon := controller.New(dbConn, "article", "article.article_crud")

	articleCon.RegisterRouter(router.Group("/api/v1/article"))

	router.Run(":8000")
}
