package model

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DatasetsModel = (*customDatasetsModel)(nil)

type (
	// DatasetsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDatasetsModel.
	DatasetsModel interface {
		datasetsModel
	}

	customDatasetsModel struct {
		*defaultDatasetsModel
	}
)

// NewDatasetsModel returns a model for the database table.
func NewDatasetsModel(conn sqlx.SqlConn) DatasetsModel {
	return &customDatasetsModel{
		defaultDatasetsModel: newDatasetsModel(conn),
	}
}
