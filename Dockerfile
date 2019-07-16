FROM golang:latest AS build-env

RUN mkdir /app
WORKDIR /app
COPY . .
ENV CGO_ENABLED 0
RUN go build

FROM alpine:latest

COPY --from=build-env /app/tezTracker /

CMD ["/tezTracker"]
