package mt

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
// Producto es una estructura que define un producto
type Producto struct {
    ID           	int     `json:"id"`
    Nombre       	string  `json:"nombre"`
    Precio       	float64 `json:"precio"`
    Stock       	int     `json:"stock"`
    Codigo       	string  `json:"codigo"`
    Publicado    	bool   	`json:"publicado"`
    FechaDeCreacion string 	`json:"fecha_de_creación"`
}
// sería main()
 func ObtenerProductos() {
	//			  func gin.Default() *gin.Engine
	engine := gin.Default() 
	// Captura la solicitud GET "/productos"
	engine.GET("/productos", func(c *gin.Context) {
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
		c.Data(http.StatusOK, "application/json", data)
		// imprimir el listado JSON por consola																			
		fmt.Printf("Impresión de productos por Consola: %v",string(data))

	})
	// Corremos nuestro servidor sobre el puerto 8080
	// http://localhost:8080/productos	
	engine.Run()

}