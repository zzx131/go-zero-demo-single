package user

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/svc"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// 申请JWT token
	l.Logger.Infof("登录入参是：", req)
	accessToken, err := l.buildToken(l.svcCtx.Config.JwtAuth.AccessSecret, map[string]interface{}{
		"key": "value",
	}, 60*5)
	return &types.UserLoginResp{
		UserId:      10,
		Nickname:    "zhangsan",
		AccessToken: accessToken,
	}, err
}

func (l *UserLoginLogic) buildToken(secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + seconds
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
