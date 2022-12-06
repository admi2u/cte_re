package logic

import (
	"context"
	"datasets/model"

	"datasets/rpc/datasets"
	"datasets/rpc/internal/svc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DatasetCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDatasetCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DatasetCreateLogic {
	return &DatasetCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DatasetCreateLogic) DatasetCreate(in *datasets.DatasetCreateReq) (*datasets.DatasetCreateResp, error) {
	dataset := new(model.Datasets)
	if err := copier.Copy(dataset, in); err != nil {
		return nil, err
	}
	sqlResult, err := l.svcCtx.DatasetsModel.Insert(l.ctx, dataset)
	if err != nil {
		return nil, err
	}
	lastInsertId, err := sqlResult.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &datasets.DatasetCreateResp{
		Id: lastInsertId,
	}, nil
}
