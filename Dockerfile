FROM golang:1.16-buster as build

WORKDIR /go/src/app
COPY . .
RUN go build -o bin/challenge

FROM golang:1.16-buster 

USER root

COPY --from=build /go/src/app/bin/challenge /usr/bin/challenge

ENTRYPOINT [ "/usr/bin/challenge" ]
