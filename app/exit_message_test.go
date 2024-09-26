package app

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/wtfutil/wtf/support"
	"gotest.tools/assert"
	"strings"
	"testing"
)

func Test_displayExitMessage(t *testing.T) {
	tests := []struct {
		name          string
		isDisplayable bool
		isContributor bool
		isSponsor     bool
		compareWith   string
		expected      string
	}{
		{
			name:          "when not displayable",
			isDisplayable: false,
			isContributor: true,
			isSponsor:     true,
			compareWith:   "equals",
			expected:      "",
		},
		{
			name:          "when contributor",
			isDisplayable: true,
			isContributor: true,
			compareWith:   "contains",
			expected:      "thank you for contributing",
		},
		{
			name:          "when sponsor",
			isDisplayable: true,
			isSponsor:     true,
			compareWith:   "contains",
			expected:      "Thank you for sponsoring",
		},
		{
			name:          "when user",
			isDisplayable: true,
			isContributor: false,
			isSponsor:     false,
			compareWith:   "contains",
			expected:      "supported by sponsorships",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wtfApp := WtfApp{}
			wtfApp.ghUser = &support.GitHubUser{
				IsContributor: tt.isContributor,
				IsSponsor:     tt.isSponsor,
			}

			actual := wtfApp.displayExitMsg(tt.isDisplayable)

			if tt.compareWith == "equals" {
				assert.Equal(t, actual, tt.expected)
			}

			if tt.compareWith == "contains" {
				assert.Equal(t, true, strings.Contains(actual, tt.expected))
			}
		})
	}
}

func Test_contributorThankYouMessage(t *testing.T) {
	wtfApp := WtfApp{}
	actual := wtfApp.contributorThankYouMessage()

	t.Run("contains main message thank message", func(t *testing.T) {
		expected := "On behalf of all the users of WTF, thank you for contributing to the source code."

		if strings.Contains(actual, expected) == false {
			t.Errorf("Expected '%s' to be present in '%s'", expected, actual)
		}
	})

	t.Run("contains a green colored message", func(t *testing.T) {
		expected := fmt.Sprintf(" %s", aurora.Green("\n\n    You rock."))

		if strings.Contains(actual, expected) == false {
			t.Errorf("Expected '%s' to be present in '%s'", expected, actual)
		}
	})
}

func Test_sponsorThankYouMessage(t *testing.T) {
	wtfApp := WtfApp{}
	actual := wtfApp.sponsorThankYouMessage()

	t.Run("contains main message thank message", func(t *testing.T) {
		expected := "Your sponsorship of WTF makes a difference. Thank you for sponsoring and supporting WTF."

		if strings.Contains(actual, expected) == false {
			t.Errorf("Expected '%s' to be present in '%s'", expected, actual)
		}
	})

	t.Run("contains a green colored message", func(t *testing.T) {
		expected := fmt.Sprintf(" %s", aurora.Green("\n\n    You're awesome."))

		if strings.Contains(actual, expected) == false {
			t.Errorf("Expected '%s' to be present in '%s'", expected, actual)
		}
	})
}

func Test_supportRequestMessage(t *testing.T) {
	wtfApp := WtfApp{}
	actual := wtfApp.supportRequestMessage()

	t.Run("contains main message thank message", func(t *testing.T) {
		expected := "The development and maintenance of WTF is supported by sponsorships."

		if strings.Contains(actual, expected) == false {
			t.Errorf("Expected '%s' to be present in '%s'", expected, actual)
		}
	})

	t.Run("contains a green colored message", func(t *testing.T) {
		expected := fmt.Sprintf("    Sponsor the development of WTF at %s\n", aurora.Green("https://github.com/sponsors/senorprogrammer"))

		if strings.Contains(actual, expected) == false {
			t.Errorf("Expected '%s' to be present in '%s'", expected, actual)
		}
	})
}
