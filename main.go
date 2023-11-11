package main

import (
	"fmt"

	"github.com/Gigi-U/primera_api_go_web_server.git/digitalhouse/m2/mt"
	"github.com/Gigi-U/primera_api_go_web_server.git/digitalhouse/m2/pg"
	"github.com/Gigi-U/primera_api_go_web_server.git/digitalhouse/m2/professor"
)

func main(){
		fmt.Println("\n----------Gin web framework---------------")
	fmt.Println("\n----------CLASE 9---------------")
	pg.WebServer()
	professor.ObtenerPersonas()
	mt.ObtenerProductos()
	fmt.Println("\n----------CLASE 10---------------")
	professor.GetProducto()
	fmt.Println("\n----------CLASE 11---------------")
	fmt.Println("\n----------CLASE 12---------------")
	fmt.Println("\n----------CLASE 13---------------")



}

/* 
Gin es un microframework de alto rendimiento que se puede utilizar para crear aplicaciones web y
microservicios en Go.
Este contiene un conjunto de funcionalidades (por ejemplo: routing, middleware, rendering, etc.) que reducen el código repetitivo y simplifican la creación de aplicaciones web y microservicios. 

1) crear repositorio
2) clonar repositorio
3) go mod init github.com/MiUsuario/nombreDeMiRepo
4) go get -u github.com/gin-gonic/gin


*/