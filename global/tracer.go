/**
 * @Title  tracer
 * @description  链路追踪
 * @Author  沈来
 * @Update  2020/8/8 14:58
 **/
package global

import "github.com/opentracing/opentracing-go"

var (
	Tracer opentracing.Tracer
)
