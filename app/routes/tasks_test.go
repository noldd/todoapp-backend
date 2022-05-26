package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
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

func TestTasks_PostIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }

    testutil.SetEnv()
    db := db.GetDB(config.GetConfig())

	tests := []struct {
		name       string
		bodyJson   interface {}
        bodyText   string
        wantStatus     int
	}{
        {
            name: "OK",
            bodyJson: map[string]interface{}{"Title": "foo", "Done": false},
            wantStatus: 201,
        },
        {
            name: "OK",
            bodyJson: map[string]interface{}{"Title": "bar", "Done": true},
            wantStatus: 201,
        },
        {
            name: "Invalid JSON",
            bodyJson: `{"Title: "foo", "Done": false}`,
            wantStatus: 400,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tasks{
				DB: db,
			}

            w := httptest.NewRecorder()

            bytes, err := json.Marshal(tt.bodyJson)
            if err != nil {
                t.Fatalf("Failed to marshal JSON: %v", err)
            }
            reader := strings.NewReader(string(bytes))

            r := httptest.NewRequest(
                http.MethodPost,
                "/",
                reader,
            )

			tr.Post(w, r)

            if gotStatus := w.Result().StatusCode; gotStatus != tt.wantStatus {
                t.Fatalf("Got %d, want %d", gotStatus, tt.wantStatus)
            }
		})
	}
}
