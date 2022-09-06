protoc315 --proto_path=protos/user-proto --go_out=./ models.proto
protoc315 --proto_path=protos/user-proto --go_out=./ userkind.proto
protoc315 --proto_path=protos/user-proto --go-grpc_out=./ usersvc.proto