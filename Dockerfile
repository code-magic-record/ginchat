FROM golang:1.17

WORKDIR /app

COPY . /app

EXPOSE 8080

CMD [ "go", "run", "." ]