package health

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_healthCheckHandler(t *testing.T) {
	tests := []struct {
		description   string
		route         string
		method        string
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "health check return positive response",
			route:         "/api/status",
			method:        "GET",
			body:          nil,
			expectedCode:  200,
			expectedError: false,
		},
	}

	app := fiber.New()
	v1 := app.Group("/api")
	v1.Get("/status", HealthCheckHandler)
	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, test.body)
		resp, err := app.Test(req, -1) // the -1 disables request latency
		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		// Check the response body is what we expect.
		expected := `"alive":true`
		if strings.Contains(string(body), expected) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				string(body), expected)
		}
	}
}
