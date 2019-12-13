FROM golang as builder
RUN mkdir /go/src/openprio_log
WORKDIR /go/src/openprio_log
COPY . .

RUN go get
RUN CGO_ENABLED=0 go build -o /go/bin/openprio_log

FROM alpine
COPY --from=builder /go/bin/openprio_log /app/openprio_log
WORKDIR /app
CMD ["/app/openprio_log"]