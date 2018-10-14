package tcp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"skinstore/common/protocol"
	"time"
)




func Start(){
	fmt.Println("start tcp server...")
	lister , err := net.Listen("tcp",":1986")
	if err != nil{
		fmt.Println(err)
	}
	for{
		if conn,err := lister.Accept();err == nil{
			go handleConn(conn)
		}
	}
}


func handleConn(conn net.Conn){
	var flag = make(chan byte)
	tmpBuf := make([]byte,0)

	buf := make([]byte,1024)
	defer conn.Close()
	for{
		n,err := conn.Read(buf)
		go heartBeating(conn,flag,6)
		if err != nil{
			fmt.Println(err)
			return
		}
		if n > 0{
			tmpBuf = append(tmpBuf,buf[:n]...)
			reader := bytes.NewBuffer(tmpBuf)
			scanner := bufio.NewScanner(reader)
			scanner.Split(func(data []byte,atEOF bool)(advance int, token []byte,err error){
				var m uint32
				binary.Read(bytes.NewReader(data[:4]),binary.BigEndian,&m)
				if !atEOF && m == protocol.MAGIC{
					if len(data) > 10{
						length := int32(0)
						err:=binary.Read(bytes.NewReader(data[6:10]),binary.BigEndian,&length)
						if err != nil {
							fmt.Println(err)
						}
						if int(length) + 10 <= len(data){
							return int(length+10),data[:length+10],nil
						}
					}
				}
				return
			})
			for scanner.Scan(){
				msg := new(protocol.Message)
				buf := scanner.Bytes()
				tmpBuf = tmpBuf[len(buf):]
				msg.Unpack(bytes.NewReader(buf))
				go handMsg(msg,conn,flag)
			}
		}
	}
}

func handMsg(msg *protocol.Message,conn net.Conn,flag chan byte){
	switch(msg.Type){
	case 0:
		fmt.Println("ping")
		resp := protocol.MakeMsg(1,nil)
		flag <- 1
		fmt.Println(len(flag))
		resp.Pack(conn)
		break
	case 1:
		fmt.Println("pong")
		break
	case 2:
		fmt.Println("text")
		fmt.Println(string(msg.Data))
		resp := protocol.MakeMsg(2,[]byte("recevie your msg"))
		resp.Pack(conn)
		break
	default:
		break
	}
}

func heartBeating(conn net.Conn,channel chan byte,timeout int){
	select{
	case <- channel:
		conn.SetDeadline(time.Now().Add(time.Duration(timeout)*time.Second))
		break;
	case <- time.After(time.Second*6):
		conn.Close()
	}
}