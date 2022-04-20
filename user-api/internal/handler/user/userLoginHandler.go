package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo-single/user-api/internal/logic/user"
	"go-zero-demo-single/user-api/internal/svc"
	"go-zero-demo-single/user-api/internal/types"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
