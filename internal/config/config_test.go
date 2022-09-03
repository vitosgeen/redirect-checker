package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Load config ok", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := LoadEnvConfig()
			assert.Equal(t, tt.wantErr, !(err != nil))
		})
	}
}
