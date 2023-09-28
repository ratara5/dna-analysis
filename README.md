# dna-analysis
### Analize DNA sequence with Go

## BUILD IMAGE AND RUN CONTAINER (DEV)
```
docker-compose up -d
```

## BUILD IMAGE (PROD)
```
docker build --no-cache -t dna-analysis-go:lite -f Dockerfile.prod .
```
## Then RUN CONTAINER (PROD)
```
docker run -d -p 8181:8080 --name dna-analysus-go-with-lite dna-analysis-go:lite
```

## API USE
```
localhost:8080/mutant/ 
POST {"dna":[L]}
```
The `mutant` POST method, allow verify if a dna sequence is a mutant dna sequence.
`[L]` is a list of lenght `n`, each list element is a `n` lenght string 
`[L]->[M]`
`[M]` is a `n*n` char matrix, valid chars: 'A' 'G' 'C' 'T'


### Body request example
```
{
“dna”:["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}
```

### Port 8080 for DEV
### Port 8181 for PROD