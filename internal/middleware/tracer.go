/**
 * @Title  tracer
 * @description  链路追踪
 * @Author  沈来
 * @Update  2020/8/8 15:03
 **/
package middleware

import (
	"CloudDisk/global"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx context.Context
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path)
		}
		defer span.Finish()

		var traceID string
		var SpanID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			traceID = spanContext.(jaeger.SpanContext).TraceID().String()
			SpanID = spanContext.(jaeger.SpanContext).SpanID().String()
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", SpanID)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
