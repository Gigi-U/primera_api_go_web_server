package crud

import "github.com/gin-gonic/gin"

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
// sería main()
func ExamplePost() {
	engine := gin.Default()

	engine.POST("/login", func(c *gin.Context) {
		var login Login
		// bindeamos (mapeamos) el body en una variable Login
		c.BindJSON(&login)
		// respuesta hacia el cliente
		c.JSON(200, gin.H{"status": login.Email})
	})

	engine.Run(":8080")
}

//* primero go run main.go y luego desde powershell o desde bagin

//! comando para hacer peticion desde bash
/*
curl --location --request POST 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
	"email": "Gise Urriza",
	"password": "secret"
}'

! comando para hacer peticion desde powershell
/*
curl 'http://localhost:8080/login' -Method POST -Headers @{'Content-Type'='application/json'} -Body '{
	"email": "Gise Urriza",
	"password": "secret"
}'
*/

/*
? desde git bash me muestra esto:
{"status":"Gise Urriza"}
*/

/*
? desde powershell me muestra esto:

StatusCode        : 200
StatusDescription : OK
Content           : {"status":"Gise Urriza"}
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 24
                    Content-Type: application/json; charset=utf-8
                    Date: Sun, 12 Nov 2023 20:51:10 GMT

                    {"status":"Gise Urriza"}
Forms             : {}
Headers           : {[Content-Length, 24], [Content-Type, application/json;
                    charset=utf-8], [Date, Sun, 12 Nov 2023 20:51:10 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 24 */