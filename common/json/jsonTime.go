package json

import (
	"time"
	"fmt"
)

type JsonTime struct {
	Data time.Time
}


func (this JsonTime) MarshalJSON() ([]byte, error) {
	value := this.Data.Unix()
	res := fmt.Sprintf("%d",value)
	return []byte(res), nil
}

func (this JsonTime) String()string{
	return fmt.Sprintf("\"%s\"", this.Data.Format("2006-01-02 15:04:05"))
}

