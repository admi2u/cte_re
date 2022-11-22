package datasets

import (
	"context"
	"cte_re/service/datasets/rpc/datasets"
	"github.com/jinzhu/copier"

	"cte_re/service/datasets/api/internal/svc"
	"cte_re/service/datasets/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DatasetCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDatasetCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DatasetCreateLogic {
	return &DatasetCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DatasetCreateLogic) DatasetCreate(req *types.DatasetCreateRequest) (resp *types.DatasetCreateResponse, err error) {
	datasetCreateReq := new(datasets.DatasetCreateReq)
	if err := copier.Copy(datasetCreateReq, req); err != nil {
		return nil, err
	}
	datasetCreateResp, err := l.svcCtx.DatasetsRpc.DatasetCreate(l.ctx, datasetCreateReq)
	if err != nil {
		return nil, err
	}

	return &types.DatasetCreateResponse{
		Id: datasetCreateResp.Id,
	}, nil
}
