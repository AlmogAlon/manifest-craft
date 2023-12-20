package component

import (
	"github.com/stretchr/testify/assert"
	"manifest-craft/models"
	"manifest-craft/storage"
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

func TestValidate(t *testing.T) {
	type testCase struct {
		components *[]models.Component
		payload    map[string]interface{}
		want       bool
	}

	s := storage.NewMemoryStorage()

	radioButtonTests := []testCase{
		{
			components: &[]models.Component{
				*&models.Component{
					ManifestID:    1,
					Source:        "active",
					Label:         "active",
					ComponentType: "RadioButton",
					InputType:     "Boolean",
				},
			},
			payload: map[string]interface{}{
				"active": true,
			},
			want: true,
		},
		{
			components: &[]models.Component{
				*&models.Component{
					ManifestID:    1,
					Source:        "active",
					Label:         "active",
					ComponentType: "TextField",
					InputType:     "Boolean",
				},
			},
			payload: map[string]interface{}{
				"active": true,
			},
			want: false,
		},
		{
			components: &[]models.Component{
				*&models.Component{
					ManifestID:    1,
					Source:        "active",
					Label:         "active",
					ComponentType: "RadioButton",
					InputType:     "String",
				},
			},
			payload: map[string]interface{}{
				"active": true,
			},
			want: false,
		},
		{
			components: &[]models.Component{
				*&models.Component{
					ManifestID:    1,
					Source:        "active",
					Label:         "active",
					ComponentType: "RadioButton",
					InputType:     "Boolean",
				},
			},
			payload: map[string]interface{}{
				"active": "1",
			},
			want: false,
		},
	}

	stringTests := []testCase{
		{
			components: &[]models.Component{
				*s.GetComponent("reason"),
			},
			payload: map[string]interface{}{
				"reason": "test",
			},
			want: true,
		},
		{
			components: &[]models.Component{
				*s.GetComponent("reason"),
			},
			payload: map[string]interface{}{
				"reason": 3,
			},
			want: false,
		},
	}

	comboxTests := []testCase{
		{
			components: &[]models.Component{
				*s.GetComponent("databaseRoles"),
			},
			payload: map[string]interface{}{
				"databaseRoles": "Full Admin",
			},
			want: true,
		},
		{
			components: &[]models.Component{
				*s.GetComponent("databaseRoles"),
			},
			payload: map[string]interface{}{
				"databaseRoles": "non-exist",
			},
			want: false,
		},
		{
			components: &[]models.Component{
				*s.GetComponent("databaseRoles"),
			},
			payload: map[string]interface{}{
				"databaseRoles": 2,
			},
			want: false,
		},
		{
			components: &[]models.Component{
				*s.GetComponent("databaseRoles"),
			},
			payload: map[string]interface{}{
				"databaseRoles": []string{"Full Admin"},
			},
			want: false,
		},
	}

	tests := append(stringTests, comboxTests...)
	tests = append(tests, radioButtonTests...)

	server := NewComponentService()
	for _, tc := range tests {
		got, _ := server.Validate(*tc.components, tc.payload)
		assert.Equal(t, tc.want, got)
	}
}
