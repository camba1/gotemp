FROM golang AS Dev
ARG SRVDIR=user
ARG SRVNAME=userServer
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

FROM Dev AS alpBuild
ARG SRVDIR=user
ARG SRVNAME=userServerAlp
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $SRVNAME ./$SRVDIR/server/

FROM alpine
ARG SRVNAME=userServerAlp
RUN apk --no-cache add ca-certificates
WORKDIR /goTemp
COPY --from=alpBuild /go/src/goTemp/$SRVNAME $SRVNAME
#CMD ["./promotionServerAlp"]
CMD ./$SRVNAME