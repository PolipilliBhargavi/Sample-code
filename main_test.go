package hello

import (
	"reflect"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name string

		want1 string
	}{
		//TODO: Add test cases
		{name: "world", want1: "hello world"},
        {name: "pullrequest", want1: "hello pullrequest"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := HelloWorld()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HelloWorld got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
