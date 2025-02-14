package router

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{

	Route{
		"ApiAuthPost",
		strings.ToUpper("Post"),
		"/api/auth",
		apiAuthPost,
	},

	Route{
		"ApiBuyItemGet",
		strings.ToUpper("Get"),
		"/api/buy/{item}",
		apiBuyItemGet,
	},

	Route{
		"ApiInfoGet",
		strings.ToUpper("Get"),
		"/api/info",
		apiInfoGet,
	},

	Route{
		"ApiSendCoinPost",
		strings.ToUpper("Post"),
		"/api/sendCoin",
		apiSendCoinPost,
	},
}
