all:
	build

build:
	build_client
	build_server
	build_implant

build_client:
	go build -o client/clinet client/client.go

build_server:
	go build -o server/server server/server.go

build_implant:
	go build -o implant/implant implant/implant.go

clean:
	rm -rf client/client
	rm -rf server/server
	rm -rf implant/implant

.PHONY: clean