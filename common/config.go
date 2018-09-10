package common

import (
	"os"
	"bufio"
	"io"
	"strings"
)

type Config struct{
	data map[string]string
	strcet string
}
var config *Config

func NewConfig() *Config{
	return config
}

func init(){
	config = &Config{data:make(map[string]string)}
	path := "./config.ini"
	f,err := os.Open(path)
	CheckErr(err)
	defer f.Close()
	r := bufio.NewReader(f)

	for{
		b,_,err := r.ReadLine()
		if err != nil{
			if err == io.EOF{
				break
			}
			CheckErr(err)
		}
		line := strings.TrimSpace(string(b))
		if strings.Index(line,"#") == 0{
			continue
		}
		n1 := strings.Index(line,"[")
		n2 := strings.Index(line,"]")
		if n1 >-1 && n2 > -1 && n2 > n1+1{
			config.strcet = strings.TrimSpace(line[n1+1:n2])
			continue
		}
		if len(config.strcet) == 0{
			continue
		}
		index := strings.Index(line,"=")
		if index< 0{
			continue
		}
		key := strings.TrimSpace(line[:index])
		if len(key)==0{
			continue
		}
		value := strings.TrimSpace(line[index+1:])
		if len(value) == 0{
			continue
		}
		pos := strings.Index(value,"\t#")
		if pos > -1{
			value = value[:pos]
		}
		pos = strings.Index(value," #")
		if pos > -1 {
			value = value[:pos]
		}
		pos = strings.Index(value,"\t//")
		if pos > -1{
			value = value[:pos]
		}
		pos = strings.Index(value," //")
		if pos > -1 {
			value = value[:pos]
		}
		config.data[config.strcet+"==="+key] = value
	}
}

func (c *Config)Get(node,key string)string{
	v,ok := c.data[node+"==="+key]
	if ok{
		return v
	}
	return ""
}
