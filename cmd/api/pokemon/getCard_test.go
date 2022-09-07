package pokemon

import "testing"

const (
// Card was denoted by id to make name unique
)

func TestURLEncode(t *testing.T) {
	nonEncodedURL := "http://test.com?q=name: \"this is some text\""
	expected := "http://test.com?q=name:%20%22this%20is%20some%20text%22"
	got := URLEncode(nonEncodedURL)

	if got != expected {
		t.Errorf("Url did not encode as expected\n\tGot '%s', and expected: '%s'", got, expected)
	}
}

func TestGetCardById(t *testing.T) {
	id, name := "gym2-2", "Blaine's Charizard"
	gotCard, err := getCardById(id)

	if err != nil {
		t.Errorf("Error when trying to get card: '%s'", err)
	}

	if gotCard.Id != id {
		t.Errorf("Got wrong card from ID\n\tGot ID: '%s', and expected ID: '%s'", gotCard.Id, id)
	}

	if gotCard.Name != name {
		t.Errorf("Got wrong card from ID\n\tGot name: '%s', and expected name: '%s'", gotCard.Name, name)
	}
}

// TODO: This test needs to be more in depth as there are lots of case scenarios
func TestGetCardByParams(t *testing.T) {
	tests := []struct {
		expectedId         string
		returnedCardsCount int
		searchParams       map[string]string
	}{
		{
			// basic test
			expectedId:         "gym2-2",
			returnedCardsCount: 1,
			searchParams: map[string]string{
				"name":  "Blaine's Charizard",
				"hp":    "100",
				"level": "50",
			},
		},
		{
			// wildcard test
			expectedId:         "gym2-2",
			returnedCardsCount: 1,
			searchParams: map[string]string{
				"name":  "Blaine's Chariz*",
				"hp":    "100",
				"level": "50",
			},
		},
		{
			// Multiple cards returned
			returnedCardsCount: 240,
			searchParams: map[string]string{
				"name": "cha*",
			},
		},
	}

	for _, test := range tests {
		gotCards, err := getCardsByParams(test.searchParams)
		if err != nil {
			t.Errorf("Error when trying to get card: '%s'", err)
		}

		if len(gotCards) != test.returnedCardsCount {
			t.Errorf("Expected %d cards, got %d with params: '%v'", test.returnedCardsCount, len(gotCards), test.searchParams)
		}

		// If we only have one card, check that its the right one
		if test.returnedCardsCount == 1 {
			if gotCards[0].Id != test.expectedId {
				t.Errorf("Got unexpected card id when searching with params: '%v'\n\tGot ID: '%s', and expected ID: '%s'", test.searchParams, gotCards[0].Id, test.expectedId)
			}
		}
	}
}
