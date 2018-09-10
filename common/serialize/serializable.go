package serialize

import (
	"io"
	"encoding/binary"
)

type Serializable interface{
	/**
	把数据写入到writer中去
	 */
	Serialize(w io.Writer) error
	/**
	把数据从reader中读出来
	 */
	Deserialize(r io.Reader) error
}

func WriteUint32(w io.Writer,num uint32) error{
	var p [4]byte
	binary.LittleEndian.PutUint32(p[:],num)
	_,err := w.Write(p[:])
	return err
}

func ReadUint32(r io.Reader) (uint32,error){
	var p [4]byte
	n,err := r.Read(p[:])
	if n <= 0 || err != nil {
		return 0,err
	}else{
		return binary.LittleEndian.Uint32(p[:]),nil
	}
}

func WriteString(w io.Writer,str string) error{
	err := WriteUint32(w,uint32(len(str)))
	if err != nil {
		return err
	}
	_,err = w.Write([]byte(str))
	return err
}

func ReadString(r io.Reader) (string,error){
	len,err := ReadUint32(r)
	if err != nil {
		return "",err
	}
	p := make([]byte,len)
	b,err := r.Read(p)
	if b<=0 || err != nil{
		return "",err
	}
	return string(p),nil
}


