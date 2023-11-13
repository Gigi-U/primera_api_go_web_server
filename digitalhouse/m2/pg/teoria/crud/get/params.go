package get

import "github.com/gin-gonic/gin"

// Definimos una pseudo base de datos donde consultaremos la info
var empleados = map[string]string{
	"644": "Empleado A",
	"755": "Empleado B",
	"777": "Empleado C",
}

// sería main()
func MyParams() {
	server := gin.Default()
	server.GET("/", PaginaPrincipal)
	server.GET("/empleados/:id", BuscarEmpleado)
	server.Run(":8085")
}


// este HANDLER responde a "/"
func PaginaPrincipal(ctxt *gin.Context){
	ctxt.String(200, "Bienvenido a la empresa Gophers!")
}

// este HANDLER responde a "/empleados/:id"
func BuscarEmpleado(ctxt *gin.Context) {
	empleado, ok := empleados[ctxt.Param("id")]
	if ok{
		ctxt.String(200, "Información del empleado %s, nombre: %s", ctxt.Param("id"), empleado)
	} else{
		ctxt.String(404, "Información del empleado ¡No existe!")
	}
}

