package handler

import (
	"context"

	inv "github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/inventory-srv/model/inventory"
	proto "github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/inventory-srv/proto/inventory"
	"github.com/micro/go-micro/util/log"
)

var (
	invService inv.Service
)

type Service struct {
}

// Init 初始化handler
func Init() {
	invService, _ = inv.GetService()
}

// Sell 库存销存
func (e *Service) Sell(ctx context.Context, req *proto.InventoryRequest, rsp *proto.InventoryResponse) (err error) {
	id, err := invService.Sell(req.BookId, req.UserId)
	if err != nil {
		log.Logf("[Sell] 销存失败，bookId：%d，userId: %d，%s", req.BookId, req.UserId, err)
		rsp.Success = false
		return
	}

	rsp.InvH = &proto.InvHistory{
		Id: id,
	}

	rsp.Success = true
	return nil
}

// Confirm 库存销存 确认
func (e *Service) Confirm(ctx context.Context, req *proto.InventoryRequest, rsp *proto.InventoryResponse) (err error) {
	err = invService.Confirm(req.HistoryId, int(req.HistoryState))
	if err != nil {
		log.Logf("[Confirm] 确认销存失败，%s", err)
		rsp.Success = false
		return
	}

	rsp.Success = true
	return nil
}
