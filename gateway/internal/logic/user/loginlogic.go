package user

import (
	"context"
	"strings"
	"time"

	"datacenter/gateway/common/response"
	"datacenter/gateway/internal/svc"
	"datacenter/gateway/internal/types"
	"datacenter/user/rpc/userclient"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.Response, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return response.Response(response.LOGIN_PARAM_INCORRECT, response.LOGIN_PARAM_INCORRECT_MSG, nil)
		// return nil, errorx.NewCodeError(response.LOGIN_PARAM_INCORRECT, response.LOGIN_PARAM_INCORRECT_MSG)
	}
	// 使用user rpc
	userInfo, err := l.svcCtx.UserRpc.GetUserInfoByUsername(l.ctx, &userclient.UserNameRequest{
		Username: req.Username,
	})
	if err != nil {
		return response.Response(response.LOGIN_USERRPC_ERROR, response.LOGIN_PARAM_INCORRECT_MSG+":"+err.Error(), nil)
		// return nil, err
	}

	// 判断登录密码
	if userInfo.Password != hash.Md5Hex([]byte(req.Password)) {
		return response.Response(response.LOGIN_PARAM_INCORRECT, response.LOGIN_PARAM_INCORRECT_MSG, nil)
		// return nil, errorx.NewCodeError(response.LOGIN_PARAM_INCORRECT, response.LOGIN_PARAM_INCORRECT_MSG)
	}

	// jwt token
	now := time.Now().Unix()
	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(accessSecret, now, accessExpire, int64(userInfo.Id))
	if err != nil {
		return response.Response(response.LOGIN_JWTTOKEN_ERROR, response.LOGIN_JWTTOKEN_ERROR_MSG+":"+err.Error(), nil)
		// return nil, err
	}

	// return response.Response(response.SUCCESS, response.SUCCESS_MSG, map[string]interface{}{
	// 	"id":            userInfo.Id,
	// 	"username":      userInfo.Username,
	// 	"access_token":  jwtToken,
	// 	"access_expire": now + accessExpire,
	// 	"refresh_after": now + accessExpire/2,
	// })

	return response.Response(response.SUCCESS, response.SUCCESS_MSG, &types.LoginResult{
		Id:           userInfo.Id,
		Username:     userInfo.Username,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	})
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
