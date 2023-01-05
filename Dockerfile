#Step 1: Cache modules
FROM golang:1.18.3-alpine3.16 AS build

WORKDIR /app/

#Install Dependencies
ADD go.mod go.sum ./
RUN go mod download

ADD . .

#Build application
RUN CGO_ENABLED=0 GOOS=linux go build -v cmd/server/main.go

#Step 2: Copy artifacts on clean image
FROM alpine
COPY --from=build /app/main /app/main
COPY --from=build /app/conf/config.yaml /app/config/

EXPOSE 80
CMD /app/main