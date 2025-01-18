package middleware

import (
	"github.com/GiairoZeppeli/utils/context"
	"github.com/GiairoZeppeli/utils/responseWrapper"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoveryMiddleware(ctx context.MyContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic occurred: %v\n%s", err, debug.Stack())
				responseWrapper.NewErrorResponse(ctx, w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
