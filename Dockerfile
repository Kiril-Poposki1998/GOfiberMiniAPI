FROM golang:1.20.5-alpine AS build

WORKDIR /src
ADD ./views /src/views

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=/src,rw \
    go get -d -v -t 
    
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=/src,rw \
    GOOS=linux go build -v -o /bin/website .
    
RUN chmod +x /bin/website

FROM scratch
ENV APP_PORT=8080
COPY --from=build /bin/website /usr/local/bin/website
COPY --from=build /src/views /src/views
EXPOSE ${APP_PORT}
CMD [ "website" ]