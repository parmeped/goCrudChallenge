package server_test

import (
	"testing"

	"github.com/goCrudChallenge/pkg/utl/server"
)

func TestNew(t *testing.T) {
	e := server.New()
	if e == nil {
		t.Errorf("Server should not be nil")
	}
}
