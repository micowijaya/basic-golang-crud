FROM golang:1.21.3 AS build

WORKDIR /app

COPY . /app

RUN go mod download
RUN go mod verify

RUN go build -o basiccrud

FROM golang:1.21.3

WORKDIR /app

COPY --from=build app/basiccrud .

CMD ./basiccrud