FROM golang as build
COPY ./ /app
WORKDIR /app
RUN go build

FROM gcr.io/distroless/base
COPY --from=build /app/udp-forward /usr/bin/udp-forward
ENTRYPOINT ["udp-forward"]
