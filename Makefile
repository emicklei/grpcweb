gen:
	protoc --go_out=. --go-grpc_out=. test/test.proto