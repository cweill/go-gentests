package testdata

import "testing"

func TestFoo15(t *testing.T) {
	tests := []struct {
		name    string
		arg     func(string) (string, error)
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo15(tt.arg); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo15() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
