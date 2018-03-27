`
CreateContainerPostgreSQL

docker run --name postgres -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres:latest

cd proto

protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. user.proto`
