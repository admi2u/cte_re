package datasets

import (
	"net/http"

	"cte_re/service/datasets/api/internal/logic/datasets"
	"cte_re/service/datasets/api/internal/svc"
	"cte_re/service/datasets/api/internal/types"
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
