package testutil

import "os"

func SetEnv() {
    envs := map[string]string {
		"DB_ENDPOINT": "localhost:3306",
		"DB_USER":     "todos",
		"DB_PASSWORD": "todos",
		"DB_NAME":     "todos",
    }

    for key, value := range envs {
        os.Setenv(key, value)
    }
}
