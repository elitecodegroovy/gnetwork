package handler

import (
	"context"

	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/orders-srv/model/order"
	proto "github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/orders-srv/proto/order"
	"github.com/go-log/log"
)

var (
	ordersService order.Service
)

type Orders struct {
}

// Init 初始化handler
func Init() {
	ordersService, _ = order.GetService()
}

// New 新增订单
func (e *Orders) New(ctx context.Context, req *proto.OrderRequest, rsp *proto.OrderResponse) (err error) {
	orderId, err := ordersService.New(req.BookId, req.UserId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Order = &proto.Order{
		Id: orderId,
	}
	return
}

// GetOrder 获取订单
func (e *Orders) GetOrder(ctx context.Context, req *proto.OrderRequest, rsp *proto.OrderResponse) (err error) {
	log.Logf("[GetOrder] 收到获取订单请求，%d", req.OrderId)

	rsp.Order, err = ordersService.GetOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Success = true
	return
}
