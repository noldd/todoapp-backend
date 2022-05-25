package app

import (
	"fmt"
	"log"
	"net/http"
	"todoapp-backend/app/model"
	"todoapp-backend/app/routes"
	"todoapp-backend/config"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Router *chi.Mux
	DB     *gorm.DB
}

func NewApp(config *config.Config) *App {
	dbURI := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Endpoint,
		config.DB.Name,
		config.DB.Charset,
	)

	dbDialector := mysql.New(mysql.Config{
		DSN: dbURI,
	})

	db, err := gorm.Open(dbDialector)
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	model.DBMigrate(db)

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
