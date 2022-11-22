package svc

import (
	"cte_re/service/datasets/model"
	"cte_re/service/datasets/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	DatasetsModel model.DatasetsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		DatasetsModel: model.NewDatasetsModel(sqlx.NewSqlConn("postgres", c.DB.DataSource)),
	}
}
