// Package routes provides HTTP routing configuration and endpoint registration.
//
// This package defines the routing structure for the Go-Kopi framework,
// including route registration, middleware assignment, and router initialization.
//
// Usage:
//
//	api := api_http.NewHTTPs(usecases)
//	routers := routes.NewEndpointRouters(api)
//	routes.InitRouter(routers, muxRouter)
package routes

import (
	"log"
	"net/http"

	"github.com/RandySteven/paipai-deposit/enums"
	rest_handler "github.com/RandySteven/paipai-deposit/handlers/rests"
	"github.com/RandySteven/paipai-deposit/middlewares"
	"github.com/gorilla/mux"
)

type (
	// HandlerFunc is the standard HTTP handler function signature.
	HandlerFunc func(w http.ResponseWriter, r *http.Request)

	// Router represents a single route configuration including its path,
	// HTTP method, handler function, and associated middlewares.
	//
	// Fields:
	//   - methodName: Identifier name for the route (used for logging)
	//   - path: URL path pattern (e.g., "/users", "/users/{id}")
	//   - handler: The handler function to execute
	//   - method: HTTP method (GET, POST, PUT, DELETE, PATCH)
	//   - middlewares: List of middleware enums to apply to this route
	Router struct {
		methodName  string
		path        string
		handler     HandlerFunc
		method      string
		middlewares []enums.Middleware
	}

	// RouterPrefix maps route prefixes to their associated routes.
	// This allows grouping routes under common URL prefixes.
	//
	// Example:
	//
	//	RouterPrefix{
	//	    enums.AuthPrefix: []*Router{...},    // /auth/*
	//	    enums.UsersPrefix: []*Router{...},   // /users/*
	//	}
	RouterPrefix map[enums.RouterPrefix][]*Router
)

// NewEndpointRouters creates and returns all API endpoint configurations.
// This is where all routes are defined and grouped by their prefix.
//
// Parameters:
//   - api: The HTTPs struct containing all HTTP handlers
//
// Returns:
//   - RouterPrefix: Map of all configured routes grouped by prefix
//
// Example - Adding new routes:
//
//	// Authentication routes (prefix: /auth)
//	endpointRouters[enums.AuthPrefix] = []*Router{
//	    Post("RegisterUser", "/register", api.UserHTTP.RegisterUser),
//	    Post("LoginUser", "/login", api.UserHTTP.LoginUser),
//	}
//
//	// User routes (prefix: /users)
//	endpointRouters[enums.UsersPrefix] = []*Router{
//	    Get("GetUsers", "/", api.UserHTTP.GetUsers, enums.AuthenticationMiddleware),
//	    Get("GetUserByID", "/{id}", api.UserHTTP.GetUserByID),
//	    Put("UpdateUser", "/{id}", api.UserHTTP.UpdateUser, enums.AuthenticationMiddleware),
//	    Delete("DeleteUser", "/{id}", api.UserHTTP.DeleteUser, enums.AuthenticationMiddleware),
//	}
func NewEndpointRouters(api *rest_handler.Deposits) RouterPrefix {
	endpointRouters := make(RouterPrefix)

	// ============================================================
	// AUTH ROUTES - Prefix: /auth
	// ============================================================
	// POST /auth/register - Register a new user account
	// POST /auth/login    - Authenticate user and get token
	// ============================================================
	endpointRouters[enums.DepositPrefix] = []*Router{
		// Post("LoginUser", "/login", api.UserRest.LoginUser),
		Post("CreateAccount", "/accounts", api.CreateAccount),
		Post("AccountInquiry", "/accounts/{account_number}", api.AccountInquiry),
		Post("BalanceInquiry", "/accounts/{account_number}/balance", api.BalanceInquiry),
		Post("Auth", "/accounts/{account_number}/auth", api.Auth),
		Post("Capture", "/accounts/{account_number}/capture", api.Capture),
		Post("TransactionHistory", "/accounts/{account_number}/transaction-history", api.TransactionHistory),
		Post("TransactionDetail", "/accounts/{account_number}/transaction-detail", api.TransactionDetail),
	}

	return endpointRouters
}

// InitRouter initializes the mux router with all configured routes and middlewares.
// It sets up global middlewares and registers all endpoint routes.
//
// Global middlewares applied (in order):
//   - LoggingMiddleware: Logs all incoming requests
//   - CorsMiddleware: Handles CORS headers
//   - TimeoutMiddleware: Sets request timeout
//   - CheckHealthMiddleware: Health check endpoint
//   - AuthenticationMiddleware: JWT authentication
//   - RateLimiterMiddleware: Rate limiting
//
// Parameters:
//   - routers: RouterPrefix containing all route configurations
//   - r: The gorilla/mux router instance
func InitRouter(routers RouterPrefix, r *mux.Router) {
	middleware := middlewares.NewMiddlewares()
	clientMiddleware := middlewares.RegisterClientMiddleware(middleware)
	serverMiddleware := middlewares.RegisterServerMiddleware(middleware)

	r.Use(
		serverMiddleware.LoggingMiddleware,
		serverMiddleware.CorsMiddleware,
		serverMiddleware.TimeoutMiddleware,
		serverMiddleware.CheckHealthMiddleware,
		clientMiddleware.AuthenticationMiddleware,
		clientMiddleware.RateLimiterMiddleware,
	)

	depositsRouter := r.PathPrefix(enums.AuthPrefix.ToString()).Subrouter()
	for _, routers := range routers[enums.DepositPrefix] {
		routers.RouterLog(enums.DepositPrefix.ToString())
		depositsRouter.HandleFunc(routers.path, routers.handler).Methods(routers.method)
	}
}

// RouterLog prints route information to the console during startup.
// Format: "METHOD | /prefix/path"
func (router *Router) RouterLog(prefix string) {
	log.Printf("%12s | %4s/ \n", router.method, prefix+router.path)
}
