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
	"github.com/rs/cors"
	"github.com/sharizzle/my-snippets/server/middleware"
	
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

			// Apply Cors Middleware
			c := cors.New(cors.Options{
				AllowedOrigins: []string{"http://localhost", "http://localhost:3000"},
				ExposedHeaders: []string{"X-Total-Count","Access-Control-Allow-Origin"},
				AllowCredentials: true,
				AllowedMethods: []string{"GET","POST","PUT","DELETE","PATCH","OPTIONS"},
				OptionsPassthrough: true,
				Debug: false,
			})
			handler = c.Handler(handler)

			// Apply Logging Middlware
			logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
  			if err != nil {
    		log.Fatal(err)
  			}
			loggingHandler := middleware.LoggingHandler(logFile)

			//attach sub route
			routePrefix.
				Path(r.Pattern).
				Handler(loggingHandler(handler)).
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
	db.DB.AutoMigrate(&code.Code{})


	//create http server
	log.Fatal(http.ListenAndServe(":"+port, router))
}