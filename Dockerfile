FROM golang:1.16.3

LABEL project="ASCII-ART-WEB-DOCKERIZE" \ 
authors="Adil and Zangar" \
link="https://git.01.alem.school/kenzhigaliyev/ascii-art-web-dockerize.git"

WORKDIR /dockerize

COPY go.mod .

RUN mkdir ascii
COPY ascii ascii/

RUN mkdir templates
COPY templates templates/

RUN mkdir fonts
COPY fonts fonts/

COPY main.go .

RUN go build -o main .

CMD ["./main"]