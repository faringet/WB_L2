package main

import (
	"CitysTempRest/controllers"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubsCreate(t *testing.T) {
	router := gin.New()

	router.POST("/", controllers.SubsCreate)

	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{
        "City": "New York",
    }`))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	assert.JSONEq(t, `{
        "post": {
            "City": "New York",
        }
    }`, res.Body.String())
}
