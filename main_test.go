package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=52&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	totalCount := 4
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Equal(t, len(list), totalCount)
}

func TestMainHandlerWhenOkAndNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
}

func TestMainHandlerWhenMissingCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=saint-petersburg", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	expected := `wrong city value`
	body := responseRecorder.Body.String()
	assert.Equal(t, body, expected)
}
