# Start all services and the web UI. Does not start server clients
start:
	docker-compose up -d usersrv customersrv productsrv promotionsrv auditsrv
	docker-compose up web
stop:
	docker-compose down

# Docker-compose sample commands
composeup:
	docker-compose up
composedown:
	docker-compose down
composebuildpromosrv:
	docker-compose build promotionsrv
composerestartpromocli:
	docker-compose rm -fsv promotioncli
	docker-compose up promotioncli

# Docker sample commands
docbuildpromosrv:
	docker build -t promosrv -f promotion/Dockerfile .
docrunpromosrv:
	docker run --env-file ./promotion/docker-compose.env -p 50051:50051 --name promosrvcont promosrv
# run and attach to existing network
docrunusersrvnet:
	docker run --env-file ./user/docker-compose.env --network=gotemp_default  -p 50051:50051 --name usersrvcont usersrv
docbuildpromocli:
	docker build -t promocli -f promotion/DockerfileCli .
docrunpromocli:
	docker run -p 50051:50051 --name promoclicont promocli

#DockerHub
hubpush:
	docker build -t $$SERVICE -f  $$FOLDER/Dockerfile .
	docker tag $$SERVICE bolbeck/gotemp_$$SERVICE
	docker push bolbeck/gotemp_$$SERVICE

hubpushcontext:
	docker build -t $$SERVICE -f  ./$$FOLDER/Dockerfile ./$$FOLDER
	docker tag $$SERVICE bolbeck/gotemp_$$SERVICE
	docker push bolbeck/gotemp_$$SERVICE

# Run service directly
runpromosrv:
	go run promotion/server/promotionServer.go
runpromocli:
	go run promotion/client/promotionClient.go


# Web App
# Directly (dev)
runweb:
	npm run dev

# Docker
docbuildweb:
	docker build -t gotempweb -f ./web/Dockerfile ./web
docrunweb:
	docker run -p 3000:3000 --name gotempwebcont gotempweb

#Docker-compose
composeupweb:
	docker-compose up web



# Compile proto files
genpromotionproto:
	protoc --proto_path=$$GOPATH/src:. --micro_out=source_relative:.. --go_out=. --go_opt=paths=source_relative promotion/proto/promotion.proto
genuserproto:
	protoc --proto_path=$$GOPATH/src:. --micro_out=source_relative:.. --go_out=. --go_opt=paths=source_relative user/proto/user.proto
gencustomerproto:
	protoc --proto_path=$$GOPATH/src:. --micro_out=source_relative:.. --go_out=. --go_opt=paths=source_relative customer/proto/customer.proto
genproductproto:
	protoc --proto_path=$$GOPATH/src:. --micro_out=source_relative:.. --go_out=. --go_opt=paths=source_relative product/proto/product.proto
genstandardFieldsproto:
	protoc --proto_path=$$GOPATH/src:. --micro_out=source_relative:.. --go_out=. --go_opt=paths=source_relative globalProtos/standardFields.proto

# Call service through the micro gateway
promoviaapigateway:
	curl --location --request POST 'http://localhost:8080/promotion/promotionSrv/getPromotions' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoyMzQzNzI1MzkxMjkxNjE4MzA1LCJjb21wYW55IjoiRHVjayBJbmMuIn0sImV4cCI6MTU5NzMzNTMzNywiaWF0IjoxNTk3MjQ4OTM3LCJpc3MiOiJnb1RlbXAudXNlcnNydiJ9.QWAvvoXQHv_Cf48PTrjK9uRvrdEblNvFOxQWjNcX79U' \
    --data-raw '{"name":"Promo1", "customerId": "ducksrus"}'

# Call service using the micro gateway running behind the ingress in K8s
authviaapigateway:
	curl --location --request POST 'http://gotemp.tst/user/userSrv/auth' \
	--header 'Content-Type: application/json' \
	--data-raw '{"pwd":"1234","email":"duck@mymail.com"}'

# K8s

startkub:
	kubectl apply -f cicd/K8s/dbsAndBroker
	kubectl apply -f cicd/K8s/services
	kubectl apply -f cicd/K8s/web
	kubectl apply -f cicd/K8s/ingress
stopkub:
	kubectl delete -f cicd/K8s/ingress
	kubectl delete -f cicd/K8s/web
	kubectl delete -f cicd/K8s/services
	kubectl delete -f cicd/K8s/dbsAndBroker


kapplyingress:
	kubectl apply -f cicd/K8s/ingress
kapplydbandborkers:
	kubectl apply -f cicd/K8s/dbsAndBroker
kapplyservices:
	kubectl apply -f cicd/K8s/services
kapplyclients:
	kubectl apply -f cicd/K8s/clients
kapplyweb:
	kubectl apply -f cicd/K8s/web
kdelete:
	kubectl delete -f $FOLDER

kstartSubset:
	kubectl apply $$(ls cicd/K8s/services/audit*.yaml | awk ' { print " -f " $$1 } ')


# Misc

encode:
	echo -n 'data' | base64
decode:
	echo -n ZGF0YQ== | base64 -d


# Vault integration on K8s

vkubsetup:
	kubectl exec vault-0 -- rm -rf /vault/file/scripts/
	kubectl exec vault-0 -- rm -rf /vault/file/policies/
	kubectl cp vault/policies vault-0:/vault/file/policies
	kubectl cp vault/scripts vault-0:/vault/file/scripts
	kubectl exec vault-0 -- /vault/file/scripts/allServices.sh  $$VAULT_TOKEN

vstartkub:
	make startkub
	make vkubpatchdeploy


vstopkub:
	make stopkub
	kubectl delete -f cicd/K8s/vault/serviceAccount

vkubteardown:
	kubectl exec vault-0 -- /vault/file/scripts/deleteAllSrv.sh $$VAULT_TOKEN
	kubectl exec vault-0 -- rm -rf /vault/file/scripts/
	kubectl exec vault-0 -- rm -rf /vault/file/policies/

vkubpatchdeploy:
	kubectl apply -f cicd/K8s/vault/serviceAccount
	kubectl patch deployment auditsrv --patch "$$(cat cicd/K8s/vault/patch/auditsrv-deployment-patch.yaml)"
	kubectl patch deployment customersrv --patch "$$(cat cicd/K8s/vault/patch/customersrv-deployment-patch.yaml)"
	kubectl patch deployment productsrv --patch "$$(cat cicd/K8s/vault/patch/productsrv-deployment-patch.yaml)"
	kubectl patch deployment promotionsrv --patch "$$(cat cicd/K8s/vault/patch/promotionsrv-deployment-patch.yaml)"
	kubectl patch deployment usersrv --patch "$$(cat cicd/K8s/vault/patch/usersrv-deployment-patch.yaml)"

# ---------------------------------------------
vkubunseal:
	kubectl exec -ti vault-0 -- vault operator unseal $$KEY
vkubui:
	kubectl port-forward vault-0 8100:8200
vkubenableseceng:
	export VAULT_TOKEN=<token>
	vault secrets enable -path=$$PATH $$TYPE

vkubtestrmsecret:
	kubectl apply -f cicd/K8s/vault/testYamlFile

