# 📘 Certamen de Testing WhiteBox en Go

Cada alumno debe:

1. Crear una rama a partir de `test/certamen/` repositorio: `github.com/cdvelop/testexample` con el formato:

   ```
   <apellido-nombre>
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
		{"Caso 1 entradas 1 y 2 expectativa 3 OK", 1, 2, 3},
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

## ✅ Caso 001 – Suma de positivos en un slice (5 pts.)

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

## ✅ Caso 002 – Máximo de dos números (10 pts.)

```go
// Max retorna el mayor de los dos números a y b.
// Ejemplo: Max(3, 7) == 7
func Max(a, b int) int {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 003 – Es par (5 pts.)

```go
// IsEven retorna true si n es par, false en caso contrario.
// Ejemplo: IsEven(4) == true
func IsEven(n int) bool {
	// TODO: implementar
	return false
}
```

---

## ✅ Caso 004 – Factorial (15 pts.)

```go
// Descripción:
// El factorial de un entero no negativo n (denotado n!) es el producto
// de todos los enteros positivos desde 1 hasta n. Por convención:
//   0! = 1
// Y para n >= 1:
//   n! = n * (n-1)!
//
// Notas importantes:
// - Se asume n >= 0; los valores negativos no están definidos para factorial
//   en este ejercicio.
// - El factorial crece muy rápido. Para n relativamente grandes puede
//   producirse desbordamiento (overflow) en tipos enteros de 32 o 64 bits.
//   Si se espera n grande se debe usar un tipo con mayor capacidad (por
//   ejemplo uint64) o big.Int de la biblioteca estándar.
//
// Ejemplos:
// Factorial(0) == 1
// Factorial(5) == 120   // 5*4*3*2*1 = 120
// Factorial(6) == 720   // 6*5*4*3*2*1 = 720
// Factorial(10) == 3628800
func Factorial(n int) int {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 005 – Palíndromo (15 pts.)

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

## ✅ Caso 006 – Reversa de string (10 pts.)

```go
// Reverse retorna el string invertido.
// Ejemplo: Reverse("hola") == "aloh"
func Reverse(s string) string {
	// TODO: implementar
	return ""
}
```

---

## ✅ Caso 007 – Promedio (10 pts.)

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

## ✅ Caso 008 – Contar vocales (10 pts.)

```go
// CountVowels retorna cuántas vocales (a,e,i,o,u) hay en el string.
// Ejemplo: CountVowels("hola mundo") == 4
func CountVowels(s string) int {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 009 – Buscar en slice (5 pts.)

```go
// Contains retorna true si el número target está dentro del slice nums.
// Ejemplo: Contains([]int{1,2,3}, 2) == true
func Contains(nums []int, target int) bool {
	// TODO: implementar
	return false
}
```

---

## ✅ Caso 010 – Fibonacci (15 pts.)

```go
// Descripción:
// La sucesión de Fibonacci es una secuencia donde cada término a partir
// del segundo es la suma de los dos anteriores. Usando la convención con
// índice 0 (0-based):
//   Fibonacci(0) = 0
//   Fibonacci(1) = 1
// Para n >= 2:
//   Fibonacci(n) = Fibonacci(n-1) + Fibonacci(n-2)
//
// Notas importantes:
// - El índice aquí es 0-based (Fibonacci(0)=0).
// - Para esta práctica supongamos n razonablemente pequeño (por ejemplo
//   n <= 45) para que el resultado quepa en un int de 32 bits.
//
// Ejemplos:
// Fibonacci(0) == 0
// Fibonacci(1) == 1
// Fibonacci(5) == 5    // secuencia: 0,1,1,2,3,5
// Fibonacci(7) == 13   // secuencia: 0,1,1,2,3,5,8,13
// Fibonacci(10) == 55
// Fibonacci(12) == 144
func Fibonacci(n int) int {
	// TODO: implementar
	return 0
}
```

---

📊 **Pauta de Evaluación (sistema por descuentos)**

Los alumnos deben ganar los puntos resolviendo los tests y cumpliendo la pauta. El puntaje máximo disponible al que pueden optar es 50 pts. no importa que resuelvan mas.

### Tabla criterios  descuentos(cada incumplimiento descuenta 5 pts):

| Ítem (descarta 5 pts si no cumplido)     | Descripción breve                                |
| ---------------------------------------- | ------------------------------------------------ |
| Rama con formato correcto                | No creó la rama con el formato solicitado        |
| Implementación                        | La función no compila o está claramente incorrecta |
| Test table-driven                      | No se usó una estructura table-driven en el test |
| ≥ 10 casos en el test                   | Menos de 10 casos en el archivo de test         |
| Variedad de escenarios                  | Casos muy repetitivos o no cubren bordes        |
| Calidad del test                        | Nombres/mensajes poco claros |
| Implementaciones duplicadas            | Si hay 2 (o más) implementaciones del mismo caso con los mismos números, se descontarán 5 pts por cada ítem repetido. |

Notas importantes:
- Los criterios se aplican de forma independiente.
- Esta pauta se aplica a cada uno de los casos individualmente.

---


