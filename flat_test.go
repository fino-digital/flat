package flat

import (
	"encoding/json"
	"reflect"
	"testing"
)

// func TestFlatten(t *testing.T) {
// 	tests := []struct {
// 		given   string
// 		want    map[string]interface{}
// 		options Options
// 	}{
// 		// test with different primitives
// 		// String: 'good morning',
// 		// Number: 1234.99,
// 		// Boolean: true,
// 		// null: null,
// 		{
// 			`{"hello":{"world":"good morning"}}`,
// 			map[string]interface{}{"hello.world": "good morning"},
// 			Options{},
// 		},
// 		{
// 			`{"hello":{"world":1234.99}}`,
// 			map[string]interface{}{"hello.world": 1234.99},
// 			Options{},
// 		},
// 		{
// 			`{"hello":{"world":true}}`,
// 			map[string]interface{}{"hello.world": true},
// 			Options{},
// 		},
// 		{
// 			`{"hello":{"world":null}}`,
// 			map[string]interface{}{"hello.world": nil},
// 			Options{},
// 		},
// 		// slice
// 		{
// 			`{"hello":{"world":["one","two"]}}`,
// 			map[string]interface{}{
// 				"hello.world.0": "one",
// 				"hello.world.1": "two",
// 			},
// 			Options{},
// 		},
// 		// nested twice
// 		{
// 			`{"hello":{"world":{"again":"good morning"}}}`,
// 			map[string]interface{}{"hello.world.again": "good morning"},
// 			Options{},
// 		},
// 		// multiple keys
// 		{
// 			`{
// 				"hello": {
// 					"lorem": {
// 						"ipsum":"again",
// 						"dolor":"sit"
// 					}
// 				},
// 				"world": {
// 					"lorem": {
// 						"ipsum":"again",
// 						"dolor":"sit"
// 					}
// 				}
// 			}`,
// 			map[string]interface{}{
// 				"hello.lorem.ipsum": "again",
// 				"hello.lorem.dolor": "sit",
// 				"world.lorem.ipsum": "again",
// 				"world.lorem.dolor": "sit"},
// 			Options{},
// 		},
// 		// empty object
// 		{
// 			`{"hello":{"empty":{"nested":{}}}}`,
// 			map[string]interface{}{"hello.empty.nested": map[string]interface{}{}},
// 			Options{},
// 		},
// 		// custom delimiter
// 		{
// 			`{"hello":{"world":{"again":"good morning"}}}`,
// 			map[string]interface{}{"hello:world:again": "good morning"},
// 			Options{Delimiter: ":"},
// 		},
// 		// custom depth
// 		{
// 			`{
// 				"hello": {
// 					"world": {
// 						"again": "good morning"
// 					}
// 				},
// 				"lorem": {
// 					"ipsum": {
// 						"dolor": "good evening"
// 					}
// 				}
// 			}
// 			`,
// 			map[string]interface{}{
// 				"hello.world": map[string]interface{}{"again": "good morning"},
// 				"lorem.ipsum": map[string]interface{}{"dolor": "good evening"},
// 			},
// 			Options{MaxDepth: 2},
// 		},

// 		// should parse array when safe = true
// 		{
// 			`{"hello":{"world":["one","two"]}}`,
// 			map[string]interface{}{
// 				"hello.world": []interface{}{"one", "two"},
// 			},
// 			Options{Safe: true},
// 		},
// 		// todo: why this does not work
// 		// {
// 		// 	`
// 		// 	{
// 		// 		"hello": [{
// 		// 			"world": {
// 		// 				"again": "foo"
// 		// 			}
// 		// 		}, {
// 		// 			"lorem": "ipsum"
// 		// 		}],
// 		// 		"another": {
// 		// 			"nested": [{
// 		// 				"array": {
// 		// 					"too": "deep"
// 		// 				}
// 		// 			}]
// 		// 		},
// 		// 		"lorem": {
// 		// 			"ipsum": "whoop"
// 		// 		}
// 		// 	}
// 		// 	`,
// 		// 	map[string]interface{}{
// 		// 		"hello": []map[string]interface{}{
// 		// 			{
// 		// 				"world": map[string]interface{}{
// 		// 					"again": "foo",
// 		// 				},
// 		// 			},
// 		// 			{
// 		// 				"lorem": "ipsum",
// 		// 			},
// 		// 		},
// 		// 		"another.nested": []map[string]interface{}{
// 		// 			{
// 		// 				"array": map[string]interface{}{
// 		// 					"too": "deep",
// 		// 				},
// 		// 			},
// 		// 		},
// 		// 		"lorem.ipsum": "whoop",
// 		// 	},
// 		// 	Options{Safe: true},
// 		// },
// 	}
// 	for i, test := range tests {
// 		var given interface{}
// 		err := json.Unmarshal([]byte(test.given), &given)
// 		if err != nil {
// 			t.Errorf("%d: failed to unmarshal test: %v", i+1, err)
// 		}
// 		got, err := Flatten(given.(map[string]interface{}), test.options)
// 		if err != nil {
// 			t.Errorf("%d: failed to flatten: %v", i+1, err)
// 		}
// 		if !reflect.DeepEqual(got, test.want) {
// 			t.Errorf("%d: mismatch, got: %v want: %v", i+1, got, test.want)
// 		}
// 	}
// }

func TestFPrimitive(t *testing.T) {
	tests := []struct {
		prefix  string
		nested  interface{}
		options Options
		want    map[string]interface{}
	}{
		// test with different primitives
		// String: 'good morning',
		// Number: 1234.99,
		// Boolean: true,
		// null: null,
		{
			"hello",
			"world",
			Options{},
			map[string]interface{}{"hello": "world"},
		},
		{
			"hello",
			1234.49,
			Options{},
			map[string]interface{}{"hello": 1234.49},
		},
		{
			"hello",
			true,
			Options{},
			map[string]interface{}{"hello": true},
		},
		{
			"hello",
			nil,
			Options{},
			map[string]interface{}{"hello": nil},
		},
	}
	for i, test := range tests {
		got, err := f(test.prefix, test.nested, test.options)
		if err != nil {
			t.Errorf("%d: failed to flatten: %v", i+1, err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: mismatch, got: %v want: %v", i+1, got, test.want)
		}
	}
}
func TestFMap(t *testing.T) {
	tests := []struct {
		prefix  string
		nested  string
		options Options
		want    map[string]interface{}
	}{
		{
			"hello",
			`{"world": "good morning"}`,
			Options{
				Delimiter: ".",
			},
			map[string]interface{}{"hello.world": "good morning"},
		},
		// nested twice
		{
			"hello",
			`{"world":{"again":"good morning"}}`,
			Options{
				Delimiter: ".",
			},
			map[string]interface{}{"hello.world.again": "good morning"},
		},
		// multiple key
		{
			"hello",
			`{
				"world": {
					"again": "good morning"
				},
				"ipsum": {
					"dolor": "good evening"
				}
			}`,
			Options{
				Delimiter: ".",
			},
			map[string]interface{}{
				"hello.world.again": "good morning",
				"hello.ipsum.dolor": "good evening",
			},
		},
		// slice
		{
			"hello",
			`{"world":["one","two"]}`,
			Options{
				Delimiter: ".",
			},
			map[string]interface{}{
				"hello.world.0": "one",
				"hello.world.1": "two",
			},
		},
	}
	for i, test := range tests {
		var nested interface{}
		err := json.Unmarshal([]byte(test.nested), &nested)
		if err != nil {
			t.Errorf("%d: failed to unmarshal test: %v", i+1, err)
		}
		got, err := f(test.prefix, nested, test.options)
		if err != nil {
			t.Errorf("%d: failed to flatten: %v", i+1, err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: mismatch, got: %v want: %v", i+1, got, test.want)
		}
	}
}
