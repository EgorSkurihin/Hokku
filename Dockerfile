FROM golang:latest

WORKDIR /

COPY . .

RUN go mod download
RUN go build -o .

CMD ["./Hokku"]