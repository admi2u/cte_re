package datasets

import (
	"datasets/api/internal/types"
	"net/http"

	"datasets/api/internal/logic/datasets"
	"datasets/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DatasetsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DatasetListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := datasets.NewDatasetsListLogic(r.Context(), svcCtx)
		resp, err := l.DatasetsList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
