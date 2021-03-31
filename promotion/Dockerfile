# Build DEv image with hot rebuild
FROM golang AS dev

ENV SRVDIR=promotion

ENV SRVNAME=${SRVDIR}Server
WORKDIR /go/src/goTemp
RUN go get github.com/githubnemo/CompileDaemon
ENV GO111MODULE=on
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./$SRVDIR ./$SRVDIR
COPY ./user/proto ./user/proto
COPY ./customer/proto ./customer/proto
COPY ./globalProtos ./globalProtos
COPY ./globalerrors ./globalerrors
COPY ./globalUtils ./globalUtils
COPY ./globalMonitoring ./globalMonitoring
#COPY ./go.mod ./go.sum ./
#RUN go get -d  -v ./...
RUN go build -o $SRVNAME ./$SRVDIR/server/
EXPOSE 50051
EXPOSE 2112
CMD ./$SRVNAME

# Compile the Alpine version of the application
FROM dev AS alpBuild

ENV SRVDIR=promotion

ENV SRVNAME=${SRVDIR}ServerAlp
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $SRVNAME ./$SRVDIR/server/

# Build the Alpine image
FROM alpine

ENV SRVDIR=promotion

ENV SRVNAME=${SRVDIR}ServerAlp
RUN apk --no-cache add ca-certificates
WORKDIR /goTemp
COPY --from=alpBuild /go/src/goTemp/$SRVNAME $SRVNAME
#CMD ["./promotionServerAlp"]
CMD ./$SRVNAME