package integration

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"testing"
	"todoapp-backend/app"
	"todoapp-backend/app/model"
	"todoapp-backend/config"
	"todoapp-backend/db"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func TestTasksList(t *testing.T) {
    url := Config.Addr + "/tasks"

    tests := []struct {
        name string
        wantStatus int
    }{
        {
            name: "OK",
            wantStatus: http.StatusOK,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create random task for sanity check
            sanityTask := randomTask(DB)

            resp, err := http.Get(url)
            if err != nil {
                t.Fatal(err)
            }

            if resp.StatusCode != tt.wantStatus {
                t.Fatalf("Wrong http status. Got: %d. Want: %d", resp.StatusCode, tt.wantStatus)
            }

            if tt.wantStatus != http.StatusOK {
                return
            }

            // Make sure the sanity check task was returned
            gotTasks := []model.Task{}
            if err := json.NewDecoder(resp.Body).Decode(&gotTasks); err != nil {
                t.Fatalf("Failed to decode response: %v", err)
            }

            found := false
            for _, gotTask := range gotTasks {
                if gotTask.Title == sanityTask.Title && gotTask.Done == sanityTask.Done {
                    found = true
                }
            }

            if found == false {
                t.Fatalf("Couldn't find sanity check task in results")
            }
        })
    }
}

func TestTasksPost(t *testing.T) {
    url := Config.Addr + "/tasks"
    tests := []struct {
        name string
        body interface{}
        wantStatus int
    }{
        {
            name: "OK",
            body: map[string]interface{}{"Title": randomTitle(), "Done": randomBool()},
            wantStatus: http.StatusCreated,
        },
        {
            name: "Invalid JSON",
            body: `{"Title: "foo", "Done": false}`,
            wantStatus: http.StatusBadRequest,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            json, mErr := json.Marshal(tt.body)
            if mErr != nil {
                t.Fatal(mErr)
            }
            reader := bytes.NewReader(json)

            resp, rErr := http.Post(url, "application/json", reader)
            if rErr != nil {
                t.Fatal(rErr)
            }

            if resp.StatusCode != tt.wantStatus {
                t.Fatalf("Wrong http status. Got: %d. Want: %d", resp.StatusCode, tt.wantStatus)
            }

            // Sanity checks only past this point
            if tt.wantStatus != http.StatusCreated {
                return
            }

            createdTask := model.Task{}
            if err := DB.Where(tt.body).First(&createdTask).Error; err != nil {
                if errors.Is(err, gorm.ErrRecordNotFound) {
                    t.Fatal("Couldn't find created task in DB")
                }
                t.Fatalf("Unexpected DB error: %v", err)
            }
        })
    }
}

var Config *config.Config
var DB *gorm.DB
func init() {
    godotenv.Load("../.env")
    Config = config.GetConfig()
    DB = db.GetDB(Config)
}

func TestMain(m *testing.M) {
    app := app.NewApp(Config)
    go func() {
        app.Run()
    }()

    exitCode := m.Run()
    os.Exit(exitCode)
}
