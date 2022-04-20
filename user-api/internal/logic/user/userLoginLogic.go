package user

import (
	"context"
	"github.com/golang-jwt/jwt/v4"

	"go-zero-demo-single/user-api/internal/svc"
	"go-zero-demo-single/user-api/internal/types"

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

func (l *UserLoginLogic) UserLogin(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// 申请JWT token
	l.Logger.Infof("登录入参是：", *req)
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, 10, 1000, 10)
	return &types.UserInfoResp{
		UserId:      10,
		Nickname:    "zhangsan",
		AccessToken: accessToken,
	}, err
}

// getJwtToken 获取JWT token
func (l *UserLoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)

	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
