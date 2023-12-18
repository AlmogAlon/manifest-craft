package component

import (
	"github.com/stretchr/testify/assert"
	"manifest-craft/models"
	"testing"
)

func TestGetComponentOptions(t *testing.T) {
	server := NewComponentService()
	component := &models.Component{
		Source:        "databaseInstances",
		ComponentType: "ComboBox",
	}

	expectedValues := []string{"Prod", "Local", "staging"}

	componentOptions := server.GetComponentOptions(component)

	assert.Equal(t, expectedValues, componentOptions)
}
