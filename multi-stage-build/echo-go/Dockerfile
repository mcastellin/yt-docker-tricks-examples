FROM golang:1.16 as builder

WORKDIR /go/src/echo

COPY go.mod *.go ./

RUN go get -d -v
RUN go test -v

# need to use the CGO_ENABLED=0 to disably dynamic linking used by the net library
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=builder /go/src/echo/echoapp /echoapp

EXPOSE 8081

CMD ["/echoapp"]