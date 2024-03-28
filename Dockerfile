FROM golang:1.22.1
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o /bin/service-discovery .
#FROM scratch
#COPY --from=0 /bin/presentation-layer /bin/presentation-layer
EXPOSE 8080
CMD ["/bin/service-discovery"]
