package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo-single/user-api/internal/logic/user"
	"go-zero-demo-single/user-api/internal/svc"
	"go-zero-demo-single/user-api/internal/types"
	"go-zero-demo-single/user-api/response"
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

		response.Response(w, resp, err)
	}
}
