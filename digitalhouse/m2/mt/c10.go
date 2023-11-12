package mt

/* Mesa de trabajo - Clase 14
Ejercitación grupal
Nivel de complejidad: medio 🔥🔥
Consigna
Un supermercado necesita un sistema para gestionar los productos frescos que tienen publicados en su página web. Para poder hacer esto, necesitan un servidor que ejecute una API que les permita manipular los productos cargados desde distintos clientes. Los campos que conforman un producto son:
Nombre | Tipo dato JSON  | 	Tipo dato Go | 	Descripción / Ejemplo
id 				number			int				Identificador en conjunto de datos/15
name 			string			string			Nombre caracterizado/Cheese - St. Andre
quantity		number			int				Cantidad almacenada/60
code_value  	string          string 			Código alfanumérico característico/S73191A
is_published	boolean			bool			El producto se encuentra publicado o no/True
expiration		string			string			Fecha de vencimiento/ 12/04/2022
price			number			float64			Precio del producto / 50.15


Con la misma API que veníamos trabajando en clase, vamos a resolver los siguientes ejercicios.
Crear una ruta /productparams que tome todos los datos de la estructura de un producto por parámetro y lo devuelva en forma de JSON. Podemos seguir el siguiente ejemplo:

Insertar este último producto a la lista de productos y verificar si lo podemos tomar con la ruta /products/:id (más adelante veremos otro método para insertar datos en nuestras listas o bases de datos).
Se necesita un endpoint que devuelva una lista de productos que estén entre ciertas cantidades de stock. Por ejemplo: los productos que tengan entre 300 y 400 unidades. La ruta se llamará /searchbyquantity y debemos pasar los límites de las cantidades por parámetro.
Se necesita un endpoint que brinde el detalle de una compra. Por parámetro se deberá pasar el code_value del producto y la cantidad de unidades a comprar. El detalle de la compra deberá ser: nombre del producto, cantidad y precio total. La ruta se deberá llamar /buy. */