# Cloud-go-ref
## Golang reference app for 101 demonstration of Golang, Docker, AWS bits, Terraform, etc...

### Compile to binary:

Format the code first:

go fmt ./...

Clean the code next:

go clean

Binary for linux:

GOOS=linux CGO_ENABLED=0 go build -a

Binary for mac:

GOOS=darwin CGO_ENABLED=0 go build -a

### Build Docker image:

docker build -t cloud-go-ref  .

### Run Docker image:

docker run  -p 80:80 -i -t cloud-go-ref
