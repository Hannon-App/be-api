FROM golang:1.20-alpine
RUN mkdir /app
WORKDIR /app 
COPY ./ /app
RUN go mod tidy 

RUN go build -o be-api 
CMD ["./be-api"]