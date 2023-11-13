package teoria

/*
& Organización de los packages-

.
El domain-driven design (DDD) es una forma de estructurar y modelar el software según el dominio al que pertenece.

Sus principios se basan en:
- Colocar los modelos y reglas de negocio de la organización en el core de la aplicación.
 
- Fundamentar un dominio complejo en un modelo de software. 

- El desarrollo de una mejor perspectiva a nivel de colaboración entre expertos del dominio
y los desarrolladores, con el objetivo de concebir un software de objetivos claros.


/cmd - 
Esta carpeta contiene los archivos de punto de entrada de la aplicación principal. Cada uno de ellos será un paquete de la carpeta cmd y tendrán su propio main.go, además del resto de código necesario (configuraciones, entornos, inyección de dependencias, etc.). Es importante recordar que los paquetes dentro de cmd tendrán que corresponder con el nombre de los puntos de entrada, es decir, con el nombre de los binarios de las aplicaciones respectivas.

/internal - 
Este paquete contiene el código de la biblioteca privada que se usa en su servicio, es específico para la función del servicio y no se comparte con otros servicios

/pkg - 
Esta carpeta contiene código que permite ser consumido por otros servicios. Esto puede incluir clientes API o funciones útiles para otros.

& CAPAS DE UNA API

Para estructurar el proyecto por dominio lo que hacemos es generar un paquete por cada entidad. Cada paquete tendrá todas las capas de la entidad que le corresponda. De esta forma, si quisiéramos quitar la entidad producto para implementarlo en otro microservicio, el cambio sería muy sencillo.

Controlador:
Se encarga de recibir las peticiones que vienen del front end y devolver las respuestas. También es quien verifica que los datos de entrada cumplan los criterios que se hayan puesto.

Servicio:
Se encarga de toda la capa de negocios de la API, procesa los datos y maneja los recursos externos, como la comunicación con otras APIs. 

Repositorio:
Se encarga de abstraer el acceso a los datos, le corresponde interactuar con la base de datos. 

* De esta manera, por ejemplo, el empleado tendrá su propio controller, service, repository. Cada paquete, entonces, será una entidad.



*/