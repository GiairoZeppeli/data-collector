package handler

import (
	"data-collector/service/history"
	"github.com/GiairoZeppeli/utils/context"
	"github.com/GiairoZeppeli/utils/responseWrapper"
	"github.com/GiairoZeppeli/utils/url"
	"net/http"
)

func CandleHistory(ctx context.MyContext, service history.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := url.GetQueryFirstParams(r)

		candleHistoryResponse, err := service.GetCandleHistory(ctx, params)
		if err != nil {
			responseWrapper.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		responseWrapper.WriteResponseJson(w, candleHistoryResponse)
	}
}
