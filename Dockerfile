FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/site .

FROM alpine:3.19
WORKDIR /app
COPY --from=build /bin/site /bin/site
COPY --from=build /src/templates ./templates
COPY --from=build /src/static ./static
EXPOSE 3000
WORKDIR /app
ENTRYPOINT ["/bin/site"]