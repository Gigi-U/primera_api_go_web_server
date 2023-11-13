package get


import (
	"context"
	"fmt"
)

func Context() {
	ctx := context.Background()

	newCtx := addValue(ctx)
	s:= newCtx.Value("bootcamp")
	fmt.Println(s)
}

func addValue(ctx context.Context)context.Context{
	return context.WithValue(ctx,"bootcamp","Go")
}

// funciones de la interfaz context | DEADLINE: retorna un booleano y un tiempo. lo q hace es finalizar una tarea dado x tiempo q se le asigne y retorna un bool indicando si terminó o no. DONE: nos dice si se cumplio con una serie de instrucciones dentro de una goroutine. ERR (ya sabemos que hace) . VALUE:  permite pasar valores entre funciones, a traves del contexto.

//^ algo importante del diseño detras del package context es que TODO DEVUELVE UNA NUEVA ESTRUCTURA CONTEXT.CONTEXT . Esto significa que debemos recordar de trabajar con el valor devuelto de las operaciones y posiblemente ANULAR contextos antiguos con otros nuevos.
//& ESTE ES UN DISEÑO CLAVE EN LA INMUTABILIDAD.