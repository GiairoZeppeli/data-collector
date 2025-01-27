package handler

import (
	"data-collector/service/market"
	"github.com/GiairoZeppeli/utils/context"
	"github.com/GiairoZeppeli/utils/responseWrapper"
	"github.com/GiairoZeppeli/utils/url"
	"net/http"
)

func PositionInfo(ctx context.MyContext, service market.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := url.GetQueryFirstParams(r)

		positionInfoResponse, err := service.GetPositionInfo(ctx, params)
		if err != nil {
			responseWrapper.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		responseWrapper.WriteResponseJson(w, positionInfoResponse)
	}
}

func OrderBook(ctx context.MyContext, service market.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := url.GetQueryFirstParams(r)

		orderBookResponse, err := service.GetOrderBook(ctx, params)
		if err != nil {
			responseWrapper.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		responseWrapper.WriteResponseJson(w, orderBookResponse)
	}
}

func FundingRate(ctx context.MyContext, service market.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := url.GetQueryFirstParams(r)

		fundingRateResponse, err := service.GetFundingRate(ctx, params)
		if err != nil {
			responseWrapper.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		responseWrapper.WriteResponseJson(w, fundingRateResponse)
	}
}

func DeliveryPrice(ctx context.MyContext, service market.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := url.GetQueryFirstParams(r)
		if _, ok := params["category"]; !ok {
			responseWrapper.NewErrorResponse(ctx, w, "category is required", http.StatusBadRequest)
			return
		}

		deliveryPriceResponse, err := service.GetDeliveryPrice(ctx, params)
		if err != nil {
			responseWrapper.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		responseWrapper.WriteResponseJson(w, deliveryPriceResponse)
	}
}

func CoinInfo(ctx context.MyContext, service market.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := url.GetQueryFirstParams(r)

		coinInfoResponse, err := service.GetCoinInfo(ctx, params)
		if err != nil {
			responseWrapper.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		responseWrapper.WriteResponseJson(w, coinInfoResponse)
	}
}

func Tickers(ctx context.MyContext, service market.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := url.GetQueryFirstParams(r)
		if _, ok := params["category"]; !ok {
			responseWrapper.NewErrorResponse(ctx, w, "category is required", http.StatusBadRequest)
			return
		}

		tickersResponse, err := service.GetTickers(ctx, params)
		if err != nil {
			responseWrapper.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		responseWrapper.WriteResponseJson(w, tickersResponse)
	}
}
