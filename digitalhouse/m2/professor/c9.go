package professor

import (
	"encoding/json"
	"fmt"
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
// o main()
func ObtenerPersonas(){

	// harcodeamos un objeto json 
	jsonPersona := `{
		"nombre":"Gi",         
		"apellido":"Urriza",         
		"edad":43,             
		"direccion":"Mi Casa",     
		"telefono":"1234546",    
		"activo": true
	}`
	// punto 1.
	// hacemos un unmarshal del json harcodeado contra nuestra estructura - vamos a recibir info de un json
	// el Unmarshal toma un Json y lo parsea a una estructura de datos. el Marshal toma una estructura y lo va a pasar a formato json.
	var persona Persona     // &persona es la posicion de memoria de la variable persona
	err := json.Unmarshal([]byte(jsonPersona), &persona)

	if err != nil {
		log.Fatal(err)
		log.Printf("ERROR: %v",err)
	}
	fmt.Println(persona)
	// tenemos que levantar el servicio de gin
	// instancia default gin
	engine:= gin.Default()

	// prueba inicial ping-pong para ver si hay comunicación - se recomienda incluir en todos los proyectos para estar seguros todo funciona ok
	// gin.H --> es un map donde se responde en formato json, clave-valor.
	engine.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"mensaje": "pong",
		})
	})

	// punto 2.
	// estructura de una persona
	personaResponse:= Persona{
		Nombre : 	"Cosme",         
		Apellido:	"Fulanito",         
		Edad: 		70,             
		Direccion:	"Su Casa",     
		Telefono: 	"1234546",    
		Activo :  	true,
	}
	// no es necesario hacer el marshal xq lo hace gin internamente
	engine.GET("/persona", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": personaResponse, 
		})
	})
	//engine.Run   --> así solo retorna error por eso se maneja el error
	if err := engine.Run(puerto); err != nil{
		log.Fatal(err)
	}
}  

/* 
Vamos a crear una aplicacion web con el framework Gin q tenga un endpoint /persona que responda con los datos de una persona
1) crear una estructura persona2) crear una persona en formato JSON (hardcodear) y aplicar la logica  para que 
	a) imprima la persona x consola
	b) imprima la persona a traves del endpoint en formato json. El endpoint debe ser de metodo GET 
*/


//^ command= go get -u github.com/gin-gonic/gin


//En vez de Go tmb se pueden usar Ch,  Mux o  eco  

//https://medium.com/@hasanshahjahan/a-deep-dive-into-gin-chi-and-mux-in-go-33b9ad861e1b

//https://github.com/go-chi/chi 

//* TODOS LOS STATUS DE HTTP.


/* 	StatusContinue           = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols = 101 // RFC 9110, 15.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 9110, 15.3.1
	StatusCreated              = 201 // RFC 9110, 15.3.2
	StatusAccepted             = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 9110, 15.3.4
	StatusNoContent            = 204 // RFC 9110, 15.3.5
	StatusResetContent         = 205 // RFC 9110, 15.3.6
	StatusPartialContent       = 206 // RFC 9110, 15.3.7
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices  = 300 // RFC 9110, 15.4.1
	StatusMovedPermanently = 301 // RFC 9110, 15.4.2
	StatusFound            = 302 // RFC 9110, 15.4.3
	StatusSeeOther         = 303 // RFC 9110, 15.4.4
	StatusNotModified      = 304 // RFC 9110, 15.4.5
	StatusUseProxy         = 305 // RFC 9110, 15.4.6

	StatusTemporaryRedirect = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              = 402 // RFC 9110, 15.5.3
	StatusForbidden                    = 403 // RFC 9110, 15.5.4
	StatusNotFound                     = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               = 408 // RFC 9110, 15.5.9
	StatusConflict                     = 409 // RFC 9110, 15.5.10
	StatusGone                         = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            = 417 // RFC 9110, 15.5.18
	StatusTeapot                       = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          = 422 // RFC 9110, 15.5.21
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
	 */