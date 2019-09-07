package server

import (
	"log"
	"net/http"
	"wikilibras-core/src/app/routes"
	"wikilibras-core/src/config"
	"wikilibras-core/src/database"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/secure"
)

func initHandlers(db *gorm.DB) []routes.Handler {
	return []routes.Handler{
		routes.NewStateRoutes(db),
		routes.NewActionRoutes(db),
		routes.NewTaskTypeRoutes(db),
		routes.NewWorkflowRoutes(db),
		routes.NewUserRoutes(db),
		routes.NewSessionRoutes(db),
		routes.NewOrientationRoutes(db),
		routes.NewContextRoutes(db),
		routes.NewTaskRoutes(db),
		routes.NewProfileRoutes(db),
		routes.NewObjectRoutes(db),
	}
}

// Start - Initialize application
func Start() {
	db := database.RunMigrations()
	defer db.Close()

	database.RunSeeds(db)

	secureMiddleware := secure.New(secure.Options{
		FrameDeny:          true,
		BrowserXssFilter:   true,
		ContentTypeNosniff: false,
	})

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1/").Subrouter()

	for _, h := range initHandlers(db) {
		for _, route := range h.GetRoutes() {
			api.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	api.Use(secureMiddleware.Handler)

	log.Fatal(
		http.ListenAndServe(
			config.GetConfig().Address, handlers.CORS()(api),
		),
	)
}
