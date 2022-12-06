package datasets

import (
	"net/http"

	"datasets/api/internal/logic/datasets"
	"datasets/api/internal/svc"
	"datasets/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DatasetCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DatasetCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := datasets.NewDatasetCreateLogic(r.Context(), svcCtx)
		resp, err := l.DatasetCreate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
