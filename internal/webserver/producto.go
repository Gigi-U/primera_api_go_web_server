package webserver

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)
type Producto struct {
    ID           	int     `json:"Id"`
    Nombre       	string  `json:"Nombre"`
    Precio       	float64 `json:"Precio"`
    Stock       	int     `json:"Stock"`
    Codigo       	string  `json:"Codigo"`
    Publicado    	bool   	`json:"Publicado"`
    FechaDeCreacion string 	`json:"FechaDeCreación"`
}

func Marshal() {
	//			  func gin.Default() *gin.Engine
	router := gin.Default() 
	// Captura la solicitud GET "/productos"
	router.GET("/productos", func(c *gin.Context) {
		// lee el archivo JSON, si hay algun error lo va a notificar sinó mostrará el json"
		//func os.ReadFile(name string) ([]byte, error)
		data, err := os.ReadFile("internal/webserver/productos.json")
        if err != nil {
		//	   JSON: func (*gin.Context).JSON(code int, obj any) 
		//                | H: type H map[string]any
            c.JSON(500, gin.H{"error": "No se pudo cargar el archivo de productos.json"})
            return
        }
		// Data: func (*gin.Context).Data(code int, contentType string, data []byte)
		c.Data(200, "application/json", data)
		// imprimir el listado JSON por consola																			
		fmt.Printf("Impresión de productos por Consola: %v",string(data))

	})
	// Corremos nuestro servidor sobre el puerto 8080
	// http://localhost:8080/productos	
	router.Run()

}