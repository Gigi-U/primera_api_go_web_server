package crud

/*

PATCH

Este método nos será de mucha utilidad cuando queramos realizar una actualización parcial de un recurso. "Emparcha" el recurso

* Una solicitud PATCH no debe tener el recurso completo sino solamente las propiedades que desea modificar. No necesita ser IDEMPOTENTE

en la url va la ruta del objeto a editar y en el body va lo que se desea modificar

^ si el servidor recibe peticiones PATCH duplicadas, pueden surgir incoherencias
*/

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User2 struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}


type UserPatchRequest struct {
	ID       *int    `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

var Users2 = []User2{
	{1, "Juan Pablo", "email_falso@fakemail.com"},
	{2, "José Perez", "email_falso@fakemail.com"},
	{3, "Jaime Pogo", "email_falso@fakemail.com"},
	{4, "Jenaro Pinto", "email_falso@fakemail.com"},
}
// sería main()
func MyPatch() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context){
		c.JSON(http.StatusOK, Users2)
		return
	})

	r.PATCH("/users/:id", func(c *gin.Context){
		userID := c.Param("id")
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var userReq UserPatchRequest
		err = c.BindJSON(&userReq)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for k, user2 := range Users2 {
			if user2.ID == userIDInt {

				switch {
					case userReq.ID != nil: 
						user2.ID = *userReq.ID
						fallthrough
					case userReq.Username != nil: 
						user2.Username = *userReq.Username
						fallthrough
					case userReq.Email != nil: 
						user2.Email = *userReq.Email
				}

				Users2[k] = user2
				c.JSON(http.StatusOK, user2)
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user %d not found", userIDInt)})
		return

	})

	r.Run(":8080")
}

// ver galeria validatos- 

//^ se usa PATCH si tenemos q modificar 1 o 2 atributos de 30 , x ejemplo. sino se usa el PUT