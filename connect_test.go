package goibcp

import (
	"fmt"
	"testing"
	"time"
)

var URL = "https://ib.froogle.in"

func Test_AutoTickle(t *testing.T) {
	ib, err := Connect(&Config{CPURL: URL, LogLevel: 3, AutoTickle: true})
	if err != nil {
		t.Error("Not able to connect")
	} else {
		fmt.Printf("IB Client: %+v \n", ib)
	}
	time.Sleep(30 * time.Minute)
	ib.Logout()
}
func Test_Connection_Settings(t *testing.T) {
	ib, err := Connect(&Config{CPURL: URL, LogLevel: 2, AutoTickle: false})
	if err != nil {
		t.Error("Not able to connect")
	} else {
		fmt.Printf("IB Client: %+v \n", ib)
	}
	var user IBUser
	ib.GetSessionInfo(&user)
	fmt.Println(&user)
	for i := 0; i < 2; i++ {
		time.Sleep(60 * time.Second)
		err = ib.Tickle()
		fmt.Printf("IB Client: %+v", ib)
	}

	fmt.Printf("End connection test")
}
