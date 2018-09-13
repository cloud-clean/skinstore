package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

var VERSION uint8 = 0
var MAGIC uint32 = 0xeea0eea0



type Message struct {
	Magic uint32
	Ver uint8
	Type uint8
	Len uint32
	Data []byte
}



func MakeMsg(t uint8,data []byte) Message{
	return Message{Magic:MAGIC,Ver:VERSION,Type:t,Len:uint32(len(data)),Data: data}
}


func (msg *Message) Pack(writer io.Writer) error{
	var err error
	err = binary.Write( writer,binary.BigEndian,&msg.Magic)
	err = binary.Write( writer,binary.BigEndian,&msg.Ver)
	err = binary.Write( writer,binary.BigEndian,&msg.Type)
	err = binary.Write( writer,binary.BigEndian,&msg.Len)
	err = binary.Write( writer,binary.BigEndian,&msg.Data)

	return err
}


func (msg *Message) Unpack(reader io.Reader) error{
	var err error
	err = binary.Read(reader,binary.BigEndian,&msg.Magic)
	err = binary.Read(reader,binary.BigEndian,&msg.Ver)
	err = binary.Read(reader,binary.BigEndian,&msg.Type)
	err = binary.Read(reader,binary.BigEndian,&msg.Len)
	msg.Data = make([]byte,msg.Len)
	err = binary.Read(reader,binary.BigEndian,&msg.Data)
	return err
}

func CheckMagic(reader io.Reader) bool{
	var m uint32
	err := binary.Read(reader,binary.BigEndian,&m)
	if err != nil{
		fmt.Println(err)
		return false
	}
	if(m == MAGIC){
		return true
	}else{
		return false
	}
}
