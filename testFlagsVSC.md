

## configuraciones de test en visual estudio code para go
```json

{
    "go.coverOnSave": true,// ejecuta las pruebas cuando se guarda el archivo
    "go.testFlags": [
        "-v",// que el test se muestra con detalle verboso
        //"-failfast", el test al fallar se detiene
        "-test.v" // no me recuerdo
    ],
    "go.coverOnSingleTest": true,
    "go.coverageOptions": "showUncoveredCodeOnly",
   
}
```