package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

type Article struct {
	ArticleID   int
	ArticleName string
	Author      string
	Text        string
}

const (
	mysqlArticleCreateDB = iota
	mysqlArticleCreateTable
	mysqlArticleInsert
	mysqlArticleDeleteByID
	mysqlArticleUpdateByID
	mysqlArticleQueryByID
)

var (
	errInvalidInsert = errors.New("insert article:insert affected 0 rows")

	articleSQLstring = []string{
		`CREATE DATABASE IF NOT EXISTS %s`,
		`CREATE TABLE IF NOT EXISTS %s (
			articleId    INT           NOT NULL AUTO_INCREMENT,
			articleName  VARCHAR(128)  NOT NULL,
			author       VARCHAR(128)  NOT NULL,
			text         TEXT          NOT NULL,
			PRIMARY KEY (articleId)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`,
		`INSERT INTO %s (articleName,author,text) VALUES (?,?,?)`,
		`DELETE FROM %s WHERE articleId = ? LIMIT 1`,
		// `UPDATE %s.%s SET articleName = ? WHERE id = ? LIMIT 1`
		// `UPDATE %s.%s SET author = ? WHERE id = ? LIMIT 1`,
		`UPDATE %s SET text = ? WHERE articleId = ? LIMIT 1`,
		`SELECT * FROM %s WHERE articleId = ? LIMIT 1`,
	}
)

//createDB
func CreateDB(db *sql.DB, DBName string) error {
	sql := fmt.Sprintf(articleSQLstring[mysqlArticleCreateDB], DBName)
	_, err := db.Exec(sql)
	return err
}

//createTable
func CreateTable(db *sql.DB, tableName string) error {
	sql := fmt.Sprintf(articleSQLstring[mysqlArticleCreateTable], tableName)
	_, err := db.Exec(sql)
	return err
}

//insertArticle
func InsertArticle(db *sql.DB, tableName string, articleName string, author string, text string) (int, error) {
	sql := fmt.Sprintf(articleSQLstring[mysqlArticleInsert], tableName)
	result, err := db.Exec(sql, articleName, author, text)
	if err != nil {
		return 0, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return 0, errInvalidInsert
	}

	articleId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(articleId), nil
}

//deleteArticleByID
func DeleteArticleByID(db *sql.DB, tableName string, id int) error {
	sql := fmt.Sprintf(articleSQLstring[mysqlArticleDeleteByID], tableName)
	_, err := db.Exec(sql, id)
	return err
}

//updateArticleByID
func UpdateArticleByID(db *sql.DB, tableName string, text string, id int) error {
	sql := fmt.Sprintf(articleSQLstring[mysqlArticleUpdateByID], tableName)
	_, err := db.Exec(sql, text, id)
	return err
}

//queryArticleByID
func QueryArticleByID(db *sql.DB, tableName string, id int) (*Article, error) {
	var art Article

	sql := fmt.Sprintf(articleSQLstring[mysqlArticleQueryByID], tableName)
	rows, err := db.Query(sql, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&art.ArticleID, &art.ArticleName, &art.Author, &art.Text); err != nil {
			return nil, err
		}
	}

	return &art, nil
}
