package application_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kiskislaya/calc_go/internal/application"
)

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    map[string]string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "valid expression",
			requestBody: map[string]string{
				"expression": "1+1",
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"result": 2.0,
			},
		},
		{
			name: "incorrect operator",
			requestBody: map[string]string{
				"expression": "2+2*",
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody: map[string]interface{}{
				"error": "incorrect operator",
			},
		},
		{
			name: "invalid expression",
			requestBody: map[string]string{
				"expression": "3+3a",
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody: map[string]interface{}{
				"error": "invalid expression",
			},
		},
		{
			name: "division by zero",
			requestBody: map[string]string{
				"expression": "4/0",
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody: map[string]interface{}{
				"error": "division by zero",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(application.CalcHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			var responseBody map[string]interface{}
			err = json.NewDecoder(rr.Body).Decode(&responseBody)
			if err != nil {
				t.Fatal(err)
			}

			for key, expectedValue := range tt.expectedBody {
				if responseBody[key] != expectedValue {
					t.Errorf("handler returned unexpected body: got %v want %v",
						responseBody[key], expectedValue)
				}
			}
		})
	}
}
