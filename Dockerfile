FROM golang:latest AS go

WORKDIR /app
ENV GOFLAGS=-buildvcs=false
COPY go.mod .
COPY go.sum .
COPY setup.sh .
RUN ./setup.sh
RUN go mod download

FROM go AS app
WORKDIR /app
COPY . .
COPY ./configs/air.toml ./.air.toml

# ENTRYPOINT ["fiber", "dev"]
ENTRYPOINT ["air", "-c", "configs/air.toml"]