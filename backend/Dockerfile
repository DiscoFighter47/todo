# builder image
FROM golang:latest AS builder

# copy source code
WORKDIR /go/src/github.com/DiscoFighter47/todo/backend
COPY . .

# fetch dependencies
RUN go get -v -t -d ./...

# build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARC=amd64 go build -a -installsuffix cgo -o binary .


# base image
FROM alpine:latest

# Security related package
RUN apk --no-cache add ca-certificates

# copy the binary
COPY --from=builder /go/src/github.com/DiscoFighter47/todo/backend/binary .

# run the binary
ENTRYPOINT [ "./binary" ]