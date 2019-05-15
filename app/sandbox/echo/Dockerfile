# build binary
FROM golang:alpine as builder
WORKDIR /go/src/github.com/mokichi13/warabi/app/sandbox/echo
COPY . .
RUN go build

# runtime image
FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/mokichi13/warabi/app/sandbox/echo/echo /echo
EXPOSE 8080
ENTRYPOINT [ "/echo" ]