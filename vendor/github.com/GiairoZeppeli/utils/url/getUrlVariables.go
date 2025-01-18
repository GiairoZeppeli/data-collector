package url

import "net/http"

func GetQueryFirstParams(r *http.Request) map[string]interface{} {
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
