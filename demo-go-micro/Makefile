pbbuild:
	protoc --go_out=plugins=grpc:. proto/user.proto

micro:
	protoc --micro_out=. --go_out=. proto/user.proto