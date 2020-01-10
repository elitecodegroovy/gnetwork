package flyweight

import (
	"strconv"
	"testing"
	"time"
)

func TestFlyWeight(t *testing.T) {
	flyweight := NewFlyweight()
	flyweight.GetHeavyStone("1")
	flyweight.GetHeavyStone("2")
	flyweight.GetHeavyStone("3")
	flyweight.GetHeavyStone("11")
	flyweight.GetHeavyStone("22")
	flyweight.GetHeavyStone("32")

	t1 := time.Now()
	flyweight.GetHeavyStone("33")
	flyweight.GetHeavyStone("44")
	flyweight.GetHeavyStone("55")
	flyweight.GetHeavyStone("33")
	flyweight.GetHeavyStone("44")
	flyweight.GetHeavyStone("55")
	flyweight.GetHeavyStone("33")
	flyweight.GetHeavyStone("44")
	flyweight.GetHeavyStone("55")
	flyweight.GetHeavyStone("33")
	flyweight.GetHeavyStone("44")
	t.Log("time elapses " + strconv.FormatInt(time.Since(t1).Milliseconds(), 10) + "")
	t2 := time.Now()
	flyweight.GetHeavyStone("33")
	flyweight.GetHeavyStone("44")
	flyweight.GetHeavyStone("55")
	flyweight.GetHeavyStone("33")
	flyweight.GetHeavyStone("44")
	flyweight.GetHeavyStone("55")
	flyweight.GetHeavyStone("33")
	flyweight.GetHeavyStone("44")
	flyweight.GetHeavyStone("55")
	flyweight.GetHeavyStone("33")
	flyweight.GetHeavyStone("44")
	t.Log("time elapses " + strconv.FormatInt(time.Since(t2).Milliseconds(), 10) + "")
}
