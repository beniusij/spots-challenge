package spots_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
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
	t.Run("rejects calls without expected parameters", func(t *testing.T) {
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

	t.Run("get valid json response", func(t *testing.T) {
		var result spots.Response
		r := initTestAPI()
		request, _ := http.NewRequest(http.MethodGet, "/spots?longitude=10.0&latitude=10.0&radius=10&type=circle", nil)
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

		//if len(result.Results) == 0 {
		//	t.Errorf("should contain results, but is empty")
		//}
	})
}
