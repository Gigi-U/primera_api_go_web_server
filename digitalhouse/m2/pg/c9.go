package pg


import "github.com/gin-gonic/gin"

func WebServer() {
	// Crea un router con gin
	router := gin.Default()
	// Captura la solicitud GET “/hello-world”
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	// Corremos nuestro servidor sobre el puerto 8080
	// http://localhost:8080/hello-world
	router.Run()
} 