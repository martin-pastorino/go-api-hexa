package middleware

import (
	"api/core"
	"context"
	"net/http"
)

const (
	AUTHORIZATION = "Authorization"
	PLATFORM	  = "Platform"
	X_API_EY= "X-API-KEY"
)



type RequiredHeadersMiddleware struct {
	headers map[string]int
}

func RequiredHeaders(next http.Handler) http.Handler {

	requireHeaders := RequiredHeadersMiddleware{	headers: map[string]int{
		AUTHORIZATION: http.StatusUnauthorized,
		PLATFORM: http.StatusBadRequest,
		X_API_EY: http.StatusBadRequest,
	}}


	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requiredHeadersValues := make(map[string]string)

		for header, status := range requireHeaders.headers {
			value := r.Header.Get(header)
			if value == "" {
				http.Error(w, header + " header is required", status)
				return
			} 
			requiredHeadersValues[header] = value
			
		}
		
		ctx = context.WithValue(ctx, core.ContextKey("requiredHeaders"), requiredHeadersValues)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
