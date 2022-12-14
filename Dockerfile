FROM golang:1.19

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o metrics-app ./cmd/main.go

CMD ["./metrics-app"]