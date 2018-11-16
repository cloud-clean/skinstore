package lot

import (
	"fmt"
	"testing"
)

func TestLotUser_Login(t *testing.T) {
	account := "cloud"
	password := "cloud"
	user := LotUser{}
	if(user.Login(account,password)){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}

}