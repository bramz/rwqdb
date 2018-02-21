package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func randomQuote(context *gin.Context) {
    context.HTML(http.StatusOK, "public/index.tmpl", gin.H{
        "title": "RootWire Quote Database",
    })
}

func submitQuote(context *gin.Context) {
}

func searchQuote(context *gin.Context) {
}

func deleteQuote(context *gin.Context) {
}

func main() {
    router := gin.Default()

//    router.LoadGLOB("public/*")
    router.LoadHTMLFiles("public/index.tmpl")

    router.GET("/", randomQuote)
    router.POST("/submit", submitQuote)
    router.POST("/search/:id", searchQuote)
    router.POST("/delete/:id", deleteQuote)

    router.Run(":8080")
}


