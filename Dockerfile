FROM golang:1.16

WORKDIR /app

COPY . /app

EXPOSE 8888

CMD [ "go", "run", "." ]