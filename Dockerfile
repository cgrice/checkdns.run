## Go build
FROM golang:alpine AS go

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /app/dnslookup
RUN go get -u github.com/DarthSim/overmind/v2
RUN which overmind

## Node build
FROM node:alpine AS node

WORKDIR /app

COPY . ./

RUN ls -la static/styles/main.css

RUN npm i && npm run build

## Caddy
FROM caddy:2-alpine

RUN apk add tmux

WORKDIR /app

COPY --from=go /app/dnslookup /app/dnslookup
COPY --from=go /go/bin/overmind /overmind
COPY --from=node /app/assets /app/assets

ADD Procfile /app/
ADD Caddyfile /app/

EXPOSE 80

CMD /overmind start

