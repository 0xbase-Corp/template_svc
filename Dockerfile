FROM cgr.dev/chainguard/go:latest as src

WORKDIR /src
COPY . .

# migrate

FROM src as build-migrate
RUN go build -o ./build/migrate ./cmd/migrate

FROM cgr.dev/chainguard/glibc-dynamic:latest as migrate

COPY --from=build-migrate /src/build/migrate /bin/migrate
COPY --from=build-migrate /src/app.env /
CMD ["/bin/migrate"]

# API

FROM src as build-api
RUN go build -o ./build/api ./cmd/0xbase

FROM cgr.dev/chainguard/glibc-dynamic:latest as api

COPY --from=build-api /src/build/api /bin/api
COPY --from=build-api /src/app.env /

# Expose port
EXPOSE 5050

CMD ["/bin/api"]