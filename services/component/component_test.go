package component

import (
	"github.com/stretchr/testify/assert"
	"manifest-craft/models"
	"testing"
)

func TestGetComponentOptions(t *testing.T) {
	type testCase struct {
		input *models.Component
		want  []string
	}

	tests := []testCase{
		{
			input: &models.Component{
				Source:        "databaseInstances",
				ComponentType: "ComboBox",
			},
			want: []string{"Prod", "Local", "staging"},
		},
		{
			input: &models.Component{
				Source:        "non-exist",
				ComponentType: "TextField",
			},
			want: []string{},
		},
	}

	server := NewComponentService()
	for _, tc := range tests {
		got, _ := server.GetComponentOptions(tc.input)
		assert.Equal(t, tc.want, got)
	}

}
