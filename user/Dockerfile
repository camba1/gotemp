# Build DEv image with hot rebuild
FROM golang AS Dev

ENV SRVDIR=user
ENV SRVNAME=userServer

WORKDIR /go/src/goTemp
RUN go get github.com/githubnemo/CompileDaemon
ENV GO111MODULE=on
COPY ./$SRVDIR ./$SRVDIR
COPY ./globalerrors ./globalerrors
COPY ./globalUtils ./globalUtils
COPY ./go.mod ./go.sum ./
RUN go get -d  -v ./...
RUN go build -o $SRVNAME ./$SRVDIR/server/
EXPOSE 50051
#CMD ["./promotionServer"]
CMD ./$SRVNAME

# Compile the Alpine version of the application
FROM Dev AS alpBuild

ENV SRVDIR=user
ENV SRVNAME=userServerAlp

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $SRVNAME ./$SRVDIR/server/

# Build the Alpine image
FROM alpine

ENV SRVNAME=userServerAlp

RUN apk --no-cache add ca-certificates
WORKDIR /goTemp
COPY --from=alpBuild /go/src/goTemp/$SRVNAME $SRVNAME
#CMD ["./promotionServerAlp"]
CMD ./$SRVNAME