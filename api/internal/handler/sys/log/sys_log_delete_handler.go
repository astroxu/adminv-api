package log

import (
	"net/http"

	"adminv-api/api/internal/logic/sys/log"
	"adminv-api/api/internal/svc"
	"adminv-api/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SysLogDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteSysLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := log.NewSysLogDeleteLogic(r.Context(), svcCtx)
		resp, err := l.SysLogDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
