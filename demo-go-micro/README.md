
### CreateContainerPostgreSQL

    docker run --name postgres -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres:latest


### Gen file proto

    cd proto
    

    protoc --proto\_path=$GOPATH/src:. --micro\_out=. --go_out=. user.proto
