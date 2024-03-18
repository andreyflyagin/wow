
build: build-network \
	build-server \
	build-client
.PHONY: build

build-network:
	docker network create wow || true
.PHONY: build-network

build-server:
	docker build . -t wow-server -f Dockerfile-server
.PHONY: build-server

build-client:
	docker build . -t wow-client -f Dockerfile-client 3
.PHONY: build-client

server:
	docker run --rm -ti --network wow --name wow-server -p 8080:8080 wow-server quotes.txt 3
.PHONY: server

client:
	docker run --network wow --rm -it --entrypoint /client wow-client wow-server:8080 3
.PHONY: client


