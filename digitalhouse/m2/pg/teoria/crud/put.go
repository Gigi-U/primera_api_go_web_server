package crud



/* se emplea para actualizar un recurso, reemplazandolo en su totalidad x un recurso nuevo. Una solicitud PUT siempre tiene un recurso completo.
* Una cualidad de las solicitudes PUT es la IDEMPOTENCIA. Es decir producir el mismo resultado si la misma solicitud se realiza varias veces.

~ Usamos PUT siempre que querramos modificar un recurso en su TOTALIDAD. 
Si ejecutamos el mismo PUT varias veces, el resultado será siempre el mismo
*/

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User1 struct{
    ID int `json:"id"`
    Username string `json:"username" binding:"required"`
}

var Users1 = []User1 {
	{1, "Juan Pablo"},
	{2, "José Perez"},
	{3, "Jaime Pogo"},
	{4, "Jenaro Pinto"},
}
// sería main()
func MyPut(){
	r := gin.Default()

	r.PUT("/users/:id", func(c *gin.Context) {
		var user1 User1
		c.BindJSON(&user1)

		idReceived := c.Param("id")
		idReceivedInt, err := strconv.Atoi(idReceived)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for k, u := range Users1 {
			if u.ID == idReceivedInt {
				user1.ID = idReceivedInt
				Users1[k] = user1
				c.JSON(http.StatusOK, user1)
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", idReceivedInt)})
	})

	r.Run(":8080")
}