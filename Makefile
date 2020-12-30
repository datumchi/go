buildall: build-identity build-collaboration build-hborderer
buildalldocker: build-identity-docker build-collaboration-docker build-hborderer-docker
dockerpushall: buildalldocker identity-dockerpush collaboration-dockerpush hborderer-dockerpush



testall:
	go get -v -t -d ./...
	ginkgo -r




################### START IDENTITY ########################
build-identity: testall
	cd services/identity && \
	go build

build-identity-docker: build-identity
	rm -Rf services/identity/identity
	docker build -f docker/Dockerfile-identity -t datumchi/identity .
	docker tag datumchi/identity:latest 248181394449.dkr.ecr.us-west-2.amazonaws.com/datumchi/identity:latest

identity-dockerpush:
	docker push 248181394449.dkr.ecr.us-west-2.amazonaws.com/datumchi/identity:latest
################### END IDENTITY ########################


################### START COLLABORATION ########################
build-collaboration: testall
	cd services/collaboration && \
	go build

build-collaboration-docker: build-collaboration
	rm -Rf services/collaboration/collaboration
	docker build -f docker/Dockerfile-collaboration -t datumchi/collaboration .
	docker tag datumchi/collaboration:latest 248181394449.dkr.ecr.us-west-2.amazonaws.com/datumchi/collaboration:latest

collaboration-dockerpush:
	docker push 248181394449.dkr.ecr.us-west-2.amazonaws.com/datumchi/collaboration:latest
################### END COLLABORATION ########################


################### START HEARTBEAT ORDERER ########################
build-hborderer: testall
	cd services/hborderer && \
	go build

build-hborderer-docker: build-hborderer
	rm -Rf services/hborderer/hborderer
	docker build -f docker/Dockerfile-hborderer -t datumchi/hborderer .
	docker tag datumchi/hborderer:latest 248181394449.dkr.ecr.us-west-2.amazonaws.com/datumchi/hborderer:latest

hborderer-dockerpush:
	docker push 248181394449.dkr.ecr.us-west-2.amazonaws.com/datumchi/hborderer:latest
################### END HEARTBEAT ORDERER ########################



generate-protocol:
	mkdir -p ./generated/protocol
	protoc -I=../protocol/protobuf --go_out=. --go-grpc_out=. ../protocol/protobuf/*.proto