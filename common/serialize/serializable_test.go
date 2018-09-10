package serialize

import (
	"testing"
	"bytes"
	"fmt"
)

func TestReadString(t *testing.T) {
	str := "seijas faijs * weir234aga"
	w := new(bytes.Buffer)
	err := WriteString(w,str)
	if err != nil{
		fmt.Print(err)
	}
	fmt.Println(ReadString(w))


}
