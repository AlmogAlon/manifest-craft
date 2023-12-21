package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"manifest-craft/models"
	"manifest-craft/services"
	"manifest-craft/storage"
	"manifest-craft/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

var validateResponse = Response{status: true, err: nil}

type Response struct {
	status bool
	err    error
}

type mockComponent struct{}

func (c *mockComponent) Validate(components []models.Component, jsonData map[string]interface{}) (bool, error) {
	return validateResponse.status, validateResponse.err
}

func (c *mockComponent) GetComponentOptions(component *models.Component) ([]string, error) {
	return []string{}, nil
}

func TestSend(t *testing.T) {
	r := tests.SetUpRouter()
	store := storage.NewMemoryStorage()
	s := services.Services{
		Component: &mockComponent{},
	}

	manifestController := NewManifestController(store, s)

	type testCase struct {
		name               string
		body               []byte
		expectedStatusCode int
		validateError      error
		validateStatus     bool
	}

	cases := []testCase{
		{
			// manifest not found
			name:               "not-found",
			body:               []byte{},
			expectedStatusCode: 404,
		},
		{
			// bad body given
			name:               "databaseAccess",
			body:               []byte{},
			expectedStatusCode: 400,
		},
		{
			// correct body
			name:               "databaseAccess",
			body:               []byte(`{"databaseAccess": true}`),
			expectedStatusCode: 200,
		},
		{
			// validate returned status
			name:               "databaseAccess",
			body:               []byte(`{"databaseAccess": true}`),
			validateStatus:     true,
			expectedStatusCode: 200,
		},
		{
			// returns 500 when validate returns error
			name:               "databaseAccess",
			body:               []byte(`{"databaseAccess": true}`),
			validateError:      fmt.Errorf("error"),
			expectedStatusCode: 500,
		},
	}

	r.POST("/form/:name", manifestController.Send)

	for _, tc := range cases {
		validateResponse.err = tc.validateError
		validateResponse.status = tc.validateStatus

		req, _ := http.NewRequest("POST", fmt.Sprintf("/form/%s", tc.name), bytes.NewBuffer(tc.body))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		jsonMap := make(map[string]interface{})
		err := json.Unmarshal(w.Body.Bytes(), &jsonMap)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expectedStatusCode, w.Code)

		if tc.expectedStatusCode != 200 {
			continue
		}

		assert.Equal(t, tc.validateStatus, jsonMap["status"])
	}
}
