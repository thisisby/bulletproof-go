package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ParserOptions struct {
	size int64
}

type JSONParser struct {
	Options ParserOptions
}

func NewParserOptions(size int64) *ParserOptions {
	return &ParserOptions{
		size: size,
	}
}

func NewJsonParser(options ParserOptions) *JSONParser {
	return &JSONParser{
		Options: options,
	}
}

type JsonKey string

func (jp *JSONParser) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") == "application/json" {
			maxBodySize := jp.Options.size
			if r.ContentLength > maxBodySize {
				http.Error(w, "Request body too large", http.StatusRequestEntityTooLarge)
				return
			}
			// Read the body into a byte slice
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				log.Print(err)
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			// Restore the io.ReadCloser to its original state
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			// Decode the body into a map
			var body interface{} // map[string]interface{} global
			if err := json.Unmarshal(bodyBytes, &body); err != nil {
				log.Print(err)
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			// Store the parsed JSON in the context
			key := JsonKey("jsonBody")
			r = r.WithContext(context.WithValue(r.Context(), key, body))
		}
		next.ServeHTTP(w, r)
	})
}
