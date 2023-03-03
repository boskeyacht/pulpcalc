FROM golang:1.19.6-bullseye


ARG CMD

WORKDIR /app

COPY . ./

RUN go mod download && go mod verify
RUN go build -v -o pulpcalc .

EXPOSE 7687

ENTRYPOINT $CMD