#FROM golang AS Dev
#WORKDIR /go/src/goTemp
#RUN go get github.com/githubnemo/CompileDaemon
#ENV GO111MODULE=on
#COPY ./promotion ./promotion
#COPY ./globalerrors ./globalerrors
#COPY ./globalUtils ./globalUtils
#COPY ./go.mod ./go.sum ./
#RUN go get -d  -v ./...
#RUN go build -o promotionServer ./promotion/server/
##RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o promotionServerAlp ./promotion/server/
#EXPOSE 50051
#CMD ["./promotionServer"]
#
#FROM Dev AS alpBuild
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o promotionServerAlp ./promotion/server/
#
#FROM alpine
#RUN apk --no-cache add ca-certificates
#WORKDIR /goTemp
##COPY --from=Dev /go/src/goTemp/promotionServerAlp promotionServerAlp
#COPY --from=alpBuild /go/src/goTemp/promotionServerAlp promotionServerAlp
#CMD ["./promotionServerAlp"]


# Build DEv image with hot rebuild
FROM golang AS Dev

ENV SRVDIR=promotion
ENV SRVNAME=promotionServer

WORKDIR /go/src/goTemp
RUN go get github.com/githubnemo/CompileDaemon
ENV GO111MODULE=on
COPY ./$SRVDIR ./$SRVDIR
COPY ./user/proto ./user/proto
COPY ./globalerrors ./globalerrors
COPY ./globalUtils ./globalUtils
COPY ./go.mod ./go.sum ./
RUN go get -d  -v ./...
RUN go build -o $SRVNAME ./$SRVDIR/server/
EXPOSE 50051
CMD ./$SRVNAME

# Compile the Alpine version of the application
FROM Dev AS alpBuild

ENV SRVDIR=promotion
ENV SRVNAME=promotionServerAlp

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $SRVNAME ./$SRVDIR/server/

# Build the Alpine image
FROM alpine

ENV SRVNAME=userServerAlp

RUN apk --no-cache add ca-certificates
WORKDIR /goTemp
COPY --from=alpBuild /go/src/goTemp/$SRVNAME $SRVNAME
#CMD ["./promotionServerAlp"]
CMD ./$SRVNAME