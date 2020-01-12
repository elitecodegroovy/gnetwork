package future

import (
	"errors"
	"fmt"
	"testing"
)

func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%s Closure!\n", msg)
	return func() (string, error) {
		return msg, nil
	}
}

func TestFuture(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		task := new(MaybeString)
		task.Success(func(s string) {
			t.Fatal("success >>>" + s)
		}).Fail(func(e error) {
			t.Log("failed" + e.Error())
		}).Execute(func() (string, error) {
			msg := fmt.Sprintf("%s Closure failure!\n", "Future pattern ")
			return "", errors.New(" >>> can't get context" + msg)
		})
		task.wg.Wait()

	})

	t.Run("failed", func(t *testing.T) {
		task := new(MaybeString)
		task.Success(func(s string) {
			fmt.Println("success >>>" + s)
		}).Fail(func(e error) {
			t.Fatal("failed" + e.Error())
		}).Execute(setContext("Test future pattern with failure"))
		task.wg.Wait()
	})
}
