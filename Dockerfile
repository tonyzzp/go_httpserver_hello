FROM golang:alpine as build
WORKDIR /build
COPY . .
RUN go build -o out .

FROM alpine
COPY --from=build /build/out server
EXPOSE 80
ENTRYPOINT [ "./server" ]
