package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Joke struct {
	ID   int    `json:"id" binding:"required"`
	Joke string `json:"joke" binding:"required"`
}

var jokes = []Joke{
	Joke{1, "4 Years ago dad used to nag me for overusing Whats App today I have to tell him 'Papa thodi der phone side rakh do yaar'."},
	Joke{2, "Papa kehte hain 'Kuch kaam-waam bhi karega?' "},
	Joke{3, "Son wanted to go on a trip. Reminded him about the things he could've done and been somewhere in life by now. A guilt trip is still a trip."},
	Joke{4, "An average Indian Father's way to show they care for their child is by being really angry & beating the shit out of them"},
	Joke{5, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, "Why did the coffee file a police report? It got mugged."},
	Joke{7, "How does a penguin build it's house? Igloos it together."},
	Joke{8, "Sarcasm is like electricity, half of India doesn't get it."},
	Joke{9, "Mayawati Ctrls + All + Dalit."},
	Joke{10, "Never say 'give me five' to a snake. Woh tumhe dus dega."},
	Joke{11, "What do you call drunk Pandavas? - High five."},
	Joke{12, "We are against reservation. - IRCTC"},
	Joke{13, "God never tasted any cough syrup, because khuda-na-khasta."},
	Joke{14, "Rahul Dravid's wristwatch is technically a wall clock."},
	Joke{15, "Vishwanathan Anand gets tensed when the waiter in the hotel says 'Sir Check'."},
}

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/jokes", JokeHandler)
	}
	router.Run(":3000")
}
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	i := rand.Intn(14) + 1
	c.JSON(http.StatusOK, jokes[i])
}
