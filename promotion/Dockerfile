FROM golang AS Dev
WORKDIR /go/src/goTemp
RUN go get github.com/githubnemo/CompileDaemon
ENV GO111MODULE=on
COPY ./promotion ./promotion
COPY ./globalerrors ./globalerrors
COPY ./go.mod ./go.sum ./
RUN go get -d  -v ./...
RUN go build -o promotionServer ./promotion/server/
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o promotionServerAlp ./promotion/server/
EXPOSE 50051
CMD ["./promotionServer"]

FROM Dev AS alpBuild
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o promotionServerAlp ./promotion/server/

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /goTemp
#COPY --from=Dev /go/src/goTemp/promotionServerAlp promotionServerAlp
COPY --from=alpBuild /go/src/goTemp/promotionServerAlp promotionServerAlp
CMD ["./promotionServerAlp"]