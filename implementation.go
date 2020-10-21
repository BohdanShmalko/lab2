package lab2

import (
  "strings" //strings functions
  "strconv" //string to number
  "math" //for power
  "errors" //for error
)

func CalculatePostfix(input string) (float64, error) {

  if input == "" { return -1, errors.New("!!! Enter something !!!") }

  operators := map[string]func(float64, float64)(float64){
    "+" : func (x, y float64) float64{return x + y},
    "-" : func (x, y float64) float64{return x - y},
    "*" : func (x, y float64) float64{return x * y},
    "/" : func (x, y float64) float64{return x / y},
    "^" : func (x, y float64) float64{return math.Pow(x, y)},
  }

  var stack []float64

  strArr := strings.Fields(input)

  for _, val := range strArr {
    if action , found := operators[val]; found {
      y := stack[len(stack) - 1]
      stack = stack[:len(stack)-1]
      x := stack[len(stack) - 1]
      stack = stack[:len(stack)-1]
      stack = append(stack, action(x,y))
    } else {
      number, err := strconv.ParseFloat(val, 64)
      if err != nil {  return -1, errors.New("!!! You entered incorrect data !!!") }
      stack = append(stack, number)
    }
  }

  if len(stack) != 1 { return -1, errors.New("!!! You entered incorrect data !!!") }
	return stack[len(stack) - 1], nil
}
