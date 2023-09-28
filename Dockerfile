FROM golang
WORKDIR /app
RUN go mod init dna-analysis
RUN apt-get update && apt-get install -y ca-certificates
RUN apt-get install -y tzdata
COPY . .
CMD ["go", "run", "."]