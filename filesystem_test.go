package filesystem

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeReadCloser struct {
	buff *bytes.Buffer
}

func newFakeReadCloser[T string | []byte](data T) *fakeReadCloser {
	return &fakeReadCloser{buff: bytes.NewBuffer([]byte(data))}
}

func (s *fakeReadCloser) Read(p []byte) (n int, err error) {
	return s.buff.Read(p)
}

func (s *fakeReadCloser) Close() error {
	s.buff.Reset()
	return nil
}

func TestReadContentOf(t *testing.T) {
	t.Run("it should pass execution to provided handler", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionReadContentOfHandler(func(path string) (Content, error) {
			assert.Equal(t, expectedPath, path)
			return []byte("TEST"), nil
		}))

		// WHEN
		result, err := ReadContentOf(fs, expectedPath)

		// THEN
		require.NoError(t, err)
		assert.Equal(t, "TEST", result.String())
	})
	t.Run("it should return up-stream error", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionReadContentOfHandler(func(path string) (Content, error) {
			assert.Equal(t, expectedPath, path)
			return nil, errors.New("something went wrong")
		}))

		// WHEN
		result, err := ReadContentOf(fs, expectedPath)

		// THEN
		require.EqualError(t, err, "failed to read content of path/to/file: something went wrong")
		assert.Nil(t, result)
	})
}

func TestStreamContentOf(t *testing.T) {
	t.Run("it should pass execution to provided handler", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionStreamContentOfHandler(func(path string) (io.ReadCloser, error) {
			assert.Equal(t, expectedPath, path)
			return newFakeReadCloser([]byte("TEST")), nil
		}))

		// WHEN
		result, err := StreamContentOf(fs, expectedPath)

		// THEN
		require.NoError(t, err)

		val, _ := io.ReadAll(result)
		assert.Equal(t, "TEST", string(val))
	})
	t.Run("it should return up-stream error", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionStreamContentOfHandler(func(path string) (io.ReadCloser, error) {
			assert.Equal(t, expectedPath, path)
			return nil, errors.New("something went wrong")
		}))

		// WHEN
		result, err := StreamContentOf(fs, expectedPath)

		// THEN
		require.EqualError(t, err, "failed to attach reader to path/to/file: something went wrong")
		assert.Nil(t, result)
	})
}

func TestCheckIfExists(t *testing.T) {
	t.Run("it should pass execution to provided handler", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionCheckIfExistsHandler(func(path string) (bool, error) {
			assert.Equal(t, expectedPath, path)
			return true, nil
		}))

		// WHEN
		result, err := CheckIfExists(fs, expectedPath)

		// THEN
		require.NoError(t, err)
		assert.True(t, result)
	})
}
