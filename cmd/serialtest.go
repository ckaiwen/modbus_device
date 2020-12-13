package main

import (
	"fmt"
	modbus "github.com/goburrow/modbus"
	"time"
)

func main() {
    handler:= modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 31
	handler.Timeout = 2 * time.Second
    err:= handler.Connect()
    if err!= nil{
   		panic(err)
    }
    defer handler.Close()
    client:= modbus.NewClient(handler)

    //read test
    //for{
	//	handler.SlaveId = 31
	//	results,err:=client.ReadHoldingRegisters(10,1)
	//	if err!= nil{
	//		panic(err)
	//	}
	//	fmt.Println(results)
    //   time.Sleep(time.Second)
	//
	//	//handler.SlaveId = 31
	//	//results,err=client.ReadHoldingRegisters(0x100,2)
	//	//if err!= nil{
	//	//	panic(err)
	//	//}
	//	//fmt.Println(results)
	//	//time.Sleep(time.Millisecond*200)
	//}

	//write test
	for{
		results,err:=client.WriteMultipleRegisters(1001,1,[]byte{0,1})
		if err!= nil{
			fmt.Println(err)
		}
		fmt.Println("1:",results)
		time.Sleep(time.Second)

		results,err=client.WriteMultipleRegisters(1001,1,[]byte{0,0})
		if err!= nil{
			fmt.Println(err)
		}
		fmt.Println("2:",results)
		time.Sleep(time.Second)
	}

}
