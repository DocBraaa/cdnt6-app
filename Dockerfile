FROM golang:1.12.5-alpine as app
LABEL maintainer="Vincent Letourneau <vincent@nanoninja.com>"
WORKDIR /go/src/app
COPY . .

RUN GOOS=linux go build -o main .

FROM scratch
COPY --from=app /go/src/app .
CMD ["./main"]
