FROM golang:1.18.3-alpine3.15 as builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY . .

RUN go mod download
RUN go build ./main.go

FROM alpine

COPY --from=builder /app/main /app/main
COPY --from=builder /app/resource/routes.yaml /app/resource/routes.yaml


WORKDIR /app

CMD [ "./main" ]