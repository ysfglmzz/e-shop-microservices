FROM golang:1.18.3-alpine3.15 as builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY . .

RUN go mod download
RUN go build cmd/order/main.go

FROM alpine

COPY --from=builder /app/main /app/main
COPY --from=builder /app/resource/application.yaml /app/resource/application.yaml
COPY --from=builder /app/resource/queue.yaml /app/resource/queue.yaml


WORKDIR /app

CMD [ "./main" ]