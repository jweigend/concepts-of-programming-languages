# 
# Makefile. Usage: make <target>.
# 


help:
	clear
	@echo "---------------------------------------------------------------------------------------------------------"
	@echo "Usage: make <target>"
	@echo "---------------------------------------------------------------------------------------------------------"
	@echo "Valid targets are:"
	@echo " - ginstall : Installs / Updates the GRPC Go Libraries "
	@echo " - gcompile : Compiles the .pb protocol files"
	@echo " - run-server : Compiles and starts the server"
	@echo " - run-client : Compiles and runs the client"
	@echo " - clean : Clean up generated files"
	@echo "---------------------------------------------------------------------------------------------------------"

ginstall:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go
	echo Make sure that protoc is installed for your platform

gcompile:
	protoc -I remote/idserv/ remote/idserv/idserv.proto --go_out=plugins=grpc:remote/idserv 

clean:
	rm remote/idserv/*.pb.go

run-server:
	go run remote/server/main.go

run-client:
	go run client/main.go

	