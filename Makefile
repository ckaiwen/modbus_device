APP=modbustest

protoc:
	protoc --go_out=plugins=grpc:.  proto/modbusdevice.proto

run:
	go run cmd/cmd.go
test:
	go run cmd/test.go
clean:
	rm -f ${APP}


