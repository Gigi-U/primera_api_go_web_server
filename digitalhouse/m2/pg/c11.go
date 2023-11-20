package pg

/* método POST: se usa principalmente p/ enviar datos al servidor.
La forma en que estan codificados los datos enel cuerpo de la petición se indica en el HEADER CONTENT TYPE.
Por ejemplo application/JSON.
El contenido de lo que enviamos al servidor se coloca en el BODY de la request. Los parámetros no se exponen en la url a diferencia del mpetodo GET. 

Si la respuesta es positiva retorna valores asociados al rango de los 200, identificando un status code de confirmación
*/

