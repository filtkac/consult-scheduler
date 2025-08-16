package integrationtest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func makeRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
