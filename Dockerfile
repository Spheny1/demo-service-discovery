#FROM golang:1.22.1
#RUN mkdir /app 
#ADD . /app/ 
#WORKDIR /app 
#RUN go build -o /bin/service-discovery 
#FROM scratch
#COPY --from=0 /bin/service-discovery /bin/service-discovery
#EXPOSE 8080
#CMD ["/bin/service-discovery"]

############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
ENV CGO_ENABLED=1
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cache gcc musl-dev
WORKDIR . 
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build -ldflags='-s -w -extldflags "-static"' -o /go/bin/service-discovery
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/service-discovery /go/bin/service-discovery
EXPOSE 8080
# Run the hello binary.
ENTRYPOINT ["/go/bin/service-discovery"]
