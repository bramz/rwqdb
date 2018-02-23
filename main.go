package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

var db *sql.DB

func submitQuote(context *gin.Context) {
	context.HTML(http.StatusOK, "public/index.tmpl", gin.H{
		"title": "RootWire Quote Database",
	})
	quote := context.PostForm("quote")
    fmt.Println(quote)

	stmt, err := db.Prepare("INSERT INTO quotes (quote) values (?)")
	if err != nil {
		fmt.Println(err)
	}

    _, err = stmt.Exec(quote)

    if err != nil {
        fmt.Println(err)
    }

}

func randomQuote(context *gin.Context) {
}

func searchQuote(context *gin.Context) {
}

func deleteQuote(context *gin.Context) {
}

func main() {
	router := gin.Default()

	connect, err := sql.Open("sqlite3", "qdb.db")

	if err != nil {
		fmt.Println(err)
	}

    db = connect
    defer db.Close()
    fmt.Println(db)
	//    router.LoadGLOB("public/*")
	router.LoadHTMLFiles("public/index.tmpl")

	router.GET("/", submitQuote)
	router.POST("/submit", submitQuote)
	router.POST("/search/:id", searchQuote)
	router.POST("/delete/:id", deleteQuote)
	router.GET("/random", randomQuote)
	router.Run(":8080")
}
