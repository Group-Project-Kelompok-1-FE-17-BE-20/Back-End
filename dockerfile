FROM golang:1.21-alpine

# membuat direktori folder
RUN mkdir /app

# set working direktory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable
RUN go build -o beapi

# run executable file
CMD ["./beapi"]