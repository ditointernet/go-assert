package assert

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Equal compares with a deep equal if two values are the same
// if not it calls t.Fatalf() with a readable diff and returns.
func Equal(t *testing.T, expected interface{}, received interface{}, desc ...interface{}) {
	require.Equal(t, expected, received, desc)
}

// NotEqual compares with a deep equal if two values are the same
// if not it calls t.Fatalf() with a readable diff and returns.
func NotEqual(t *testing.T, expected interface{}, received interface{}, desc ...interface{}) {
	require.NotEqual(t, expected, received, desc)
}
