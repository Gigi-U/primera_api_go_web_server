package webserver

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)
const (
	filename = "./productos.json"
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
type Productos struct{
	productos []Producto
}

func Marshal() {
	router := gin.Default()
	// Captura la solicitud GET "/productos"
	router.GET("/productos", func(c *gin.Context) {
		// lee el archivo JSON, si hay algun error lo va a notificar sinó mostrará el json"
		data, err := os.ReadFile("internal/webserver/productos.json")
        if err != nil {
            c.JSON(500, gin.H{"error": "No se pudo cargar el archivo de productos.json"})
            return
        }
        c.Data(200, "application/json", data)
		fmt.Printf("Impresión de productos por Consola: %v",string(data))

	})
	// Corremos nuestro servidor sobre el puerto 8080
	// http://localhost:8080/productos	
	router.Run()

}