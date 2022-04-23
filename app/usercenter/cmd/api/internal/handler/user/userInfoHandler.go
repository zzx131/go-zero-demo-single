package user

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/token"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/logic/user"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/response"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/svc"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/types"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取token信息
		parser := token.NewTokenParser(token.WithResetDuration(time.Second))
		tok, err := parser.ParseToken(r, svcCtx.Config.JwtAuth.AccessSecret, "")
		value := tok.Claims.(jwt.MapClaims)["key"]
		fmt.Println(value)

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
