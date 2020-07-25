FROM golang:latest AS builder

RUN mkdir /go/src/subscriber
WORKDIR /go/src/subscriber
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o subscriber .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/subscriber/ .
CMD ["./subscriber"]