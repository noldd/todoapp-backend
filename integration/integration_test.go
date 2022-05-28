package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"todoapp-backend/app"
	"todoapp-backend/config"

	"github.com/joho/godotenv"
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
            resp, err := http.Get(url)
            if err != nil {
                t.Fatal(err)
            }

            if resp.StatusCode != tt.wantStatus {
                t.Fatalf("Wrong http status. Got: %d. Want: %d", resp.StatusCode, tt.wantStatus)
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
            body: map[string]interface{}{"Title": "foo", "Done": false},
            wantStatus: http.StatusCreated,
        },
        {
            name: "OK",
            body: map[string]interface{}{"Title": "bar", "Done": true},
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
        })
    }
}

var Config *config.Config
func init() {
    godotenv.Load("../.env")
    Config = config.GetConfig()
}

func TestMain(m *testing.M) {
    app := app.NewApp(Config)
    go func() {
        app.Run()
    }()

    exitCode := m.Run()
    os.Exit(exitCode)
}
