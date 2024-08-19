package pkg

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponder_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/success", func(ctx *gin.Context) {
		responder := NewResponse(ctx)
		responder.Success("Operation successful", gin.H{"key": "value"})
	})

	req, _ := http.NewRequest("GET", "/success", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expected := `{"status":200,"message":"Operation successful","data":{"key":"value"}}`
	assert.JSONEq(t, expected, recorder.Body.String())
}

func TestResponder_Created(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/created", func(ctx *gin.Context) {
		responder := NewResponse(ctx)
		responder.Created("Resource created", gin.H{"id": 1})
	})

	req, _ := http.NewRequest("POST", "/created", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	expected := `{"status":201,"message":"Resource created","data":{"id":1}}`
	assert.JSONEq(t, expected, recorder.Body.String())
}

func TestResponder_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/badrequest", func(ctx *gin.Context) {
		responder := NewResponse(ctx)
		responder.BadRequest("Invalid request", "error details")
	})

	req, _ := http.NewRequest("GET", "/badrequest", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	expected := `{"status":400,"message":"Invalid request","error":"error details"}`
	assert.JSONEq(t, expected, recorder.Body.String())
}

func TestResponder_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/unauthorized", func(ctx *gin.Context) {
		responder := NewResponse(ctx)
		responder.Unauthorized("Unauthorized", "invalid token")
	})

	req, _ := http.NewRequest("GET", "/unauthorized", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	expected := `{"status":401,"message":"Unauthorized","error":"invalid token"}`
	assert.JSONEq(t, expected, recorder.Body.String())
}

func TestResponder_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/notfound", func(ctx *gin.Context) {
		responder := NewResponse(ctx)
		responder.NotFound("Resource not found", "details")
	})

	req, _ := http.NewRequest("GET", "/notfound", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	expected := `{"status":404,"message":"Resource not found","error":"details"}`
	assert.JSONEq(t, expected, recorder.Body.String())
}

func TestResponder_InternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/servererror", func(ctx *gin.Context) {
		responder := NewResponse(ctx)
		responder.InternalServerError("Internal error", "error details")
	})

	req, _ := http.NewRequest("GET", "/servererror", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	expected := `{"status":500,"message":"Internal error","error":"error details"}`
	assert.JSONEq(t, expected, recorder.Body.String())
}
