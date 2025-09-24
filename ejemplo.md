

1. **Cree su rama** en GitHub con el formato:

   ```
   test/certamen/nombre_alumno/caso_001
   ```
2. **Implemente la función incompleta**.
3. **Complete el test**, generando al menos **10 casos distintos**.

Te muestro cómo quedaría el **primer caso** redactado como lo pedirías a tus alumnos:

---

## 📌 Caso 001 – Suma de dos números

**Instrucciones:**

1. Desde el repositorio `github.com/cdvelop/testexample`, crea una rama a partir de la rama base `test/certamen/` con el siguiente formato:

   ```
   test/certamen/<tu_nombre>/caso_001
   ```

2. Dentro del paquete asignado, implementa la siguiente función incompleta:

```go
// Add debe retornar la suma de los dos enteros a y b.
func Add(a, b int) int {
	// TODO: implementar
	return 0
}
```

3. Crea un archivo de pruebas `add_test.go` donde escribas **al menos 10 casos de prueba diferentes** para verificar el correcto funcionamiento de `Add`.

   * Incluye casos con números positivos.
   * Incluye casos con cero.
   * Incluye casos con números negativos.
   * Incluye casos mixtos (positivo + negativo).
   * (Puedes inventar más, lo importante es llegar a 10).

4. El test debe seguir la estructura de pruebas con **table-driven tests** en Go, pero eres libre de agregar más validaciones si lo consideras necesario.

Ejemplo de estructura mínima de test (debes completarlo):

```go
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		// TODO: agregar al menos 10 casos
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
```

5. Una vez completado, sube tu rama al repositorio remoto.

---

