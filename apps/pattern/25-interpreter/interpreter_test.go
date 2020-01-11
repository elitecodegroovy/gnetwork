package interpreter

import (
	"strconv"
	"strings"
	"testing"
)

func TestInterpreter(t *testing.T) {
	stack := polishNotationStack{}
	operators := strings.Split("3 4 sum 2 sub", " ")

	for _, operatorString := range operators {
		if operatorString == SUM || operatorString == SUB {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := operatorFactory(operatorString, left, right)
			res := value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				panic(err)
			}

			temp := value(val)
			stack.Push(&temp)
		}
	}

	if int(stack.Pop().Read()) != int(5) {
		t.Fatal("failed")
	}
}
