package spots_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"spotlas-challenge/config"
	spots "spotlas-challenge/spots"
	//"github.com/gin-gonic/gin"
	"testing"
)

func initTestAPI() *gin.Engine {
	router := gin.New()

	router.GET("/spots", spots.GetSpotsByRadius)

	return router
}

func Test_GetSpotsByRadius(t *testing.T) {
	config.InitTestDb()

	t.Run("should reject calls without expected parameters", func(t *testing.T) {
		r := initTestAPI()
		request, _ := http.NewRequest(http.MethodGet, "/spots", nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		got := response.Body.String()
		want := "{\"message\":\"Invalid payload\"}"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if response.Code != http.StatusNotAcceptable {
			t.Errorf("got %q, want %q", response.Code, http.StatusNotAcceptable)
		}
	})

	t.Run("should reject calls with invalid parameter values", func(t *testing.T) {
		r := initTestAPI()
		request, _ := http.NewRequest(http.MethodGet, "/spots?longitude=&latitude=&radius=&type=", nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		got := response.Body.String()
		want := "{\"message\":\"Invalid payload\"}"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("get valid json response for circle", func(t *testing.T) {
		var result spots.Response
		r := initTestAPI()
		request, _ := http.NewRequest(http.MethodGet, "/spots?longitude=-8.473656&latitude=51.899216&radius=100&type=circle", nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("response failed to parse back")
		}

		expected := "ok"

		if result.Status != expected {
			t.Errorf("should contain %q, got %q", expected, result.Status)
		}

		if len(result.Results) == 0 {
			t.Errorf("should contain results, but is empty")
		}
	})

	t.Run("get valid json response for square", func(t *testing.T) {
		var result spots.Response
		r := initTestAPI()
		request, _ := http.NewRequest(http.MethodGet, "/spots?longitude=-8.473656&latitude=51.899216&radius=100&type=square", nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("response failed to parse back")
		}

		expected := "ok"

		if result.Status != expected {
			t.Errorf("should contain %q, got %q", expected, result.Status)
		}

		if len(result.Results) == 0 {
			t.Errorf("should contain results, but is empty")
		}
	})
}
