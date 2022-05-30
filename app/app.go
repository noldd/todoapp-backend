package app

import (
	"log"
	"net/http"
	"todoapp-backend/app/controller"
	"todoapp-backend/config"
	"todoapp-backend/db"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type App struct {
	Router *chi.Mux
	DB     *gorm.DB
	Config *config.Config
}

func NewApp(config *config.Config) *App {
	db := db.GetDB(config)
	r := chi.NewRouter()

	tasks := controller.NewTasksController(db)
	r.Route("/tasks", tasks.Router)

	return &App{
		Router: r,
		DB:     db,
		Config: config,
	}
}

func (a *App) Run() {
	log.Printf("Serving on %s", a.Config.Addr)
	log.Fatal(http.ListenAndServe(":"+a.Config.Port, a.Router))
}
