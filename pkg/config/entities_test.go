package config_test

import (
	"testing"

	"github.com/Delta-a-Sierra/wedding_website/pkg/config"
)

func TestPresentationMethod_Validate(t *testing.T) {
	tests := map[string]struct {
		presentationMethod config.PresentationMethod
		expectedError      bool
	}{
		"Valid Presentation Method": {
			presentationMethod: config.HTMXPM,
			expectedError:      false,
		},
		"Invalid Presentation Method": {
			presentationMethod: "invalid",
			expectedError:      true,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			err := tc.presentationMethod.Validate()
			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got error: %v", tc.expectedError, err)
			}
		})
	}
}
