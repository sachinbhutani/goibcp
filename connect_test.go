package goibcp

import (
	"fmt"
	"testing"
	"time"
)

func Test_AutoTickle(t *testing.T) {
	ib, err := Connect(&Config{CPURL: "http://localhost:5000", LogLevel: 3, AutoTickle: true})
	if err != nil {
		t.Error("Not able to connect")
	} else {
		fmt.Printf("IB Client: %+v \n", ib)
	}
	time.Sleep(3 * time.Minute)
	ib.Logout()
}
func Test_Connection_Settings(t *testing.T) {
	ib, err := Connect(&Config{CPURL: "http://localhost:5000", LogLevel: 3, AutoTickle: false})
	if err != nil {
		t.Error("Not able to connect")
	} else {
		fmt.Printf("IB Client: %+v \n", ib)
	}
	for i := 0; i < 3; i++ {
		time.Sleep(60 * time.Second)
		err = ib.Tickle()
	}
	fmt.Printf("End connection test")
}
