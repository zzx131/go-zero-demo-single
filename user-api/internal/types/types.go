// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoReq struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName,optional"`
	Password string `json:"password,optional"`
}

type UserInfoResp struct {
	UserId      int64  `json:"userId"`
	Nickname    string `json:"nickname"`
	AccessToken string `json:"accessToken"`
}