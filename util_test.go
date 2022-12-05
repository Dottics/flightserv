package flightserv

import (
	"fmt"
	"testing"
)

func TestParseBody(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "empty object",
			input:  `{}`,
			output: `{}`,
		},
		{
			name:   "single line",
			input:  `{ "name": "james", "age": 23 }`,
			output: `{"name":"james","age":23}`,
		},
		{
			name: "multi line",
			input: `{ 
				"name": "james", 
				"age": 23 
			}`,
			output: `{"name":"james","age":23}`,
		},
		{
			name: "an array",
			input: `{ 
				"name": "james", 
				"age": 23,
				"favNumbers": [ 1, 4, "ten" ]
			}`,
			output: `{"name":"james","age":23,"favNumbers":[1,4,"ten"]}`,
		},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			o := ParseBody(tc.input)
			if o != tc.output {
				t.Errorf("expected string '%s' got '%s'", tc.output, o)
			}
		})
	}
}
