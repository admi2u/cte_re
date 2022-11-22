package datasets

import (
	"context"
	datasetsRpc "cte_re/service/datasets/rpc/datasets"
	"github.com/jinzhu/copier"

	"cte_re/service/datasets/api/internal/svc"
	"cte_re/service/datasets/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DatasetsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDatasetsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DatasetsListLogic {
	return &DatasetsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DatasetsListLogic) DatasetsList(in *types.DatasetListRequest) (resp *types.DatasetListResponse, err error) {
	datasetListReq := new(datasetsRpc.DatasetListReq)
	if err := copier.Copy(datasetListReq, in); err != nil {
		return nil, err
	}
	data, err := l.svcCtx.DatasetsRpc.DatasetList(l.ctx, datasetListReq)
	if err != nil {
		return nil, err
	}
	var datasetsList []*types.Dataset
	for _, item := range data.Datasets {
		datasetsList = append(datasetsList, &types.Dataset{
			Id:          item.Id,
			Name:        item.Name,
			Title:       item.Title,
			Summary:     item.Summary,
			Description: item.Description,
			DatasetSize: item.DatasetSize,
			DatasetType: item.DatasetType,
			CreateAt:    item.CreateAt,
			UpdateAt:    item.UpdateAt,
		})
	}

	return &types.DatasetListResponse{
		Total:    data.Total,
		Datasets: datasetsList,
	}, nil
}
