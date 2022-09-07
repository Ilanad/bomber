package providers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devops-kung-fu/bomber/providers/ossindex"
	"github.com/devops-kung-fu/bomber/providers/osv"
)

func TestNewProvider(t *testing.T) {
	provider, err := NewProvider("ossindex")
	assert.NoError(t, err)
	assert.IsType(t, ossindex.OSSIndexProvider{}, provider)

	provider, err = NewProvider("osv")
	assert.NoError(t, err)
	assert.IsType(t, osv.OSVProvider{}, provider)

	provider, err = NewProvider("test")
	assert.Error(t, err)
}