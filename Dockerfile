FROM golang:1.21.3

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o basiccrud

CMD ./basiccrud