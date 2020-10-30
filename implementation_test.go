package lab2

import (
    "testing"
    gocheck "gopkg.in/check.v1"
		"fmt"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type MySuite struct{}
var _ = gocheck.Suite(&MySuite{})

func (s *MySuite) TestCalculatePostfix(c *gocheck.C) {

var testMap = map[string]float64 {
	"2 2 8 - * 31 +" : 19, // 2 * (2 - 8) + 31
	"146 275 + 438 + 200 -" : 659, //146 + 275 + 438 - 200
	"435 5 / 411 3 * + 13 -" : 1307 , //435 / 5 + 411 * 3 - 13
	"6 3 ^ 348 2 / + 610 3 * - 210 3 * + 3 6 ^ + 12 -" : -93, //6 ^ 3 + 348 / 2 - 610 * 3 + 210 * 3 + 3 ^ 6 - 12
	"614 22300 2 / + 4 6 ^ + 25 5 ^ 12 / + 93 +" : 829755.0833333334, //614 + 22300 / 2 + 4 ^ 6 + 25 ^ 5 / 12 + 93
	"33448 2212 16 / 12 2 ^ * + 16 4 / 28 * + 483 -" : 52985, //33448 + 2212 / 16 * 12 ^ 2 + 16 / 4 * 28 - 483
  "+ + 2 22 31" : -1, //incorrect data
	"ksksk sllsl " : -1, //incorrect data
  "" : -1, // empty string
}

var ( res float64; err error )
for key , val := range testMap {
	res, err = CalculatePostfix(key)
	if err != nil {
      c.Error(err)
			//fmt.Println("ERROR : ", err)
		} else {
			c.Check(res, gocheck.Equals, val)
		}
  }
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("2 2 +")
	fmt.Println(res)

	// Output: 4
}
