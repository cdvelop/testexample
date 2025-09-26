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

## ✅ Caso 011 – Contar palabras (5 pts.)

```go
// WordCount retorna cuántas palabras hay en el string s.
// Una palabra se define como una secuencia de caracteres separados por espacios.
// Ejemplo: WordCount("hola mundo go") == 3
func WordCount(s string) int {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 012 – Número primo (10 pts.)

```go
// Descripción:
// Un número primo es aquel mayor que 1 que solo tiene dos divisores: 1 y él mismo.
// Es decir, no puede dividirse exactamente (sin residuo) por ningún otro número.
//
// Ejemplos de números primos: 2, 3, 5, 7, 11, 13...
// Ejemplos de no primos: 1 (no cuenta como primo), 4 (2*2), 6 (2*3), 8 (2*4).
//
// IsPrime retorna true si n es un número primo.
// Ejemplo: IsPrime(7) == true
// Ejemplo: IsPrime(10) == false
func IsPrime(n int) bool {
	// TODO: implementar
	return false
}

```

---

## ✅ Caso 013 – Ordenar números (10 pts.)

```go
// SortInts retorna una copia del slice nums ordenado en forma ascendente.
// Ejemplo: SortInts([]int{3,1,2}) == []int{1,2,3}
func SortInts(nums []int) []int {
	// TODO: implementar
	return nil
}
```

---

## ✅ Caso 014 – Contar ocurrencias (5 pts.)

```go
// CountOccurrences retorna cuántas veces aparece target en el slice nums.
// Ejemplo: CountOccurrences([]int{1,2,2,3,2}, 2) == 3
func CountOccurrences(nums []int, target int) int {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 015 – Convertir Celsius a Fahrenheit (5 pts.)

```go
// CelsiusToFahrenheit convierte grados Celsius a Fahrenheit.
// Fórmula: F = C * 9/5 + 32
// Ejemplo: CelsiusToFahrenheit(0) == 32
func CelsiusToFahrenheit(c float64) float64 {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 016 – Números pares hasta N (10 pts.)

```go
// EvenNumbersUntil retorna un slice con todos los números pares desde 0 hasta n (incluido).
// Ejemplo: EvenNumbersUntil(6) == []int{0,2,4,6}
func EvenNumbersUntil(n int) []int {
	// TODO: implementar
	return nil
}
```

---

## ✅ Caso 017 – Buscar máximo en slice (5 pts.)

```go
// MaxInSlice retorna el valor máximo dentro de nums.
// Si el slice está vacío, debe retornar 0.
// Ejemplo: MaxInSlice([]int{3,7,2}) == 7
func MaxInSlice(nums []int) int {
	// TODO: implementar
	return 0
}
```

---

## ✅ Caso 018 – Palabra más larga (10 pts.)

```go
// LongestWord retorna la palabra más larga en el string s.
// En caso de empate, retorna la primera encontrada.
// Ejemplo: LongestWord("go es divertido") == "divertido"
func LongestWord(s string) string {
	// TODO: implementar
	return ""
}
```

---

## ✅ Caso 019 – Tabla de multiplicar (5 pts.)

```go
// MultiplicationTable retorna un slice con los resultados de la tabla del número n, desde 1 hasta 10.
// Ejemplo: MultiplicationTable(3) == []int{3,6,9,12,15,18,21,24,27,30}
func MultiplicationTable(n int) []int {
	// TODO: implementar
	return nil
}
```

---

## ✅ Caso 020 – Invertir slice de enteros (10 pts.)

```go
// ReverseSlice retorna un nuevo slice con los elementos de nums en orden inverso.
// Ejemplo: ReverseSlice([]int{1,2,3}) == []int{3,2,1}
func ReverseSlice(nums []int) []int {
	// TODO: implementar
	return nil
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


