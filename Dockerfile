FROM golang:1.20

RUN mkdir -p /github.com/sowjumn/probes

WORKDIR /github.com/sowjumn/probes

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o app .

CMD ["./app"]