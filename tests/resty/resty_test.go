package resty_test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestExampleAPI(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		Get("https://jsonplaceholder.typicode.com/posts/1")

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, resp.String(), "userId")
}
