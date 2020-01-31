package subscriber

import (
	"context"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/orders-srv/model/order"
	payS "github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/payment-srv/proto/payment"
	"github.com/micro/go-micro/util/log"
)

var (
	ordersService order.Service
)

// Init 初始化handler
func Init() {
	ordersService, _ = order.GetService()
}

// PayOrder 订单支付消息
func PayOrder(ctx context.Context, event *payS.PayEvent) (err error) {
	log.Logf("[PayOrder] 收到支付订单通知，%d，%d", event.OrderId, event.State)

	err = ordersService.UpdateOrderState(event.OrderId, int(event.State))
	if err != nil {
		log.Logf("[PayOrder] 收到支付订单通知，更新状态异常，%s", err)
		return
	}
	return
}
