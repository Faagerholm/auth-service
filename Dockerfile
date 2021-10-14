# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as build

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /auth

FROM busybox
COPY --from=build /auth /auth
COPY --from=build /wait /wait
CMD /wait && /auth
