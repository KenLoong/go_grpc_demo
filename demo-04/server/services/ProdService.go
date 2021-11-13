package services

import context "context"

type ProdService struct {
}

func (s *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (resp *ProResponse, err error) {
	return &ProResponse{ProdStock: 555}, nil
}
