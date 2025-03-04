# Generate proto files
set GO_PATH=C:\path\to\go
set PATH=%PATH%;C:\path\to\go\bin

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative protofiles\data_streaming\streamingData.proto
