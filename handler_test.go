package lab2

import (
	"fmt"
	"strings"
)

func ExampleComputeHandler() {

	good := &ComputeHandler{Input: strings.NewReader("2 2 8 - * 31 +")}
	_ = good.Compute()

	bad := &ComputeHandler{Input: strings.NewReader("some error")}
	err := bad.Compute()
	fmt.Println(err)

	//Output:
	//RESULT :  19
	//!!! You entered incorrect data !!!

}
