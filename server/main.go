package main

import (
	"os"
	"log"
	"net/http"
	"github.com/sharizzle/my-snippets/server/db"
	"github.com/gorilla/mux"
	"github.com/sharizzle/my-snippets/server/service/code"
	customRouter "github.com/sharizzle/my-snippets/server/router"
	"github.com/joho/godotenv"
)

func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	customRouter.AppRoutes = append(customRouter.AppRoutes, code.Routes)

	for _, route := range customRouter.AppRoutes {

		//create subroute
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		//loop through each sub route
		for _, r := range route.SubRoutes {

			var handler http.Handler
			handler = r.HandlerFunc

			//attach sub route
			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}

	}

	return router
}

func main(){

	//init router
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	router := NewRouter()

	//Setup database
	db.DB = db.SetupDB()

	//create http server
	log.Fatal(http.ListenAndServe(":"+port, router))
}