package goibcp

import (
	"fmt"
	"testing"
)

func Test_Connection_Settings(t *testing.T) {
	ib, err := Connect(&Config{CPURL: "http://localhost:5000", LogLevel: 2})
	if err != nil {
		t.Error("Not able to connect")
	} else {
		fmt.Printf("IB Client: %+v \n", ib)
	}
}

// func Test_Logout(t *testing.T) {
// 	ib, err := Connect(&Config{CPURL: "http://localhost:5000", LogLevel: 2})
// 	if err != nil {
// 		t.Error("Not able to connect")
// 	} else {
// 		ib.Logout()
// 	}
// }
