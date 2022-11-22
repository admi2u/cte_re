package svc

import (
	"cte_re/service/datasets/api/internal/config"
	"cte_re/service/datasets/rpc/datasetsclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	DatasetsRpc datasetsclient.Datasets
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		DatasetsRpc: datasetsclient.NewDatasets(zrpc.MustNewClient(c.DatasetsRpcConf)),
	}
}
