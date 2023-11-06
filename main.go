package main

import (
	"fmt"

	"github.com/Gigi-U/primera_api_go_web_server.git/internal/webserver"
)

func main(){
		fmt.Println("\n----------Gin web framework---------------")
	//webserver.WebServer()
	webserver.Marshal()


}

/* 
Gin es un microframework de alto rendimiento que se puede utilizar para crear aplicaciones web y
microservicios en Go.
Este contiene un conjunto de funcionalidades (por ejemplo: routing, middleware, rendering, etc.) que reducen el código repetitivo y simplifican la creación de aplicaciones web y microservicios. 

1) crear repositorio
2) clonar repositorio
3) go mod init github.com/MiUsuario/nombreDeMiRepo
4) go get -u github.com/gin-gonic/gin

PRUEBA TERMINAL =

{"message":"Hello World!"}



*/