package api

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckHealth(t *testing.T) {
	t.Run("health check - OK", func(t *testing.T) {
		r := httptest.NewRequest("GET", "http://api.com/status", nil)
		w := httptest.NewRecorder()

		CheckHealth(w, r)
		response := w.Result()
		body, err := ioutil.ReadAll(response.Body)
		statusCode := response.StatusCode
		contentType := response.Header.Get("Content-Type")

		assert.Equal(t, string(body), "health check - OK")
		assert.Equal(t, err, nil)
		assert.Equal(t, statusCode, 200)
		assert.Equal(t, contentType, "text/plain; charset=utf-8")
	})
}
