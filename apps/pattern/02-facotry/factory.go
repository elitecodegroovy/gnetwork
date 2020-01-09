package facotry

import (
	"errors"
	"fmt"
)

type Payment interface {
	pay(account float32) string
}

const (
	ZhiFuBao = 1
	WeChat   = 2
	Union    = 3
)

type ZhiFuBaoPayment struct{}
type WeChatPayment struct{}
type UnionPayment struct{}

func (z *ZhiFuBaoPayment) pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using ZhiFuBaoPayment\n", amount)
}

func (w *WeChatPayment) pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using WeChatPayment\n", amount)
}
func (u *UnionPayment) pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using UnionPayment\n", amount)
}

func GetPayment(paymentType int) (Payment, error) {
	switch paymentType {
	case ZhiFuBao:
		return new(ZhiFuBaoPayment), nil
	case WeChat:
		return new(WeChatPayment), nil
	case Union:
		return new(UnionPayment), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment type %d can not be recognized.", paymentType))
	}
}
