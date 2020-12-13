package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "sa_system/modbus_device/proto"
)

func main() {
    conn,err:= grpc.Dial(":8000",grpc.WithInsecure())
    if err!= nil{
    	log.Fatal("dial err:v%",err)
	}
    defer conn.Close()

    client:=pb.NewDeviceClient(conn)

    //read1
    go func() {
    	for{
			r,err:=client.ReadHoldingRegisters(context.Background(), &pb.ReadHoldingRegistersRequest{SlaveId: 31,Address: 10,Num: 50})
			if err!= nil{
				fmt.Println(err)
			}
			fmt.Println("readFun1:", r)
			//time.Sleep(time.Millisecond*200)
		}
	}()

    //read2
	go func() {
		for  {
			r,err:=client.ReadHoldingRegisters(context.Background(), &pb.ReadHoldingRegistersRequest{SlaveId: 31,Address: 256,Num: 100})
			if err!= nil{
				fmt.Println(err)
			}
			fmt.Println("readFun2:", r)
			//time.Sleep(time.Millisecond*200)
		}
	}()

	//write 1
	go func() {
		for  {
			r,err:=client.WriteMultipleRegisters(context.Background(),&pb.WriteMultipleRegistersRequest{SlaveId: 31,Address: 1000,Num: 1,Value: []byte{0,1}})
			if err!= nil{
				fmt.Println("WriteFun1_1 err:",err)
			}
			fmt.Println("WriteFun1_1:", r)
			//time.Sleep(time.Millisecond*5)

			r,err=client.WriteMultipleRegisters(context.Background(),&pb.WriteMultipleRegistersRequest{SlaveId: 31,Address: 1000,Num: 1,Value: []byte{0,0}})
			if err!= nil{
				fmt.Println("WriteFun1_2 err:",err)
			}
			fmt.Println("WriteFun1_2:", r)
			//time.Sleep(time.Millisecond*5)
		}
	}()

	//write 2
	go func() {
		for  {
			r,err:=client.WriteMultipleRegisters(context.Background(),&pb.WriteMultipleRegistersRequest{SlaveId: 31,Address: 1001,Num: 1,Value: []byte{0,1}})
			if err!= nil{
				fmt.Println("WriteFun2_1 err:",err)
			}
			fmt.Println("WriteFun2_1:", r)
			//time.Sleep(time.Millisecond*5)

			r,err=client.WriteMultipleRegisters(context.Background(),&pb.WriteMultipleRegistersRequest{SlaveId: 31,Address: 1001,Num: 1,Value: []byte{0,0}})
			if err!= nil{
				fmt.Println("WriteFun2_2 err:",err)
			}
			fmt.Println("WriteFun2_2:", r)
			//time.Sleep(time.Millisecond*5)
		}
	}()

	select {

	}
}
