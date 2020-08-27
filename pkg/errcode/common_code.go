/**
 * @Title  common_code
 * @Description  公共错误码
 * @Author  沈来
 * @Update  2020/8/3 16:52
 **/
package errcode

var(
	Success                   = NewError(10000,"成功")
	ServerError               = NewError(10001,"服务器内部错误")
	InvalidParams             = NewError(10002,"入参错误")
	NotFound                  = NewError(10003,"找不到")
	UnauthorizedAuthNotExist  = NewError(10004,"鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10005,"鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10006,"鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10007,"鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10008,"请求过多")
)