package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMode(t *testing.T) {
	t.Run("check user read/write/execute", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, Mode(0700), ModeUserReadWriteExecute)
	})
	t.Run("check group read/write/execute", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, Mode(0070), ModeGroupReadWriteExecute)
	})
	t.Run("check others read/write/execute", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, Mode(0007), ModeOthersReadWriteExecute)
	})
	t.Run("check all read/write/execute", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, Mode(0777), ModeAllReadWriteExecute)
	})
}
