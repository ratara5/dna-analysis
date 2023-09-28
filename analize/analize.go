package analize

import (
	"sync"
)

// CuatroEnFila comprueba la existencia de repeticiones de 4 en horizontal
func CuatroEnFila(matriz [][]string, canal chan int, wg *sync.WaitGroup, goroutinesTerminadas *int) {
	defer wg.Done()
	n := len(matriz)
	count := 0
	// Verificar filas
	for i := 0; i < n; i++ {
		fila := matriz[i]
		for j := 0; j <= len(fila)-4; j++ {
			if fila[j] == fila[j+1] && fila[j] == fila[j+2] && fila[j] == fila[j+3] {
				count++
			}
			canal <- count
		}
	}
	*goroutinesTerminadas++
}

// CuatroEnColumna comprueba la existencia de repeticiones de 4 en vertical
func CuatroEnColumna(matriz [][]string, canal chan int, wg *sync.WaitGroup, goroutinesTerminadas *int) {
	defer wg.Done()
	n := len(matriz)
	// Verificar columnas
	count := 0
	for j := 0; j < n; j++ {
		for i := 0; i <= n-4; i++ {
			if matriz[i][j] == matriz[i+1][j] && matriz[i][j] == matriz[i+2][j] && matriz[i][j] == matriz[i+3][j] {
				count++
			}
			canal <- count
		}
	}
	*goroutinesTerminadas++
}

// CuatroEnDiagonalDesc comprueba la existencia de repeticiones de 4 en diagonal descendente
func CuatroEnDiagonalDesc(matriz [][]string, canal chan int, wg *sync.WaitGroup, goroutinesTerminadas *int) {
	defer wg.Done()
	n := len(matriz)
	count := 0
	// Verificar diagonales descendentes
	for i := 0; i <= n-4; i++ {
		for j := 0; j <= n-4; j++ {
			if matriz[i][j] == matriz[i+1][j+1] && matriz[i][j] == matriz[i+2][j+2] && matriz[i][j] == matriz[i+3][j+3] {
				count++
			}
			canal <- count
		}
	}
	*goroutinesTerminadas++
}

// CuatroEnDiagonalAsc comprueba la existencia de repeticiones de 4 en diagonal ascendente
func CuatroEnDiagonalAsc(matriz [][]string, canal chan int, wg *sync.WaitGroup, goroutinesTerminadas *int) {
	defer wg.Done()
	n := len(matriz)
	count := 0
	// Verificar diagonales ascendentes
	for i := 3; i < n; i++ {
		for j := 0; j <= n-4; j++ {
			if matriz[i][j] == matriz[i-1][j+1] && matriz[i][j] == matriz[i-2][j+2] && matriz[i][j] == matriz[i-3][j+3] {
				count++
			}
			canal <- count
		}
	}
	*goroutinesTerminadas++
}