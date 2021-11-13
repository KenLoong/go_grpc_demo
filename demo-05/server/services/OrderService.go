package services

import (
	context "context"
	"fmt"
)

type OrderService struct {
}

func (o *OrderService) NewOrder(ctx context.Context, req *OrderRequest) (*OrderResponse, error) {
	err := req.OrderMain.Validate() //验证是否符合要求
	if err != nil {
		return &OrderResponse{
			Status:  "error",
			Message: err.Error(),
		}, nil
	}
	fmt.Println(req.OrderMain)
	return &OrderResponse{
		Status:  "ok",
		Message: "success",
	}, nil
}
