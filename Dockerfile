FROM golang:1.20.5-alpine AS build

WORKDIR /src
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=/src \
    go get -d -v -t 
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=/src,rw \
    GOOS=linux go build -v -o /bin/website .
RUN chmod +x /bin/website

FROM scratch
ENV PORT=8080
COPY --from=build /bin/website /usr/local/bin/website
EXPOSE ${PORT}
CMD [ "website" ]