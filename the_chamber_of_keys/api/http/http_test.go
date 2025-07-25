package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ck "the_chamber_of_keys/pkg/chamber_of_keys"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func newTestServer() *httptest.Server {
	gin.SetMode(gin.TestMode)

	chamber, _ := ck.NewChamber()
	router := NewRouter(chamber)

	return httptest.NewServer(router)
}

func TestInsertString(t *testing.T) {
	server := newTestServer()
	defer server.Close()

	insertBody := `{"key":"spell","value":"Lumos","ttl":3600}`
	resp, err := http.Post(server.URL+"/string", "application/json", strings.NewReader(insertBody))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPushItem(t *testing.T) {
	server := newTestServer()
	defer server.Close()

	client := &http.Client{}

	pushBody := `{"value":"Expelliarmus","ttl":600}`
	pushReq, err := http.NewRequest("POST", server.URL+"/list/spellbook/items?position=front", strings.NewReader(pushBody))
	pushReq.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	pushResp, err := client.Do(pushReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, pushResp.StatusCode)
	pushResp.Body.Close()
}
