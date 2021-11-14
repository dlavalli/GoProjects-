# Generate a compileg go version of log.proto
compile:
		protoc .\api\v1\log.proto --go_out=. --go_opt=paths=source_relative --proto_path=.

test:
	go test -race .\...
