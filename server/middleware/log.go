package middleware

import (
  "io"
  "net/http"
  "github.com/gorilla/handlers"
)

  func LoggingHandler(dst io.Writer) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
	  return handlers.LoggingHandler(dst, h)
	}
  }