# gRPC random

### Generating proto files

```bash
protoc --proto_path=./proto ./proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=./pb 

protoc --proto_path=./proto ./proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go --go_out=./pb
```


[YT vid](https://www.youtube.com/watch?v=a6G5-LUlFO4)