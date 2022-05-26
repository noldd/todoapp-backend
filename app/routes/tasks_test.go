package routes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"todoapp-backend/app/model"
	"todoapp-backend/config"
	"todoapp-backend/db"
	"todoapp-backend/testutil"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func TestTasks_Router(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		r chi.Router
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tasks{
				DB: tt.fields.DB,
			}
			tr.Router(tt.args.r)
		})
	}
}

func TestNewTasksRouter(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *Tasks
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTasksRouter(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTasksRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTasks_List(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tasks{
				DB: tt.fields.DB,
			}
			tr.List(tt.args.w, tt.args.r)
		})
	}
}

func TestTasks_ListIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }

    testutil.SetEnv()
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
		})
	}
}


func TestTasks_PostIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }

    testutil.SetEnv()
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
                t.Fatalf("Failed to read response body: %v", err)
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
