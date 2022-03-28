package calculate

import (
	"context"
	"github.com/go-kit/log"
)

// 服务中间件， 执行服务某项操作之前执行操作
type CalculateMiddleware struct {
	logger log.Logger
	next   Service
}

func (c CalculateMiddleware) Calculate(ctx context.Context, in CalculateIn) (rs ResultAck) {
	// something
	return c.next.Calculate(ctx, in)
}

func NewCalculateMiddleware(logger log.Logger) ServiceMiddleware {
	return func(service Service) Service {
		return &CalculateMiddleware{
			logger: logger,
			next:   service,
		}
	}
}

// ServiceMiddleware define service middleware
type ServiceMiddleware func(Service) Service
