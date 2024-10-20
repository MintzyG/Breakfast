package yogurt

import (
	MW "breakfast/middleware"
	"fmt"
	"net/http"
)

func Run(mux *http.ServeMux) {
	fmt.Println("Connecting YogurtController")
	mux.Handle("POST /api/v1/yogurt", MW.AuthMiddleware(http.HandlerFunc(createTask)))
}
