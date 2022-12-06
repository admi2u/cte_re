package logic

import (
	"context"
	"datasets/rpc/datasets"
	"datasets/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DatasetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDatasetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DatasetListLogic {
	return &DatasetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DatasetListLogic) DatasetList(in *datasets.DatasetListReq) (*datasets.DatasetListResp, error) {
	data, err := l.svcCtx.DatasetsModel.GetPageList(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}
	var list []*datasets.DatasetListData
	for _, item := range data {
		list = append(list, &datasets.DatasetListData{
			Id:          item.Id,
			Name:        item.Name,
			Title:       item.Title,
			Summary:     item.Summary,
			Description: item.Description,
			DatasetSize: item.DatasetSize,
			DatasetType: item.DatasetType,
			CreateAt:    item.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:    item.UpdateAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &datasets.DatasetListResp{
		Total:    int64(len(list)),
		Datasets: list,
	}, nil
}
