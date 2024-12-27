FROM golang:1.23.4

WORKDIR /cmd/web

COPY . .
RUN go build ./cmd/web

EXPOSE 4000

CMD ["./web"]
