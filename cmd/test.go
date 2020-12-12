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
	handler.SlaveId = 1
	handler.Timeout = 10 * time.Second
    err:= handler.Connect()
    if err!= nil{
   		panic(err)
    }
    defer handler.Close()
    client:= modbus.NewClient(handler)

    for{
		handler.SlaveId = 31
		results,err:=client.ReadHoldingRegisters(0x100,2)
		if err!= nil{
			panic(err)
		}
		fmt.Println(results)
        time.Sleep(time.Second)

		handler.SlaveId = 1
		results,err=client.ReadHoldingRegisters(0x0,2)
		if err!= nil{
			panic(err)
		}
		fmt.Println(results)
		time.Sleep(time.Millisecond*200)
	}


}
