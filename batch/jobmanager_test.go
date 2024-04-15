package batch_test

import (
	"errors"
	"testing"

	"github.com/remiges-tech/alya/batch"
	"github.com/stretchr/testify/assert"
)

func TestRegisterInitializer(t *testing.T) {
	jm := batch.NewJobManager(nil, nil, nil)

	// Create a mock initializer
	mockInitializer := &MockInitializer{}

	// Test registering a new initializer
	err := jm.RegisterInitializer("app1", mockInitializer)
	assert.NoError(t, err)

	// Test registering a duplicate initializer
	err = jm.RegisterInitializer("app1", mockInitializer)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, batch.ErrInitializerAlreadyRegistered))
	assert.Equal(t, "initializer already registered for this app: app=app1", err.Error())
}

type MockInitializer struct{}

func (i *MockInitializer) Init(app string) (batch.InitBlock, error) {
	return nil, nil
}
