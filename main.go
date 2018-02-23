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
	context.HTML(http.StatusOK, "public/submit.tmpl", gin.H{
		"header": "Submit Quote",
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
    rows, err := db.Query("SELECT id, quote FROM quotes ORDER BY RANDOM() LIMIT 1")
    if err != nil {
        panic(err)
    }

    var id int
    var quote string

    for rows.Next(){
        err = rows.Scan(&id, &quote)

        if err != nil {
            panic(err)
        }

        context.HTML(http.StatusOK, "public/index.tmpl", gin.H{
            "header": "Random Quotes",
            "quote": quote,
        })
    }

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

	router.LoadHTMLFiles("public/index.tmpl", "public/submit.tmpl")

	router.GET("/", randomQuote)
	router.POST("/submit", submitQuote)
	router.POST("/search/:id", searchQuote)
	router.POST("/delete/:id", deleteQuote)
	router.GET("/random", randomQuote)
	router.Run(":8080")
}
