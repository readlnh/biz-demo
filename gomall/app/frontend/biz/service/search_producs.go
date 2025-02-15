package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	product "github.com/readlnh/biz-demo/gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/readlnh/biz-demo/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/readlnh/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type SearchProducsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProducsService(Context context.Context, RequestContext *app.RequestContext) *SearchProducsService {
	return &SearchProducsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProducsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	p, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{Query: req.Q})
	if err != nil {
		return nil, err
	}

	// fmt.Println(req.Q)
	return utils.H{
		"q":     req.Q,
		"items": p.Results,
	}, nil
}
