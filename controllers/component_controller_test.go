package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"manifest-craft/services"
	"manifest-craft/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetValues(t *testing.T) {
	r := SetUpRouter()
	store := storage.NewMemoryStorage()
	s := services.Get()

	componentController := NewComponentController(store, s)

	type testCase struct {
		source     string
		statusCode int
		want       []string
	}

	tests := []testCase{
		{
			source:     "databaseInstances",
			statusCode: 200,
			want:       []string{"Prod", "Local", "staging"},
		},
		{
			source:     "not-found",
			statusCode: 404,
			want:       []string{},
		},
	}

	r.GET("/values/:source", componentController.GetValues)

	for _, tc := range tests {

		req, _ := http.NewRequest("GET", fmt.Sprintf("/values/%s", tc.source), nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, tc.statusCode, w.Code)

		if tc.statusCode != 200 {
			return
		}

		jsonMap := make(map[string][]string)
		err := json.Unmarshal(w.Body.Bytes(), &jsonMap)
		if err != nil {
			t.Failed()
		}
		assert.Equal(t, tc.want, jsonMap["values"])
	}
}
