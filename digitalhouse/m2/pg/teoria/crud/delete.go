package crud
/* 
DELETE = eliminar recurso de una colección. a traves de id o algun elemento/s

url x defecto+ coleccion + id
 */

 import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User3 struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
}

var Users3 = []User3{
	{1, "Juan Pablo"},
	{2, "José Perez"},
	{3, "Jaime Pogo"},
	{4, "Jenaro Pinto"},
}
// sería main()
func MyDelete() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context){
		c.JSON(http.StatusOK, Users3)
		return
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		idReceived := c.Param("id")
		idReceivedInt, err := strconv.Atoi(idReceived)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		for k, u := range Users3 {
			if u.ID == idReceivedInt {
				Users3 = append(Users3[:k], Users3[k+1:]...)
				c.JSON(http.StatusOK, gin.H{"status":fmt.Sprintf("user %d deleted", idReceivedInt)})
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user %d not found", idReceivedInt)})

	})

	r.Run(":8080")
} 