# IExplicacion y solcion del error en pruebas Go

## Contexto
Se nos asignó la tarea de corregir un error en las pruebas unitarias de un proyecto en Go.  
El problema original ocurría porque el mensaje de error devuelto por la función `NewDatabaseEngine` **no coincidía** con el mensaje esperado en el test.

## Problema detectado
El test esperaba el mensaje:

db not found

Pero la función devolvía:

NO existe esta db: post

Esto provocaba que la prueba fallara con el siguiente resultado:

Error esperado: db not found, obtenido: NO existe esta db:%!(EXTRA string=post)

## Solución 
Se modificó la función `NewDatabaseEngine` para que, en el caso de recibir un tipo de base de datos no soportado, retorne exactamente el mensaje que la prueba espera:

```go
default:
    return nil, errors.New("db not found")

