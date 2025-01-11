package handler

import (
	"github.com/GiairoZeppeli/utils/context"
	"github.com/GiairoZeppeli/utils/responseWrapper"
	"net/http"
)

func Ping(ctx context.MyContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := responseWrapper.WriteResponse(w, http.StatusOK, responseWrapper.StatusResponse{Status: "pong"}); err != nil {
			responseWrapper.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
		}
	}

}
