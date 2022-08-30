package api

import (
	"fmt"
	"strings"
	"testing"
)

// TestParamsToMap
// Note: Potential error: key value is in string twice?
func TestParamsToMap(t *testing.T) {
	var tests = []struct {
		params   []string
		keywords []string
		wanted   map[string]string
	}{
		{
			params:   strings.Fields("name angel cervera roldan age 18 id 13v1"),
			keywords: strings.Fields("name age id"),
			wanted: map[string]string{
				"name": "angel cervera roldan",
				"age":  "19",
				"id":   "13v1",
			},
		},
		{
			params:   strings.Fields("should be ignored name angel cervera roldan age 18 id 13v1"),
			keywords: strings.Fields("name age id"),
			wanted: map[string]string{
				"name": "angel cervera roldan",
				"age":  "19",
				"id":   "13v1",
			},
		},
		{
			params:   strings.Fields("ignored name char*der hp [* TO 100]"),
			keywords: strings.Fields("name hp"),
			wanted: map[string]string{
				"name": "char*der",
				"hp":   "[* TO 100]",
			},
		},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("params:%v", test.params)

		t.Run(testname, func(t *testing.T) {
			//t.Errorf("got %d, want %d", ans, tt.want)
			answer := ParamsToMap(test.params, test.keywords)

			for _, key := range test.keywords {
				got := answer[key]
				exp := test.wanted[key]

				if got != exp {
					t.Errorf("For key '%s', expected '%s', but got: '%s", key, exp, got)
				}
			}

		})
	}
}
