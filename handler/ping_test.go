package handler

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PingTestSuite struct {
	suite.Suite
}

func (suite *PingTestSuite) TestPing() {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/ping", nil)
	Ping(w, r)

	suite.Nil(err)
	suite.Equal("{\"success\": \"pong\"}", w.Body.String())
	suite.Equal(http.StatusOK, w.Code)
	suite.Equal("application/json", w.Header().Get("Content-Type"))
}

func TestPingTestSuite(t *testing.T) {
	suite.Run(t, new(PingTestSuite))
}
