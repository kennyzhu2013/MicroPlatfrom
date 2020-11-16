package config

import (
	"testing"
)

func TestValues(t *testing.T) {
	data := []byte(`{"foo": "bar", "baz": {"bar": "cat"}}`)

	testData := []struct {
		path  []string
		value string
	}{
		{
			[]string{"foo"},
			"bar",
		},
		{
			[]string{"baz", "bar"},
			"cat",
		},
	}

	values, err := newValues(&ChangeSet{
		Data: data,
	})

	if err != nil {
		t.Fatal(err)
	}

	for _, test := range testData {
		if v := String(""); v != test.value {
			t.Fatalf("Expected %s got %s for path %v", test.value, v, test.path)
		}
	}
}