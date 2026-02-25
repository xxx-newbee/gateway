package adapter

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/types"
	"github.com/xxx-newbee/order/order"
	"github.com/xxx-newbee/order/order_client"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type OrderClientAdapter struct {
	cli order_client.Order
}

func NewOrderClientAdapter(rpcClient zrpc.Client) *OrderClientAdapter {
	return &OrderClientAdapter{
		cli: order_client.NewOrder(rpcClient),
	}
}

func (e *OrderClientAdapter) SeckillStock(ctx context.Context, req *types.SeckillRequest, opts ...grpc.CallOption) (*types.SeckillResponse, error) {
	protoReq := &order.SeckillOrderRequest{
		UserId:     req.UserId,
		ActivityId: req.ActivityId,
		ProductId:  req.ProductId,
	}
	resp, err := e.cli.SeckillOrder(ctx, protoReq)
	if err != nil {
		return nil, err
	}
	return &types.SeckillResponse{
		OrderNo: resp.OrderNo,
		Success: resp.Success,
		Msg:     resp.Msg,
	}, nil
}
