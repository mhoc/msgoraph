package client

import (
	"context"
	"testing"

	"github.com/mhoc/msgoraph/scopes"
)

func TestHeadlessClientInitialization(t *testing.T) {
	applicationID := ""
	applicationSecret := ""
	c := NewHeadless(applicationID, applicationSecret, scopes.All(scopes.PermissionTypeApplication))
	err := c.InitializeCredentials(context.Background())
	if err != nil {
		t.Fatalf(err.Error())
	}
}
