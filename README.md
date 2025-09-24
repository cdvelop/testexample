# 📘 Certamen de Testing WhiteBox en Go

Cada alumno debe:

1. Crear una rama a partir de `test/certamen/` repositorio: `github.com/cdvelop/testexample` con el formato:

   ```
   test/certamen/<apellido-nombre>/caso_XXX
   ```
2. Completar la función indicada con el nombre de archivo: `nombre_funcion.go`.

3. Escribir un test en Go con **al menos 10 casos distintos** con el nombre de archivo: `nombre_funcion_test.go`.

Ejemplo de estructura mínima de test (debes completarlo):

```go
func TestAdd(t *testing.T) {
	tests := []struct {
		name string // nombre del caso
		a, b int // datos de entrada
		expectativa int // expectativa
	}{
		{"Caso 1 expectativa correcta", 1, 2, 3},
		// TODO: agregar al menos 10 casos
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expectativa {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expectativa)
			}
		})
	}
}
```


4. Subir la rama al repositorio remoto.

---

## ✅ Caso 001 – Suma de positivos en un slice

```go
// SumPositives retorna la suma de los enteros positivos dentro del slice nums.
// Ejemplo: SumPositives([]int{-1, 2, 3, 0}) == 5
// Si no hay positivos, debe retornar 0.
func SumPositives(nums []int) int {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 002 – Máximo de dos números

```go
// Max retorna el mayor de los dos números a y b.
// Ejemplo: Max(3, 7) == 7
func Max(a, b int) int {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 003 – Es par

```go
// IsEven retorna true si n es par, false en caso contrario.
// Ejemplo: IsEven(4) == true
func IsEven(n int) bool {
	// TODO: implementar
	return false
}
```

---

## ✅ Caso 004 – Factorial

```go
// Factorial retorna el factorial de n (n!).
// Se asume que n >= 0.
// Ejemplo: Factorial(5) == 120
func Factorial(n int) int {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 005 – Palíndromo

```go
// IsPalindrome retorna true si la palabra es un palíndromo.
// Un palíndromo es una palabra que se lee igual de izquierda a derecha y de derecha a izquierda.
// Ejemplo: IsPalindrome("radar") == true
// Ejemplo adicional: IsPalindrome("level") == true
func IsPalindrome(s string) bool {
	// TODO: implementar
	return false
}
```

---

## ✅ Caso 006 – Reversa de string

```go
// Reverse retorna el string invertido.
// Ejemplo: Reverse("hola") == "aloh"
func Reverse(s string) string {
	// TODO: implementar
	return ""
}
```

---

## ✅ Caso 007 – Promedio

```go
// Average retorna el promedio de una lista de enteros.
// Si la lista está vacía, debe retornar 0.
// Ejemplo: Average([]int{2, 4, 6}) == 4.0
func Average(nums []int) float64 {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 008 – Contar vocales

```go
// CountVowels retorna cuántas vocales (a,e,i,o,u) hay en el string.
// Ejemplo: CountVowels("hola mundo") == 4
func CountVowels(s string) int {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 009 – Buscar en slice

```go
// Contains retorna true si el número target está dentro del slice nums.
// Ejemplo: Contains([]int{1,2,3}, 2) == true
func Contains(nums []int, target int) bool {
	// TODO: implementar
	return false
}
```

---

## ✅ Caso 010 – Fibonacci

```go
// Fibonacci retorna el n-ésimo número de Fibonacci.
// Ejemplo: Fibonacci(0)=0, Fibonacci(1)=1, Fibonacci(5)=5
// Ejemplo adicional: Fibonacci(7) == 13
func Fibonacci(n int) int {
	// TODO: implementar
	return 0
}
```

---

📊 **Pauta de Evaluación para todos los casos (40 pts c/u)**

| Criterio                     | Descripción                           | Puntaje    |
| ---------------------------- | ------------------------------------- | ---------- |
| Creación de rama correcta    | Rama creada con el formato solicitado | 5 pts      |
| Implementación de la función | Función implementada correctamente    | 10 pts     |
| Estructura del test          | Table-driven test en Go               | 5 pts      |
| Cantidad de casos            | Al menos 10 casos distintos           | 10 pts     |
| Variedad de casos            | Cobertura de escenarios diversos      | 5 pts      |
| Calidad del test             | Casos bien nombrados, errores claros  | 5 pts      |
| **Total**                    |                                       | **40 pts** |

---


