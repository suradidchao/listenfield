# builder image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o listenfield app/main.go


# generate clean, final image for end users
FROM alpine:3.11.3
WORKDIR /app
COPY --from=builder /build/app/config config
COPY --from=builder /build/listenfield .

# arguments that can be overridden
CMD [ "./listenfield"]