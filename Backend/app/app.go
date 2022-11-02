package app

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	// router.Use(cors.New(cors.Config{
	//     AllowOrigins:     []string{"*"},
	//     AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
	//     AllowHeaders:     []string{"Origin"},
	//     ExposeHeaders:    []string{"Content-Length"},
	//     AllowCredentials: true,
	//     AllowOriginFunc: func(origin string) bool {
	//         return origin == "https://github.com"
	//     },
	//     MaxAge: 12 * time.Hour,
	// }))
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	// config.AllowAllOrigins = true
	router.Use(cors.New(config))

	//define routes
	router.GET("/greet", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"Title": "CodeName", "Desc": "Example API Call", "Content": "Hello Jacob and Paul"})
	})
	//starting server by default localhost :8080
	router.Run()
}
