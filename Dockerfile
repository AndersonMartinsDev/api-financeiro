# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-buster AS build

WORKDIR /usr/src/app

COPY . . 

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-api-financeiro

# ## Deploy
# FROM gcr.io/distroless/base-debian10

# COPY --from=build /go-api-financeiro /go-api-financeiro

EXPOSE 5000

ENTRYPOINT ["/go-api-financeiro"]
