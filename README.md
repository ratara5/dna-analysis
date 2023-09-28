# dna-analysis
Analizar una secuencia de ADN con Go
(Analize DNA sequence with Go)

## DOCKERIZAR. (DOCKERIZATION)
### CONSTRUIR UNA IMAGEN Y CORRER UN CONTENEDOR PARA DESARROLLO. (BUILD IMAGE AND RUN CONTAINER (DEV))
```
docker-compose up -d
```
### Ver LOGS DE CONTENEDOR PARA DESARROLLO. (See DEV CONTAINER LOGS)
```
docker logs -f dna-analysis-go-container
```

### CONSTRUIR UNA IMAGEN PARA PRODUCCIÓN. (BUILD IMAGE (PROD))
```
docker build --no-cache -t dna-analysis-go:lite -f Dockerfile.prod .
```
### Entonces CORRER UN CONTENEDOR PARA PRODUCCIÓN. (Then RUN CONTAINER (PROD))
```
docker run -d -p 8181:8080 --name dna-analysus-go-with-lite dna-analysis-go:lite
```

## USO DE LA API. (API USE)
```
localhost:8080/mutant/ 
POST {"dna":[L]}
```
El método POST `mutant`, permite verificar si una secuencia ADN es una secuencia ADN de mutante.
`[L]` es una lista de longitud `n`, cada elemento de la lista es un string de longitud `n`.  
`[L]->[M]`. 
`[M]` es una matriz `n*n` de caracteres, los caracteres válidos son: 'A' 'G' 'C' 'T'.

(The `mutant` POST method, allow verify if a dna sequence is a mutant dna sequence.
`[L]` is a list of lenght `n`, each list element is a `n` lenght string.  
`[L]->[M]`. 
`[M]` is a `n*n` char matrix, valid chars: 'A' 'G' 'C' 'T'.)


### Ejemplo del Body Request. (Body request example)
```
{
“dna”:["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}
```

Puerto 8080 para Desarrollo (Port 8080 for DEV) - Puerto 8181 para Producción (Port 8181 for PROD)
