FROM golang:1.19


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get install -y postgresql-client
#
RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o techtask ./cmd/main.go

CMD ["./techtask"]
EXPOSE 8000

#
#COPY ./ ./
#
#ENV GO111MODULE=on
#RUN mkdir src/app
#WORKDIR /src/app
#
#COPY go.mod .
#COPY . .
#
#
#RUN go build -o main
#CMD ["/app/main"]


