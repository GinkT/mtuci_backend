FROM golang:alpine

WORKDIR /usr/local/go/src/mtuci_backend/cmd

COPY . /usr/local/go/src/mtuci_backend

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT ["/usr/local/go/src/mtuci_backend/cmd/main"]