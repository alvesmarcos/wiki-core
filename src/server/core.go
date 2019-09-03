package server

import (
	"log"
	"net/http"
	"wikilibras-core/src/app/routes"
	"wikilibras-core/src/database"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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
	}
}

// Start - Initialize application
func Start(withSeeds bool) {
	db := database.RunMigrations()

	if withSeeds {
		database.RunSeeds(db)
	}
	defer db.Close()

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1/").Subrouter()

	for _, h := range initHandlers(db) {
		for _, route := range h.GetRoutes() {
			api.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(api)))
}
