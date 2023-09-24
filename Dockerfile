FROM golang:1.20.5-alpine AS build

RUN mkdir /src
ADD . /src
ADD ./go.mod /src
ADD ./go.sum /src

WORKDIR /src
RUN go get -d -v -t 
RUN GOOS=linux go build -v -o website .
RUN chmod +x website

FROM scratch
COPY --from=build /src/website /usr/local/bin/website
EXPOSE 8080
CMD [ "website" ]