FROM golang:latest

WORKDIR /home
COPY . /home

RUN cd /home 
RUN env GOOS=linux GOARCH=amd64 go  build -o todo

CMD ["/home/todo"]