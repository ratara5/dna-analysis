package analize

import (
	"fmt"
	"sync"
	"time"
	"github.com/ratara5/dna-analysis/util"
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

//IsMutant Verifica si una secuencia de dna contenida en una matriz de nxn corresponde a la secuencia de un mutante
func IsMutant(dna []string) bool {
	//Generación de Matriz de DNA
	resultadoMatriz := util.CrearMatriz(dna)
	//-----------------------------------------------------------------------------------------------------
	//TODO: Validar cadenas: que la longitud sea válida y que no haya caracteres diferentes de 'CAGT'
	//TODO: Una secuencia de 5 caracteres repetidos cuenta como dos secuencias de 4 caracteres repetidos?
	//-----------------------------------------------------------------------------------------------------

	//Analisis de secuencia
	///Canal para recibir conteo de cadenas de 4
	resultadosCanal := make(chan int, 4)
	///WaitGroup para esperar goroutines de revisión y conteo de cadenas de 4
	var wg sync.WaitGroup
	///Contador de rutinas terminadas
	var goroutinesTerminadas int

	///Llamado y adición a WaitGroup de goroutines de revisión y conteo de cadenas de 4
	wg.Add(1)
	go CuatroEnFila(resultadoMatriz, resultadosCanal, &wg, &goroutinesTerminadas)
	wg.Add(1)
	go CuatroEnColumna(resultadoMatriz, resultadosCanal, &wg, &goroutinesTerminadas)
	wg.Add(1)
	go CuatroEnDiagonalDesc(resultadoMatriz, resultadosCanal, &wg, &goroutinesTerminadas)
	wg.Add(1)
	go CuatroEnDiagonalAsc(resultadoMatriz, resultadosCanal, &wg, &goroutinesTerminadas)

	///Comprobación continua de conteo de cadenas de 4
	var sumaResultados int
	for goroutinesTerminadas < 4 {
		select {
		case resultado := <-resultadosCanal:
			sumaResultados += resultado
			fmt.Println("Resultado:", resultado)
			if sumaResultados == 2 {
				//fmt.Println("La suma de resultados es igual a ", sumaResultados, ", el programa termina.")
				return true
			}
		case <-time.After(time.Millisecond * 100):
			// Comprobar en intervalos de tiempo si se cumple la condición
			if sumaResultados == 2 {
				//fmt.Println("La suma de resultados es igual a ", sumaResultados, ", el programa termina.")
				return true
			}
		}
	}

	///Espera terminación de todas las goroutines
	wg.Wait()
	fmt.Println("Todas las GoRoutines han terminado!!!\nLa suma de resultados es igual a ", sumaResultados, ", el programa termina.")
	return false

}
