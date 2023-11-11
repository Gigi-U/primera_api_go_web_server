package professor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Persona es una estructura que define una persona
type Persona struct {
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Edad            int    `json:"edad"`
	Direccion       string `json:"direccion"`
	Telefono        string `json:"telefono"`
	Activo       	bool   `json:"activo"`
}
const (
	puerto = ":8080"
)
// o main
func ObtenerPersonas() {

	// punto 1 - unmarsheamos un onjeto json
	jsonPersona:= `{
		"Nombre" : "Gi",         
		"Apellido":	"Urriza",         
		"Edad": 43,             
		"Direccion": "Mi Casa",     
		"Telefono": "1234546",    
		"Activo" :  true,
	}`
	var persona Persona


	if err := json.Unmarshal([]byte(jsonPersona), &persona); err != nil {
		log.Fatal(err)
		log.Printf("ERROR: %v",err)
	}

	// punto 2 
	// instancia default gin
	engine:= gin.Default()
	// prueba inicial ping-pong para ver si hay comunicación
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"mensaje": "pong" })
	})

	personaResponse:= Persona {
		Nombre : "Cosme",         
		Apellido:	"Fulanito",         
		Edad: 70,             
		Direccion: "Su Casa",     
		Telefono: "1234546",    
		Activo :  true,
	}
	// creamos el 

	engine.GET("/persona", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": personaResponse, 
		})
	})

	if err := engine.Run(puerto); err != nil{
		log.Fatal(err)
	}
}  