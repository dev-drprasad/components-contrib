// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------
// +build integration

package secretmanager

import (
	"github.com/dapr/components-contrib/secretstores"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestIntegrationGetSecret requires AWS specific environments for authentication AWS_DEFAULT_REGION AWS_ACCESS_KEY_ID,
// AWS_SECRET_ACCESS_KkEY and AWS_SESSION_TOKEN
func TestIntegrationGetSecret(t *testing.T) {
	secretName := "/aws/secret/testing"
	sm := NewSecretManager()
	err := sm.Init(secretstores.Metadata{
		Properties: map[string]string{
			"Region":       os.Getenv("AWS_DEFAULT_REGION"),
			"AccessKey":    os.Getenv("AWS_ACCESS_KEY_ID"),
			"SecretKey":    os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"SessionToken": os.Getenv("AWS_SESSION_TOKEN"),
		},
	})
	assert.Nil(t, err)
	response, err := sm.GetSecret(secretstores.GetSecretRequest{
		Name:     secretName,
		Metadata: map[string]string{},
	})
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
