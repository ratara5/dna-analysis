package analize

import (
	"sync"
	"time"
	"github.com/ratara5/dna-analysis/util"
)

//IsMutant Verifica si una secuencia de dna contenida en una matriz de nxn corresponde a la secuencia de un mutante
func IsMutant(dna []string) bool {
	
	//Generación de Matriz de DNA
	resultadoMatriz := util.CrearMatriz(dna)

	//--------------------------------------------------------------------------------------------------------
	//TODO: Validar cadenas: que la longitud sea válida (igual a la de la lista) y que no haya caracteres diferentes de 'C' 'A' 'G' 'T'
	//TODO: Una secuencia de 5 caracteres repetidos cuenta como dos secuencias de 4 caracteres repetidos?
	//--------------------------------------------------------------------------------------------------------

	//Analisis de secuencia
	///Canal para recibir conteo de cadenas de 4
	resultadosCanal := make(chan int, 4)
	///WaitGroup para esperar goroutines de revisión y conteo de cadenas de 4
	var wg sync.WaitGroup
	///Contador de rutinas terminadas
	var goroutinesTerminadas int

	///Llamada y adición a WaitGroup de goroutines de revisión y conteo de cadenas de 4
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
			if sumaResultados == 2 {
				return true
			}
		case <-time.After(time.Millisecond * 100):
			// Comprobar en intervalos de tiempo si se cumple la condición
			if sumaResultados == 2 {
				return true
			}
		}
	}

	///Espera terminación de todas las goroutines
	wg.Wait()
	return false

}