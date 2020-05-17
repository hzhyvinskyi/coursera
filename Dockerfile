############################################
# STEP 1 # Build a binary using alpine image
############################################

FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
# COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o app main.go

WORKDIR /dist

RUN cp /build/app .

#############################################
# STEP 2 # Build a small image for production
#############################################

FROM scratch

COPY --from=builder /dist/app /

ENTRYPOINT [ "/app" ]
