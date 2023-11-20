package pg

//&VARIABLES DE ENTORNO:
/*
* ¿Qué son y cómo utilizarlas?
Una variable de entorno es un valor dinámico que el sistema operativo, y otros programas, pueden utilizar para determinar la información específica de tu computadora.
En otras palabras, una variable de entorno es algo que representa otra cosa, como una ubicación en el equipo, un número de versión, una lista de objetos, un usuario o una contraseña. Es también una variable dinámica que puede afectar al comportamiento de los procesos en ejecución en una computadora.
Normalmente, estos valores hacen referencia a archivos, directorios y funciones comunes del sistema cuya ruta concreta puede variar, pero que otros programas necesitan poder conocer.
Una manera bastante popular de configurar las variables de entorno es la que se conoce como los ficheros .env, también llamados coloquialmente como “dotenv files” o directamente “dotfiles”. De hecho, hoy en día, muchas tecnologías soportan este tipo de ficheros: los propios IDEs, Docker y algunos frameworks web son solo unos pocos ejemplos de ello

* Instalación
Para instalar el pkg godotenv, desde la consola de comandos, ejecutamos la siguiente instrucción:

? go get -u github.com/joho/godotenv

Tendremos que crear el archivo .env junto al main del proyecto y agregar variables.

MY_USER=nombre_de_usuario
MY_PASS=abcd1234


Código de ejemplo

 */
 import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error al intentar cargar archivo .env")
    }
    usuario := os.Getenv("MY_USER")
    password := os.Getenv("MY_PASS")

    println("Usuario sacado de variables de Entorno: ", usuario)
    println("Password sacado de variables de Entorno: ", 
password)
} 

/* Antes de correr nuestro proyecto, tenemos que sumarle 
go mod init:

? go mod init Variables-de-entorno

Y también:
? go mod tidy 
para asegurar que nuestro fichero go.mod refleje las dependencias de todas las posibles combinaciones de sistema operativo, arquitectura y build tags.

? go mod init Variables-de-entorno
Finalmente, nuestro proyecto debería quedar así:

repo_variables_de_entorno
(tuerca).env
go.mod
go.sum
main.go */

//& MANEJO DE DATOS:
/*
3 FORMAS:
1) 
Archivo txt
El archivo de texto es plano y puede ser muy útil para poder almacenar datos. Si bien es compatible con todos los sistemas, no es una opción muy rápida ni muy segura de almacenar nuestros datos.

2) 
JSON
JSON es un lenguaje de modelado de datos. Es un formato flexible, ligero y fácilmente transferible. Los valores pueden ser cadenas, números o booleanos, así como otros objetos JSON, con cualquier nivel de anidación. Todos los lenguajes disponen de funciones para interpretar archivos JSON y así su contenido puede ser consumido por la aplicación. 

3)
Base de datos SQL
Las bases de datos son muchísimo más eficientes y permiten almacenar una gran cantidad de información de una forma organizada para su futura consulta, realización de búsquedas, nuevo ingreso de datos, etc. Todo esto se realiza de una forma rápida y simple.
*/
