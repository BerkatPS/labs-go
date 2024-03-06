generate_grpc:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	--plugin=/Users/berkatsaragih/go/bin/protoc-gen-go-grpc \
	client.proto

remove :
	rm message/*.go
