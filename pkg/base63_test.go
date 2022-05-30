package pkg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncodeBase63(t *testing.T) {
	tests := []struct {
		name           string
		n              int64
		expectedResult string
		expectedErr    error
	}{
		{
			name:           "check positive number",
			n:              1,
			expectedResult: "0000000001",
			expectedErr:    nil,
		},
		{
			name:           "check a negative number",
			n:              -1,
			expectedResult: "",
			expectedErr:    ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := EncodeBase63(tt.n)
			require.Equal(t, tt.expectedResult, actual)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestDecodeBase63(t *testing.T) {
	tests := []struct {
		name           string
		shortenedURL   string
		expectedResult int64
		expectedErr    error
	}{
		{
			name:           "check positive number",
			shortenedURL:   "0000000001",
			expectedResult: 1,
			expectedErr:    nil,
		},
		{
			name:           "check a invalid input",
			shortenedURL:   "###",
			expectedResult: 0,
			expectedErr:    ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := DecodeBase63(tt.shortenedURL)
			require.Equal(t, tt.expectedResult, actual)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}
