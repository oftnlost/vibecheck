package main

import (
    // web framework
    "github.com/gin-gonic/gin"

    // pretrained sentiment analysis models
    "github.com/grassmudhorses/vader-go/lexicon"
    "github.com/grassmudhorses/vader-go/sentitext"
)

func ping(c *gin.Context) {
    c.JSON(200, gin.H{"ping": "pong"})
}

func postSentiment(c *gin.Context) {
    inputString := c.PostForm("input")
    parsedInputString := sentitext.Parse(inputString, lexicon.DefaultLexicon)
    sentiment := sentitext.PolarityScore(parsedInputString)

    returnJson := gin.H{
        "input": inputString,
        "output": gin.H{
            "positive": sentiment.Positive,
            "negative": sentiment.Negative,
            "neural": sentiment.Neutral,
        },
    }

    c.JSON(200, returnJson)
}

func main() {
    router := gin.Default()

    api := router.Group("/api")
    {
        api.GET("/ping", ping)
        api.POST("/sentiment", postSentiment)
    }

    router.Run(":5000")
}

