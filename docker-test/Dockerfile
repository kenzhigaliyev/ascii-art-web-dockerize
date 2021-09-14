FROM golang:latest

RUN mkdir /build
WORKDIR /build
# RUN export GO111MODULE=on


COPY ./ ./

RUN go build -o main .

CMD ["./main"]