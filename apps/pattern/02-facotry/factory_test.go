package facotry

import (
	"strings"
	"testing"
)

func TestGetPayment(t *testing.T) {
	var amount float32 = 9.0
	payment, err := GetPayment(1)
	if err != nil {
		t.Fatal("Can't get Payment instance.")
	}
	paymentResult := payment.pay(amount)
	if !strings.Contains(paymentResult, "ZhiFuBao") {
		t.Fatal("Payment is not correct.")
	}

	payment, err = GetPayment(2)
	if err != nil {
		t.Fatal("Can't get Payment instance.")
	}
	paymentResult = payment.pay(amount)
	if !strings.Contains(paymentResult, "WeChat") {
		t.Fatal("Payment is not correct.")
	}

	payment, err = GetPayment(3)
	if err != nil {
		t.Fatal("Can't get Payment instance.")
	}
	paymentResult = payment.pay(amount)
	if !strings.Contains(paymentResult, "UnionPayment") {
		t.Fatal("Payment is not correct.")
	}

	payment, err = GetPayment(4)
	if err == nil {
		t.Fatal("The Default payment doesn't work.")
	}
}
