package testdata

import (
	"reflect"
	"testing"
)

func TestFoo25(t *testing.T) {
	tests := []struct {
		name    string
		arg     interface{}
		want    string
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1, err := Foo25(tt.arg)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo25() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Foo25() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. Foo25() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}
