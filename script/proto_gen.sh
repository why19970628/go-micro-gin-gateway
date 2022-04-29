cd ../models/protos
protoc --micro_out=../ --go_out=../ ./Users.proto
protoc --micro_out=../ --go_out=../ ./UserService.proto