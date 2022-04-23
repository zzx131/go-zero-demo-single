package user

import (
	"go-zero-demo-single/app/usercenter/cmd/api/internal/logic/user"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/response"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/svc"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)

		response.Response(w, resp, err)
	}
}
