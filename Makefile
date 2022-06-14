compile:
	protoc api/v1/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.
test:
	go test -race ./...

compile2:
	protoc api/v1/*.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
