package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//单例模式测试代码
func TestCounter(t *testing.T) {
	counter1 := GetInstance()

	assert.EqualValues(t, 1, counter1.AddOne())
	println("count 1 : ", counter1.count)

	counter2 := GetInstance()

	assert.EqualValues(t, 2, counter2.AddOne())
}
