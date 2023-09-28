package util

import (
	"math/rand"
	"strings"
	"time"
)

// GenerarDNA Genera una secuencia de ADN aleatoria como una lista de n elementos, todos los cuales comparten un tamaño de longitud caracteres
func GenerarDNA(n, longitud int, caracteresEntrada string) []string {
	palabras := []string{}
	rand.Seed(time.Now().UnixNano()) //Deprecated

	if len(caracteresEntrada) != 4 {
		return []string{"La entrada debe contener exactamente cuatro caracteres."}
	}

	for i := 0; i < n; i++ {
		var palabra strings.Builder
		for j := 0; j < longitud; j++ {
			palabra.WriteByte(caracteresEntrada[rand.Intn(4)])
		}
		palabras = append(palabras, palabra.String())
	}

	return palabras
}

// CrearMatriz Crea una matriz a partir de una lista. La cantidad de filas de la matriz es igual a la cantidad de elementos de la lista; la cantidad de columnas de la matriz es igual a la cantidad de caracteres de cada elemento, todos los elementos de la lista deben tener la misma cantidad de caracteres.
func CrearMatriz(lista []string) [][]string {
	n := len(lista)
	matriz := make([][]string, n)

	for i, fila := range lista {
		// Dividir el string en caracteres para obtener las columnas
		columnas := strings.Split(fila, "")
		matriz[i] = columnas
	}

	return matriz
}
