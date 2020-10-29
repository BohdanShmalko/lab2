package lab2

import (
	"io"
	"fmt"
)

type ComputeHandler struct {
	Input io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	stringExpression := ReaderToExpression(ch.Input)
	res, err := CalculatePostfix(stringExpression)
	if err != nil {return err}

	if ch.Output == nil {
		fmt.Println("RESULT : ", res)
		return nil
	}

	stringResult := fmt.Sprintf("%f", res)
	ch.Output.Write([]byte(stringResult))
	return nil
}

func ReaderToExpression(r io.Reader) string {
	data := make([]byte, 64)
	stringExpression := ""
		for{
				n, err := r.Read(data)
				if err == io.EOF{ break }
				stringExpression += string(data[:n])
		}
		return stringExpression
}