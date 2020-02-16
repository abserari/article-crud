package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	mysql "github.com/yhyddr/article-crud/article/mysql/"
)

type ArticleController struct {
	db        *sql.DB
	tableName string
}

func New(db *sql.DB, tableName string) *ArticleController {
	return &ArticleController{
		db:        db,
		tableName: tableName,
	}
}

// RegisterRouter
func (a *ArticleController) RegisterRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}

	err := mysql.CreateTable(a.db, a.tableName)
	if err != nil {
		log.Fatal(err)
	}

	r.POST("/create", a.create)
	r.POST("/delete", a.deleteByID)
	r.POST("/update", a.updateByID)
	r.POST("/query", a.queryByID)
}

// create
func (a *ArticleController) create(c *gin.Context) {
	var (
		req struct {
			ArticleName string `json:"articlename"    binding:"required"`
			Author      string `json:"author"         binding:"required"`
			Text        string `json:"text"           binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	id, err := mysql.InsertArticle(a.db, a.tableName, req.ArticleName, req.Author, req.Text)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "ID": id})
}

// delete
func (a *ArticleController) deleteByID(c *gin.Context) {
	var (
		req struct {
			ArticleID int `json:"articleid"    binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = mysql.DeleteArticleByID(a.db, a.tableName, req.ArticleID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// update
func (a *ArticleController) updateByID(c *gin.Context) {
	var (
		req struct {
			ArticleID int    `json:"articleid"     binding:"required"`
			Text      string `json:"text"          binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	art, err := mysql.UpdateArticleByID(a.db, a.tableName, req.Text, req.ArticleID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "art": art})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// query
func (a *ArticleController) queryByID(c *gin.Context) {
	var (
		req struct {
			ArticleID int `json:"articleid"     binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	art, err := mysql.QueryArticleByID(a.db, a.tableName, req.ArticleID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "art": art})
}
