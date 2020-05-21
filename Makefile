composeup:
	docker-compose up
composedown:
	docker-compose down
composebuildpromosrv:
	docker-compose build promotionsrv
composerestartpromocli:
	docker-compose rm -fsv promotioncli
	docker-compose up promotioncli
docbuildpromosrv:
	docker build -t promosrv -f /promotion/Dockerfile .

docrunpromosrv:
	docker run -p 50051:50051 --name promosrvcont promosrv
docbuildpromocli:
	docker build -t promocli -f /promotion/DockerfileCli .
docrunpromocli:
	docker run -p 50051:50051 --name promoclicont promocli

runpromosrv:
	go run promotion/server/promotionServer.go
runpromocli:
	go run promotion/client/promotionClient.go

genpromotionproto:
	protoc --proto_path=$$GOPATH/src:. --micro_out=source_relative:.. --go_out=. --go_opt=paths=source_relative promotion/promotion.proto
