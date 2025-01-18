package context

import (
	"context"
	"go.uber.org/zap"
	"net/http"
)

type MyContext struct {
	Ctx    context.Context
	Logger *zap.SugaredLogger
}

func NewMyContext(ctx context.Context, logger *zap.SugaredLogger) MyContext {
	return MyContext{
		Ctx:    ctx,
		Logger: logger,
	}
}

func (ctx MyContext) GetQueryFirstParams(r *http.Request) map[string]interface{} {
	queryParams := r.URL.Query()
	if len(queryParams) == 0 {
		return nil
	}

	result := make(map[string]interface{}, len(queryParams))
	for key, values := range queryParams {
		if len(values) > 0 {
			result[key] = values[0]
		}
	}
	return result
}
