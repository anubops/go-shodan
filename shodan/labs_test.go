package shodan

import (
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_CalcHoneyScore(t *testing.T) {
	setUpTestServe()
	defer tearDownTestServe()

	ip := "192.168.0.1"

	mux.HandleFunc(fmt.Sprintf(honeyscorePath, ip), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `0.5`)
	})

	score, err := client.CalcHoneyScore(ip)
	assert.Nil(t, err)
	assert.Equal(t, 0.5, score)
}

func TestClient_CalcHoneyScore_invalidIP(t *testing.T) {
	client := NewClient(nil, testClientToken)
	_, err := client.CalcHoneyScore("invalid-ip")

	assert.NotNil(t, err)
	_, ok := err.(*net.ParseError)
	assert.True(t, ok)
}

func TestClient_CalcHoneyScore_invalidBaseURL(t *testing.T) {
	client := NewClient(nil, testClientToken)
	client.BaseURL = ":/1232.22"
	_, err := client.CalcHoneyScore("192.168.0.1")
	assert.NotNil(t, err)
}
