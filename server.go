package modbus_device

import (
	"context"
	modbus "github.com/goburrow/modbus"
	"log"
	pb "sa_system/modbus_device/proto"
	"sync"
	"time"
)

const DevMinComTime= 10 //硬件最小操作间隔时间,如果ｍodbus一返回数据立马发送,可能存在丢数据的情况

type Server struct {
	devLock    sync.Mutex              //485硬件锁,只有超时或者返回后才能进行下次读写
	devHandler *modbus.RTUClientHandler //modbus485硬件句柄
    devClient  modbus.Client           //modbus485硬件读写客户端
}

func (s *Server) ModbusDevInit(handler *modbus.RTUClientHandler ,client modbus.Client )  {
	s.devHandler =handler
	s.devClient=client
}

func (s *Server) ReadHoldingRegisters(ctx context.Context, in *pb.ReadHoldingRegistersRequest) (*pb.ReadHoldingRegistersResponse, error){
	var res pb.ReadHoldingRegistersResponse

	s.devLock.Lock()
    s.devHandler.SlaveId= byte(in.SlaveId) //设置站号
	results,err:= s.devClient.ReadHoldingRegisters(uint16(in.Address),uint16(in.Num))
	go func() {
		time.Sleep(time.Millisecond* DevMinComTime)
		s.devLock.Unlock()
	}()

	if err!=nil{
		log.Println("ReadHoldingRegisters err:",err)
		res.ErrCode= pb.ErrorCode_TIMEOUT
	}else{
		res.ErrCode=pb.ErrorCode_NORMAL
		res.Results=results
	}
	return &res,nil
}

func (s *Server) WriteMultipleRegisters(ctx context.Context, in *pb.WriteMultipleRegistersRequest) (*pb.WriteMultipleRegistersResponse, error){
	var res pb.WriteMultipleRegistersResponse

	s.devLock.Lock()
	s.devHandler.SlaveId= byte(in.SlaveId) //设置站号
	_,err:= s.devClient.WriteMultipleRegisters(uint16(in.Address),uint16(in.Num),in.Value)

	go func() { //数据返回后需要间隔一段时间才能重新发送
		time.Sleep(time.Millisecond* DevMinComTime)
		s.devLock.Unlock()
	}()

	if err!=nil{
		log.Println("WriteMultipleRegisters err:",err)
		res.ErrCode= pb.ErrorCode_TIMEOUT
	}else{
		res.ErrCode=pb.ErrorCode_NORMAL
	}
	return &res,nil
}
