package app

import (
	"log"
	"net/http"
	"todoapp-backend/app/routes"
	"todoapp-backend/config"
	"todoapp-backend/db"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type App struct {
	Router *chi.Mux
	DB     *gorm.DB
}

func NewApp(config *config.Config) *App {
    db := db.GetDB(config)
	r := chi.NewRouter()

	tasks := routes.NewTasksRouter(db)
	r.Route("/tasks", tasks.Router)

	return &App{
		Router: r,
		DB:     db,
	}
}

func (a *App) Run(host string) {
	log.Printf("Serving on http://localhost%s", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
