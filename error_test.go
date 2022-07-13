package zispc

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		setError(fmt.Errorf(tt.name))
		e := fmt.Errorf(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if err := Error(); (err == nil) || !errors.Is(err, e) {
				t.Errorf("Error() error = %v, e = %v", err, e)
			}
		})
	}
}
