package testdata

import "testing"

func TestFoo13(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		f func()
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo13(tt.f); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo13() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
