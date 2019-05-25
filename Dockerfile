FROM golang:1.12-alpine as builder

RUN apk add --no-cache ca-certificates git

ENV PROJECT frontend
WORKDIR /go/src/$PROJECT

ENV GO111MODULE on
COPY . .
RUN go install .

FROM alpine as release
RUN apk add --no-cache ca-certificates \
    busybox-extras net-tools bind-tools
WORKDIR /frontend
COPY --from=builder /go/bin/frontend /app/frontend

EXPOSE 8000
ENTRYPOINT ["/app/frontend"]