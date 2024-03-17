FROM golang:alpine AS builder

WORKDIR /usr

ADD go.mod .

COPY . .

RUN go build -o ./bin/application cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/bin/application .
COPY config/local.yaml .

ENV CONFIG_PATH=./local.yml

ENTRYPOINT ["/application"]