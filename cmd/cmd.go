package main

import (
	"github.com/goburrow/modbus"
	"google.golang.org/grpc"
	"log"
	"net"
	"sa_system/modbus_device"
	pb "sa_system/modbus_device/proto"
	"time"
)

func main() {
	//tcp listen
	listen,err:=net.Listen("tcp",":8000")
	if err!= nil{
		log.Fatal("listen err:v%",err)
	}
   //dev init
	//handler,client,err:=DevInit()
	//if err!=nil{
	//	log.Fatal("Dev Init Err:v%",err)
	//}
	handler:= modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 2 * time.Second
	handler.Connect()
	defer handler.Close()
	client:= modbus.NewClient(handler)

	//server init
	srv:=modbus_device.Server{}
	srv.ModbusDevInit(handler,client)

	//grpc register
	grpcServer:= grpc.NewServer()
	pb.RegisterDeviceServer(grpcServer,&srv)
   //server
	err=grpcServer.Serve(listen)
	if err!=nil{
		log.Fatal("grpc server err:v%",err)
	}
}

func DevInit() (*modbus.RTUClientHandler,*modbus.Client,error) {
	handler:= modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 5 * time.Second
	err:= handler.Connect()
	if err!= nil{
		return nil, nil, err
	}
	defer handler.Close()
	client:= modbus.NewClient(handler)

	return handler,&client,nil
}