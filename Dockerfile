FROM golang:1.19.6-bullseye

ARG CMD

WORKDIR /app

COPY . ./

RUN rm -rf client
RUN go mod download && go mod verify
RUN go build -v -o /app/pulpcalc .

EXPOSE 8080

CMD [$CMD]