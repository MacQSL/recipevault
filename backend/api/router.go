package api

import (
	"database/sql"
	"net/http"
	"recipevault/api/handler"
	"recipevault/api/middleware"
	"recipevault/api/service"
	"recipevault/util"
)

// Add routes and inject dependencies to handlers
func addRoutes(mux *http.ServeMux, log *util.Logger, db *sql.DB) {

	// Initialize services
	cs := service.InitCookbookService(db)
	rs := service.InitRecipeService(db)

	// Initialize route middleware
	cookbookAuth := middleware.CookbookAuthMiddleware(log, cs)

	// Get health check
	mux.Handle("GET /health", handler.GetHealth())
	// Get all user Cookbooks
	mux.Handle("GET /cookbooks", handler.GetUserCookbooksWithRecipes(log, cs))
	// Deprecated? Get Cookbook by ID, includes Recipes
	mux.Handle("GET /cookbooks/{cookbookID}", cookbookAuth(handler.GetUserCookbook(log, cs)))
	// Deprecated? Get Cookbook Recipes
	mux.Handle("GET /cookbooks/{cookbookID}/recipes", cookbookAuth(handler.GetCookbookRecipes(log, rs)))
	// Get Recipe by ID
	mux.Handle("GET /cookbooks/{cookbookID}/recipes/{recipeID}", cookbookAuth(handler.GetRecipe(log, rs)))

	mux.Handle("/", http.NotFoundHandler())
}

// Creates a new router, adds routes and applies middleware
func NewRouter(mux *http.ServeMux, log *util.Logger, db *sql.DB) http.Handler {
	mux.Handle("/api/", http.StripPrefix("/api", mux)) // prepend /api/

	// Add above routes to the handler
	addRoutes(mux, log, db)

	// Cast as handler to wrap with middleware
	var router http.Handler = mux

	// Initialize middleware with dependencies
	logMiddleware := middleware.LoggerMiddleware(log)
	authMiddleware := middleware.AuthMiddleware(log)

	// Apply middleware to ALL routes
	router = logMiddleware(router)
	router = middleware.HeadersMiddleware(router)
	router = authMiddleware(router)

	return router
}
