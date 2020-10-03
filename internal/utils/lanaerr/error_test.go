package lanaerr_test

import (
	"errors"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/utils/lanaerr"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	// Given
	err := errors.New("my custom error")
	statusCode := http.StatusNotFound

	// When
	lanaErr := lanaerr.New(err, statusCode)

	// Then
	assert.EqualError(t, lanaErr, "my custom error")
	assert.EqualError(t, lanaErr.GetError(), "my custom error")
	assert.Equal(t, 404, lanaErr.GetStatusCode())
}

func TestNewBuild(t *testing.T) {
	// Given
	err := errors.New("my custom error")
	statusCode := http.StatusNotFound

	// When
	lanaErr := lanaerr.Empty().WithCode(statusCode).WithErr(err)

	// Then
	assert.EqualError(t, lanaErr, "my custom error")
	assert.EqualError(t, lanaErr.GetError(), "my custom error")
	assert.Equal(t, 404, lanaErr.GetStatusCode())
}

func TestNewFromLanaError(t *testing.T) {
	// Given
	err := errors.New("my custom error")
	statusCode := http.StatusNotFound
	lErr := lanaerr.New(err, statusCode)

	// When
	newLErr := lanaerr.FromErr(lErr)

	// Then
	assert.EqualError(t, newLErr, "my custom error")
	assert.EqualError(t, newLErr.GetError(), "my custom error")
	assert.Equal(t, 404, newLErr.GetStatusCode())
}

func TestNewFromStandardError(t *testing.T) {
	// Given
	err := errors.New("my custom error")

	// When
	newLErr := lanaerr.FromErr(err)

	// Then
	assert.EqualError(t, newLErr, "my custom error")
	assert.EqualError(t, newLErr.GetError(), "my custom error")
	assert.Equal(t, 500, newLErr.GetStatusCode())
}
