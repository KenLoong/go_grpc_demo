package services

import context "context"

type ProdService struct {
}

func (s *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (resp *ProResponse, err error) {
	count := 0
	if req.ProdArea == ProdAreas_A {
		count = 111
	} else if req.ProdArea == ProdAreas_B {
		count = 222
	} else if req.ProdArea == ProdAreas_C {
		count = 333
	}
	return &ProResponse{ProdStock: int32(count)}, nil
}

func (s *ProdService) GetProdStocks(ctx context.Context, req *QuerySize) (*ProStockList, error) {
	size := req.GetSize()
	res := make([]*ProResponse, 0)
	var i int32 = 0
	for i = 0; i < size; i++ {
		res = append(res, &ProResponse{ProdStock: i + 1})
	}
	return &ProStockList{Prodres: res}, nil
}

func (s *ProdService) GetProdInfo(context.Context, *ProdRequest) (*ProModel, error) {
	res := &ProModel{ProdName: "牛奶", ProdPrice: 22.3, ProdId: 5}
	return res, nil
}
