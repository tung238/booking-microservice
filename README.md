# booking-microservice

generate service:
protoc --plugin=%GOPATH%/bin/protoc-gen-go.exe --plugin=%GOPATH%/bin/protoc-gen-micro.exe -I=. --micro_out=. --go_out=. greeter.proto