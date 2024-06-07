package comparator_test

import (
	"compass/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindCoincidences_shouldMatch2contacts(t *testing.T) {
	input := []comparator.Contact{
		{
			ID:        1001,
			FirstName: "C",
			LastName:  "F",
			Email:     "mollis.lectus.pede@outlook.net",
			ZipCode:   0,
			Address:   "449-6990 Tellus. Rd.",
		},
		{
			ID:        1002,
			FirstName: "C",
			LastName:  "French",
			Email:     "mollis.lectus.pede@outlook.net",
			ZipCode:   39746,
			Address:   "449-6990 Tellus. Rd.",
		},
		{
			ID:        1003,
			FirstName: "Ciara",
			LastName:  "F",
			Email:     "non.lacinia.at@zoho.ca",
			ZipCode:   39746,
			Address:   "",
		},
	}

	results := comparator.FindCoincidences(input)

	expected := []comparator.Result{
		{
			ContactSource: 1001,
			ContactMatch:  1002,
			Accuracy:      "High",
		},
		{
			ContactSource: 1001,
			ContactMatch:  1003,
			Accuracy:      "Low",
		},
	}
	assert.Equal(t, expected, results)
}

func TestFindCoincidences_shouldNoMatchContacts(t *testing.T) {
	input := []comparator.Contact{
		{
			ID:        1001,
			FirstName: "C",
			LastName:  "F",
			Email:     "mollis.lectus.pede@outlook.net",
			ZipCode:   0,
			Address:   "449-6990 Tellus. Rd.",
		},
		{
			ID:        1002,
			FirstName: "Test",
			LastName:  "Test",
			Email:     "test@outlook.net",
			ZipCode:   39746,
			Address:   "New Address",
		},
		{
			ID:        1003,
			FirstName: "B",
			LastName:  "A",
			Email:     "non.lacinia.at@zoho.ca",
			ZipCode:   39745,
			Address:   "New Address",
		},
	}

	results := comparator.FindCoincidences(input)

	assert.Len(t, results, 0)
}
