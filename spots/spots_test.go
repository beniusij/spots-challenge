package spots_test

import (
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

func TestGETSpots(t *testing.T) {
	t.Run("rejects calls without expected parameters", func(t *testing.T) {
		r := initTestAPI()
		request, _ := http.NewRequest(http.MethodGet, "/spots?latitude=&longitude=&radius=&type=", nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		got := response.Body.String()
		want := "{\"message\":\"Invalid payload\"}"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
