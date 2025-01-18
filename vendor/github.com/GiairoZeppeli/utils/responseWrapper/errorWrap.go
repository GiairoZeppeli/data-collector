package responseWrapper

import (
	"encoding/json"
	"github.com/GiairoZeppeli/utils/context"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx context.MyContext, w http.ResponseWriter, message string, statusCode int) {
	ctx.Logger.Error(message)

	errRes := ErrorResponse{Message: message}

	jsonErrRes, err := json.Marshal(errRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(statusCode)
	w.Write(jsonErrRes)
}
