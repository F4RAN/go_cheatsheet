package httpserver

import (
	"fmt"
	"log"
	"net/http"
)

// 1. Simple HandleFunc
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// 2. Controller-style handler method
type UserController struct{}

func (c *UserController) profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Profile"))
}

// 3. Handler object: implements ServeHTTP
type AdminHandler struct{}

func (h *AdminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin Area"))
}

// 4. Middleware: accepts Handler, returns Handler
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// 5. Wildcard route: captures the rest of the path
func wildcardHandler(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")

	w.Write([]byte("Wildcard path: " + path))
}

func main() {
	router := http.NewServeMux()

	userController := &UserController{}
	adminHandler := &AdminHandler{}

	// Usecase 1: simple function route
	router.HandleFunc("/health", healthHandler)

	// Usecase 2: controller method route
	router.HandleFunc("/profile", userController.profileHandler)

	// Usecase 3: object that implements http.Handler
	router.Handle("/admin", adminHandler)

	// Usecase 4: route with middleware
	router.Handle("/secure-profile",
		AuthMiddleware(
			http.HandlerFunc(userController.profileHandler),
		),
	)
	
	// Usecase 5: wildcard route
	router.HandleFunc("/files/{path...}", wildcardHandler)
	
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
