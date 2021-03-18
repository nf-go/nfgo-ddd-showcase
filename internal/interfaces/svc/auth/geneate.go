//go:generate protoc -I . -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate --go_out=. --go-grpc_out=. --validate_out=lang=go:. auth.proto

package auth
