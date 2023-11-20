package professor

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Producto es una estructura que define un producto
type Producto struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	IsPublished bool      `json:"is_published"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

// store simula una base de datos en memoria. es un slice de Producto
type Store struct {
	Productos []Producto
}
// sería main()
func GetProducto() {

	// Carga la base de datos en memoria
	// creo el store
	store := Store{}
	// llamo al método
	store.LoadStore()

	engine := gin.Default()
		//v1 = versionado de apis
	group := engine.Group("/api/v1")
	{
		group.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"mensaje": "pong",
			})
		})
		// la api puede crecer, y asi como tenemos grupoProducto podemos tener grupoVendedores, grupoProveedores etc.
		grupoProducto := group.Group("/producto")
		{
			grupoProducto.GET("", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"data": store.Productos,
				})
			})
			grupoProducto.GET("/search/:parametroPrecio", func(ctx *gin.Context) {
			// implementacion de search
				precioParametro := ctx.Param("parametroPrecio")

				precioCasteado, err := strconv.ParseFloat(precioParametro, 64)
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"mensaje": "parametro invalido",
					})
					return
				}

				var result []Producto
				for _, producto := range store.Productos {
					if producto.Price > precioCasteado {
						result = append(result, producto)
					}
				}

				ctx.JSON(http.StatusOK, gin.H{
					"data": result,
				})
				// no necesita agregar el return xq está implícito en la funcion JSON

			})

		}

	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

// LoadStore es un método que carga la base de datos en memoria (el Slice)
func (s *Store) LoadStore() {
	s.Productos = []Producto{
		{
			Id:          "1",
			Name:        "Coco Cola",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.5,
		},
		{
			Id:          "2",
			Name:        "Pepsito",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.5,
		},
		{
			Id:          "3",
			Name:        "Fantastica",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.5,
		},
	}
}

/* context --> se usa para pasar info entre las  ≠ capas

x ejemplo le puedo pasar un Usuario.
Dde el handler, en base a una request le podes pasar que cuando se autentique un user puedas obtener el rol, y si x ej. es admin se pueda propagar el contexto desde el handler hta el repositorio y que se pueda hacer una validacion a nivel repositorio, de que ese usuario que es de tipo admin pueda hacer algo sobre la base o no.
X ejemplo si no es admin, que permita hacer todos los GET, pero no el resto. Una forma de propagar ese usuario entre las 3 capas, HANDLER --> SERVICE --> REPOSITORIO, es a traves del contexto.

En el handler envolves el contexto, le agregas el usuario al contexto y pasas el contexto con el usuario agregado, y en la ultima capa lo recuperas con el context.Value, le paso la key y recupero el valor
*/

//! se agrega el context aqui a modo de ejemplo, pero en sí se usa cuando hay separacion de capas

//^ QueryParam VS PARAM

// queryParam =  los parametros se pasan por query. esto hace que puedas o no pasar datos. haces un get en localhost:8080/persona/ y donde vaya la respuesta del parametro llevará (?)

// param = es un parametro especifico que agregamos a una url ej: localhost:8080/persona/{id}