package routes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todoapp-backend/app/model"
	"todoapp-backend/config"
	"todoapp-backend/db"
	"todoapp-backend/testutil"

	"gorm.io/gorm"
)

func TestTasks_ListIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }

    db := db.GetDB(config.GetConfig())

	tests := []struct {
		name       string
        wantStatus     int
	}{
        {
            name: "OK",
            wantStatus: 200,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tasks{
				DB: db,
			}

            w := httptest.NewRecorder()

            r := httptest.NewRequest(
                http.MethodGet,
                "/",
                nil,
            )
            r.Close = true

			tr.List(w, r)

            if gotStatus := w.Result().StatusCode; gotStatus != tt.wantStatus {
                t.Fatalf("Got %d, want %d", gotStatus, tt.wantStatus)
            }

            // Make sure that the initial test data is found in the response
            var tasks []model.Task
            decoder := json.NewDecoder(w.Result().Body)
            if err := decoder.Decode(&tasks); err != nil {
                t.Fatalf("Error decoding response: %v", err)
            }

            found := false
            for _, task := range tasks {
                if task.Title == "Foo" && task.Done == true {
                    found = true
                    break
                }
            }

            if found == false {
                t.Error("Couldn't find a done Foo task in results")
            }
		})
	}
}


func TestTasks_PostIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }

    db := db.GetDB(config.GetConfig())

	tests := []struct {
		name       string
		body   interface {}
        wantStatus     int
	}{
        {
            name: "OK",
            body: map[string]interface{}{"Title": "foo", "Done": false},
            wantStatus: 201,
        },
        {
            name: "OK",
            body: map[string]interface{}{"Title": "bar", "Done": true},
            wantStatus: 201,
        },
        {
            name: "Invalid JSON",
            body: `{"Title: "foo", "Done": false}`,
            wantStatus: 400,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tasks{
				DB: db,
			}

            w := httptest.NewRecorder()

            bytes, err := json.Marshal(tt.body)
            if err != nil {
                t.Fatalf("Failed to marshal JSON: %v", err)
            }
            reader := strings.NewReader(string(bytes))

            r := httptest.NewRequest(
                http.MethodPost,
                "/",
                reader,
            )
            r.Close = true

			tr.Post(w, r)

            resp := w.Result()
            defer resp.Body.Close()
            if gotStatus := resp.StatusCode; gotStatus != tt.wantStatus {
                t.Fatalf("Got %d, want %d", gotStatus, tt.wantStatus)
                return
            }

            // If the request was successful, then make sure the task was
            // created in the database.
            if tt.wantStatus != http.StatusCreated {
                return
            }

            body, readErr := ioutil.ReadAll(resp.Body)
            if readErr != nil {
                t.Fatalf("Failed to read response body: %v", readErr)
            }

            var task model.Task
            if err := json.Unmarshal(body, &task); err != nil {
                t.Fatalf("Failed to parse response body: %v", err)
            }

            var foundTask model.Task
            dbErr := tr.DB.First(&foundTask, task.ID).Error
            if dbErr != nil {
                switch {
                case errors.Is(dbErr, gorm.ErrRecordNotFound):
                    t.Fatalf("Couldn't find created task")
                default:
                    t.Fatalf("Unexpected error: %v", dbErr)
                }
            }
		})
	}
}

func init() {
    testutil.SetupDB()
}
