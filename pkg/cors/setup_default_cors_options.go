package cors

import "bulletproof-go/internal/core/middleware"

func SetupDefaultCorsOptions(origins []string, methods []string, headers []string) *middleware.CORSOptions {
	var defCorsOpts = middleware.CORSOptions{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}
	if len(origins) == 0 {
		origins = append(origins, defCorsOpts.AllowedOrigins...)
	}
	if len(headers) == 0 {
		headers = append(headers, defCorsOpts.AllowedHeaders...)
	}
	if len(methods) == 0 {
		methods = append(methods, defCorsOpts.AllowedMethods...)
	}
	return middleware.NewCorsOptions(origins, methods, headers)
}
