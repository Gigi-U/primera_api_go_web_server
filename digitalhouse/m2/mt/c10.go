package mt

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var productsList = []Product{}

// sería main()
func MesaC10() {
	loadProducts("/digitalhouse/m2/mt/products.json", &productsList)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", GetAllProducts())
		products.GET(":id", GetProduct())
		products.GET("/search", SearchProduct())
	}
	r.Run(":8080")
}

// loadProducts carga los productos desde un archivo json
func loadProducts(path string, list *[]Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}

// GetAllProducts traer todos los productos almacenados
func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, productsList)
	}
}

// GetProduct traer un producto por id
func GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid id"})
			return
		}
		for _, product := range productsList {
			if product.Id == id {
				ctx.JSON(200, product)
				return
			}
		}
		ctx.JSON(404, gin.H{"error": "product not found"})
	}
}

// SearchProduct traer un producto por nombre o categoria
func SearchProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Query("priceGt")
		priceGt, err := strconv.ParseFloat(query, 32)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid price"})
			return
		}
		list := []Product{}
		for _, product := range productsList {
			if product.Price > priceGt {
				list = append(list, product)
			}
		}
		ctx.JSON(200, list)
	}
}
/*  del profe
package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Producto es una estructura que define ...
type Producto struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	IsPublished bool      `json:"is_published"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

// store es una base de datos en memoria
type Store struct {
	Productos []Producto
}

func main() {

	// Carga la base de datos en memoria
	store := Store{}
	store.LoadStore()

	engine := gin.Default()

	group := engine.Group("/api/v1")
	{
		group.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"mensaje": "pong",
			})
		})

		grupoProducto := group.Group("/producto")
		{
			grupoProducto.GET("", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"data": store.Productos,
				})
			})

			grupoProducto.GET("/search/:parametroPrecio", func(ctx *gin.Context) {

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

			})

		}

	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

// LoadStore carga la base de datos en memoria
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


*/


//! el contexto sirve para pasar info entre capas