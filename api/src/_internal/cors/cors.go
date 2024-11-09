package cors

import (
  "net/http"
)

func EnableCors(w *http.ResponseWriter) {
  (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}
