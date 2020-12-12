APP=test

protoc:

	protoc --go_out=plugins=grpc:.  /protoc/modbusse

clean:
	rm -f ${APP}
