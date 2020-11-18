FROM golang:1.15.3 as builder

WORKDIR /go/src/github.com/DTS-STN/benefit-service
RUN mkdir -p /go/src/github.com/DTS-STN/benefit-service
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service .

FROM scratch
COPY --from=builder /go/src/github.com/DTS-STN/benefit-service/service .
ENTRYPOINT ["./service"]