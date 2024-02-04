# Generate a compiled go version of log.proto
compile:
		d:\Tools\protoc-23.4-win64\bin\protoc .\api\v1\log.proto --go_out=. --go_opt=paths=source_relative --proto_path=.

test:
	go test -race .\...
