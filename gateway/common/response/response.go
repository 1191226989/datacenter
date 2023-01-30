package response

import "datacenter/gateway/internal/types"

// 封装统一格式的body响应
func Response(code int, msg string, data interface{}) (*types.Response, error) {
	return &types.Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}, nil
}
