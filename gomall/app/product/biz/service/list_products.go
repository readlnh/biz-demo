package service

import (
	"context"

	"github.com/readlnh/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/readlnh/biz-demo/gomall/app/product/biz/model"
	product "github.com/readlnh/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	// fmt.Println("ListProductsService")

	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	// fmt.Println(c)
	if err != nil {
		return nil, err
	}

	resp = &product.ListProductsResp{}
	for _, v1 := range c {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       v.Price,
			})
		}
	}

	// fmt.Println(resp)

	return resp, nil
}
