package responseWrapper

import "net/http"

func WriteResponseJson(w http.ResponseWriter, response string) {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(response))
}
