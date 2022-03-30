package product

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
)

func HandlerList(ctx context.Context, in *proto.ProductListRequest) (*proto.ProductListResponse, aerror.Error) {
	if in == nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "request is invalid")
	}

	total, products, err := model.ProductSearch(ctx, in.Query, (in.Page-1)*in.Limit, in.Limit)
	if err != nil {
		return nil, err
	}

	logger.Infof("total: %d", total)
	logger.Infof("products: %s", logger.ToJson(products))

	productCategoryMap, err := getProductIdCategoryMap(ctx, products)
	if err != nil {
		return nil, err
	}

	var productList []*proto.Product
	for _, p := range products {
		category, ok := productCategoryMap[p.ID]
		if !ok {
			logger.Waringf("product: %d has no category", p.ID)
			continue
		}

		productList = append(productList, &proto.Product{
			ID:         p.ID,
			Name:       p.Name,
			Desc:       p.Desc,
			CategoryId: category.ID,
		})
	}

	resp := &proto.ProductListResponse{
		Page:    in.Page,
		Limit:   in.Limit,
		Total:   total,
		HasNext: (in.Page-1)*in.Limit+int64(len(productList)) < total,
		Data:    productList,
	}

	return resp, nil
}

func getProductIdCategoryMap(ctx context.Context, products []*charon.Product) (map[int64]*charon.CategoryWide, aerror.Error) {
	var productIds []int64
	for _, p := range products {
		productIds = append(productIds, p.ID)
	}

	categories, err := model.SearchCategoryByProductIds(ctx, productIds)
	if err != nil {
		return nil, err
	}

	productIdCategoryMap := make(map[int64]*charon.CategoryWide)
	for _, c := range categories {
		productIdCategoryMap[c.ProductId] = c
	}

	return productIdCategoryMap, nil
}