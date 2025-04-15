FROM golang:1.20 AS server_builder
WORKDIR /daneshjooyar
COPY . .
RUN go build -o daneshjooyar
CMD ["./daneshjooyar"]






