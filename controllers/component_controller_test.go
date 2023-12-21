package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"manifest-craft/services"
	"manifest-craft/storage"
	"manifest-craft/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetValues(t *testing.T) {
	r := tests.SetUpRouter()
	store := storage.NewMemoryStorage()
	s := services.Get()

	componentController := NewComponentController(store, s)

	type testCase struct {
		source     string
		statusCode int
		want       []string
	}

	cases := []testCase{
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

	for _, tc := range cases {

		req, _ := http.NewRequest("GET", fmt.Sprintf("/values/%s", tc.source), nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, tc.statusCode, w.Code)

		if tc.statusCode != 200 {
			continue
		}

		jsonMap := make(map[string][]string)
		err := json.Unmarshal(w.Body.Bytes(), &jsonMap)

		assert.NoError(t, err)
		assert.Equal(t, tc.want, jsonMap["values"])
	}
}
